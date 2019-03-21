package vendorsService

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"../../../db"
	"../../../model"
	"golang.org/x/crypto/bcrypt"

	//	"../../../util/crypto"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"../../../util/random"
	"github.com/cydev/zero"
)

type Result struct {
	Vender  string
	Average float32
}

type DeviceList struct {
	DeviceID   uint
	DeviceName string
}

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateVendor(vendor *model.Vendor) (*model.Vendor, error) {
	// check duplicate username
	v := &model.Vendor{}
	if res := db.ORM.Where("username = ?", vendor.UserName).First(&v).RecordNotFound(); !res {
		err := errors.New(vendor.UserName + " is already registered")
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(vendor.Password), 10)
	if err != nil {
		return nil, err
	}
	vendor.Password = string(password)
	// email send
	verifyCode := random.GenerateRandomString(10)
	//	hashPass := crypto.GenerateHash(verifyCode)
	fmt.Println(verifyCode)
	db.ORM.UpdateColumn("password", verifyCode)
	// email.SendForgotEmail(vendor.Email, vendor.UserName, verifyCode)

	// Insert Data
	if err := db.ORM.Create(&vendor).Error; err != nil {
		return nil, err
	}
	return vendor, err
}

// CreateVendorByWorking creates a user
func CreateVendorByWorking(wkinfo *model.WorktimeInfo) (*model.WorktimeInfo, error) {

	opentm := wkinfo.OpenTime
	specdt := wkinfo.SpecialDate
	fmt.Println(wkinfo.VendorID)
	//	res.Where(model.Customer{Email: pro.Email}).Assign(*pro).FirstOrCreate(&pro)
	//	err := res.Find(&pro).Error

	db.ORM.Table("worktimes").Where("vendor_id = ?", wkinfo.VendorID).Delete(&opentm)

	for _, o := range opentm {
		o.VendorID = wkinfo.VendorID
		db.ORM.Table("worktimes").NewRecord(o)
		db.ORM.Table("worktimes").Create(&o)

		//		db.ORM.Last(&o)
	}
	//	db.ORM.Table("worktimes").Where("vendor_id = ?", wkinfo.VendorID).Update("vendor_id", wkinfo.VendorID).Find(&opentm)
	db.ORM.Table("specialtimes").Where("vendor_id = ?", wkinfo.VendorID).Delete(&specdt)
	specdt = wkinfo.SpecialDate
	for _, s := range specdt {
		s.VendorID = wkinfo.VendorID
		db.ORM.Table("specialtimes").NewRecord(s)
		db.ORM.Table("specialtimes").Create(&s)
	}
	//	db.ORM.Table("specialtimes").Where("vendor_id = ?", wkinfo.VendorID).Update("vendor_id", wkinfo.VendorID).Find(&specdt)
	return wkinfo, nil
}

// CreateUser creates a user
func CreateVendorWithEmail(vendor *model.Vendor) (*model.Vendor, error) {
	// check duplicate username
	v := &model.Vendor{}

	//	vendor.UserName = vendor.Email
	if res := db.ORM.Where("username = ?", vendor.UserName).First(&v).RecordNotFound(); !res {
		err := errors.New(vendor.UserName + " is already registered")
		return nil, err
	}

	verifyCode := random.GenerateRandomString(8)
	//password, err := bcrypt.GenerateFromPassword([]byte(verifyCode), 8)
	//if err != nil {
	//	return nil, err
	//}
	//vendor.Password = string(password)
	vendor.Password = string(verifyCode)
	// email send
	fmt.Println(verifyCode)
	fmt.Println(vendor.Email)
	fmt.Println(vendor.UserName)

	// go email.SendForgotEmail(vendor.Email, vendor.UserName, verifyCode)

	// Insert Data
	if err := db.ORM.Create(&vendor).Error; err != nil {
		return nil, err
	}
	return vendor, nil
}

// ReadUser reads a user
//func ReadVendor(id uint) ([]*model.DeviceInfo, error) {
func ReadVendor(id uint) (*model.Vendor, error) {
	vendor := &model.Vendor{}
	devinfo := []*model.DeviceInfo{}

	res := db.ORM
	res.Table("vendors").Select("vendor_devices.device_id as device_id, devices.device_name as device_name").
		Joins("left join vendor_devices on vendors.id = vendor_devices.vendor_id left join devices on device_id = devices.id").
		Find(&devinfo, "vendors.id = ?", id)

		//	err := res.Find(&vendor, "vendors.id = ?", id).Error
	// Read Data

	vendor.DeviceList = devinfo
	err := db.ORM.First(&vendor, "id = ?", id).Error

	return vendor, err
}

//ReadVendorByWorking
func ReadVendorByWorking(vendorID uint) (*model.WorktimeInfo, error) {
	wkinfo := &model.WorktimeInfo{}
	opentm := []*model.Worktime{}
	specdt := []*model.Specialtime{}

	db.ORM.Where("vendor_id = ?", vendorID).Find(&opentm)
	db.ORM.Where("vendor_id = ? AND kind = ?", vendorID, 1).Find(&specdt)

	wkinfo.OpenTime = opentm
	wkinfo.SpecialDate = specdt
	return wkinfo, nil
}

// UpdateUser reads a user
func UpdateVendor(vendor *model.Vendor) (*model.Vendor, error) {
	// Create change info
	err := db.ORM.Model(vendor).Updates(vendor).Error
	return vendor, err
}

// DeleteUser deletes user with object id
func DeleteVendor(id uint) error {
	vendor := &model.Vendor{}

	err := db.ORM.Where("id = ?", id).Delete(vendor).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadVendors(query string, offset int, count int, field string, sort int) ([]*model.Vendor, int, error) {
	vendors := []*model.Vendor{}
	totalCount := 0

	res := db.ORM

	results := []*Result{}
	res.Table("reviews").Where("type = ?", 1).Select("vender_id as vender, avg(score) as average").Group("vender_id").Scan(&results)

	fmt.Println(results)
	venderIds := []string{}
	for _, r := range results {
		fmt.Println(r.Vender, r.Average)
		venderIds = append(venderIds, r.Vender)
		res.Table("vendors").Where("vendors.id = ?", r.Vender).Update("rate", r.Average)
	}
	dbb := res.Table("vendors")
	//	dbb := res.Table("vendors").Where("vendors.id IN (?)", venderIds)
	// if query != "" {
	// 	query = "%" + query + "%"
	// 	vdd = vdd.Where("username LIKE ? OR password ?", query, query)
	// }
	// get total count of collection with initial query
	if err := dbb.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// add page feature

	if offset != 0 || count != 0 {
		dbb = dbb.Offset(offset)
		dbb = dbb.Limit(count)
	}
	// add sort feature
	if field != "" && sort != 0 {
		if sort > 0 {
			dbb = dbb.Order(field)
		} else {
			dbb = dbb.Order(field + " desc")
		}
	}
	err := dbb.Find(&vendors).Error

	return vendors, totalCount, err
}

// ReadUserByUsername returns user
func ReadVendorByUsername(username string) (*model.Vendor, error) {
	vendor := &model.Vendor{}
	fmt.Println("time star:", time.Now())
	res := db.ORM.Where("username = ?", username).First(&vendor).RecordNotFound()
	if res {
		return nil, errors.New("User doesnot exist")
	}

	return vendor, nil
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Vendor, int, error) {
	vendors := []*model.Vendor{}
	totalCount := 0

	res := db.ORM
	// if userID != 0 {
	// 	res = res.Where("userID = ?", userID)
	// }
	if query != "" {
		query = "%" + query + "%"
		res = res.Where(fmt.Sprintf("%v LIKE ? OR %v LIKE ?", field, field), query, query)
	}
	// get total count of collection with initial query
	res.Find(&vendors).Count(&totalCount)

	// add page feature
	if offset != 0 || count != 0 {
		res = res.Offset(offset)
		res = res.Limit(count)
	}
	// add sort feature
	if field != "" && sort != 0 {
		if sort > 0 {
			res = res.Order(field)
		} else {
			res = res.Order(field + " desc")
		}
	}
	err := res.Find(&vendors).Error

	return vendors, totalCount, err
}

// ReadByCustomer returns user
func ReadByCustomer(count int, vendorID uint) ([]*model.Customer, int, error) {
	customerlist := []*model.Customer{}
	totalCount := 0

	res := db.ORM
	order := &model.Order{}
	db.ORM.Find(&order)

	res = res.Table("customers").Select("customers.*").
		Joins("left join orders on customers.id = orders.customer_id")

	res.Where("orders.vendor_id = ?", vendorID).Group("orders.customer_id").Find(&customerlist)

	//	err := res.Find(&customers).Error

	return customerlist, totalCount, nil
}

// ReadByFilter return user
func ReadByFilter(postval *model.PostVendor) ([]*model.Vendor, error) {
	vendors := []*model.Vendor{}

	//	totalCount := 0

	res := db.ORM

	patterns := []string{}
	values := []interface{}{}
	fields := []interface{}{}
	pattern := "%v = ?"
	comma := " AND "

	if postval.PostCode != "" {
		patterns = append(patterns, pattern)
		fields = append(fields, "post_code")
		values = append(values, postval.PostCode)
	}

	if postval.HasParking == 1 {
		patterns = append(patterns, pattern)
		fields = append(fields, "has_parking")
		values = append(values, postval.HasParking)
	}
	if postval.DropOff == 1 {
		patterns = append(patterns, pattern)
		fields = append(fields, "drop_off")
		values = append(values, postval.DropOff)
	}
	if postval.Collect == 1 {
		patterns = append(patterns, pattern)
		fields = append(fields, "collect")
		values = append(values, postval.Collect)
	}

	whereFormat := strings.Join(patterns, comma)

	whereStr := fmt.Sprintf(whereFormat, fields...)
	fmt.Println("************", postval.Collect)
	if postval.PostCode != "" || postval.HasParking == 1 || postval.DropOff == 1 || postval.Collect == 1 {
		fmt.Println(whereStr, values)
		res = res.Where(whereStr, values...)
	}

	if postval.Sort == "rating" {
		res = res.Order("rate desc")
	}

	err := res.Find(&vendors).Error

	// get total count of collection with initial query
	//	res.Find(&vendors).Count(&totalCount)

	// add sort feature

	devinfo := []*model.DeviceInfo{}
	wt := []*model.Worktime{}
	//	vv := []*model.VendorDevices{}
	time := time.Now()
	hour := time.Hour()
	res = db.ORM
	res.Find(&wt)

	for _, v := range vendors {
		for _, w := range wt {
			if w.VendorID == v.ID {
				if w.StartTime <= hour && w.CloseTime >= hour {
					v.IsOpen = 1
					break
				} else {
					v.IsOpen = 0

				}
			}

		}
	}

	for _, vd := range vendors {
		res.Table("vendors").Select("vendor_devices.device_id as device_id, devices.device_name as device_name, devices.image as device_image").
			Joins("left join vendor_devices on vendors.id = vendor_devices.vendor_id left join devices on device_id = devices.id").
			Find(&devinfo, "vendors.id = ?", vd.ID)
		for _, k := range devinfo {
			if !zero.IsZero(k) {
				vd.DeviceList = devinfo
			} else {
				vd.DeviceList = nil
			}
		}
	}

	vendor2 := []*model.Vendor{}
	for _, v := range vendors {
		if postval.IsOpen == 0 || (postval.IsOpen != 0 && v.IsOpen == postval.IsOpen) {
			c := true
			for _, devID := range postval.DeviceID {
				b := false
				for _, dev := range v.DeviceList {
					if dev.DeviceID == devID {
						b = true
					}
				}

				if b == false {
					c = false
				}
			}

			if len(postval.DeviceID) > 0 {
				if c == true {
					fmt.Println(postval.IsOpen)
					vendor2 = append(vendor2, v)
				}
			} else {
				vendor2 = append(vendor2, v)
			}
		}
	}

	if postval.PostCode != "" {
		url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + postval.PostCode
		url = url + "&sensor=false&key=AIzaSyBl74CajeHYL2nszJXq-rQVhYrN6-9mG7A"

		req, _ := http.NewRequest("GET", url, nil)
		res_api, _ := http.DefaultClient.Do(req)

		defer res_api.Body.Close()

		body, _ := ioutil.ReadAll(res_api.Body)

		regioninfo := model.RegionInfo{}

		json.Unmarshal(body, &regioninfo)

		lat := regioninfo.Results[0].Geometry.Location.Lat
		lng := regioninfo.Results[0].Geometry.Location.Lng

		fmt.Println("lat, lng", lat, lng)

		vendors3 := []*model.Vendor{}
		for _, v := range vendor2 {
			deglen := 110.25
			x := lat - v.Lat
			y := (lng - v.Lng) * math.Cos(v.Lat)
			distance := deglen * math.Sqrt(x*x+y*y)
			v.Distance = distance
			vendors3 = append(vendors3, v)
		}

		if postval.Sort == "distance" {
			sort.Slice(vendors3, func(i, j int) bool { return vendors3[i].Distance < vendors3[j].Distance })
		}

		return vendors3, err
	}
	return vendor2, err
}
