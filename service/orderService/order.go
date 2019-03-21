package orderService

import (
	"fmt"
	"strings"
	"time"

	"../../db"
	"../../model"
	"../../util/email"
	"../../util/random"
	"../../util/timeHelper"
)

type Result struct {
	Vender string
	// Total   int64
	Average float32
}

// InitService inits service
func InitService() {

}

// CreateOrder creates a device
func CreateOrder(order *model.Order) (*model.Order, error) {
	// Insert Data
	//	if err := db.ORM.Create(&order).Error; err != nil {
	//		return nil, err
	//	}
	//	err := db.ORM.Last(&order).Error
	//return order, err
	if err := db.ORM.Related(&order, "[]model.Service").Error; err != nil {
		return nil, err
	}
	err := db.ORM.Preload("[]model.Service").Last(&order).Error
	return order, err
}

// SetOrder creates a device
func SetOrderByCustomer(obc *model.BookingInfo) (*model.BookingInfo, error) {
	// book := &model.BookingInfo{}
	fmt.Println(obc)

	order := obc.Order
	pro := obc.Profile
	vendor := &model.Vendor{}
	//refKey := []string{}
	dash := "-"
	res := db.ORM

	res.Table("customers").Where(model.Customer{Email: pro.Email}).Assign(*pro).FirstOrCreate(&pro)
	res.Table("vendors").Where(model.Vendor{ID: order.VendorID}).Assign(&vendor).FirstOrCreate(vendor)

	if order.Status == 1 { // status : fixed => coin_amount += 50
		res.Table("customers").Where(model.Customer{Email: pro.Email}).Update("coin_amount", pro.CoinAmount+50)
		res.Table("vendors").Where(model.Vendor{ID: order.VendorID}).Update("coin_amount", vendor.CoinAmount+50)
	}

	err := res.Find(&pro).Error
	fmt.Println(pro.Email, err)
	verifyCode1 := random.GenerateRandomKey(4, 0)
	verifyCode2 := random.GenerateRandomKey(4, 1)
	verifyCode3 := random.GenerateRandomKey(4, 2)
	verifyCode4 := random.GenerateRandomKey(4, 3)
	order.CustomerID = pro.ID

	//	res1.Table("order_service").Where(model.OrderService{ServiceID: os.ServiceID}).Assign(&os).FirstOrCreate(os)

	refKey := verifyCode1 + dash + verifyCode2 + dash + verifyCode3 + dash + verifyCode4

	order.ReferenceKey = string(refKey)
	// email send
	fmt.Println(refKey)

	db.ORM.NewRecord(*order)
	db.ORM.Create(order)

	res.Table("orders").Select("orders.*, services.service_name, devices.device_name, vendors.username as vendor_name," +
		"customers.username as customer_name, makes.make_name, models.model_name,  devices.image as device_image").
		Joins("left join services on services.id = orders.service_id left join devices on devices.id = orders.device_id " +
			"left join vendors on vendors.id = orders.vendor_id left join customers on customers.id = orders.customer_id " +
			"left join makes on makes.id = orders.make left join models on models.id = orders.model").Find(&order)

	DevList := order.ServiceList
	fmt.Println("0000000", DevList)
	for _, v := range DevList {
		os1 := model.OrderService{}
		os1.OrderID = order.ID
		os1.ServiceID = v
		fmt.Println(os1.OrderID, v)
		db.ORM.NewRecord(os1)
		db.ORM.Create(&os1)
	}

	// go email.SendEmail(pro.Email, order, refKey)

	return obc, nil

}

// ReadOrder reads a order
func ReadOrder(id uint) (*model.Order, error) {
	order := &model.Order{}
	// Read Data
	err := db.ORM.Table("orders").Select("orders.*, services.service_name, devices.device_name, vendors.username as vendor_name, customers.username as customer_name, makes.make_name, models.model_name, devices.image as device_image").
		Joins("left join services on services.id = orders.service_id left join devices on devices.id = orders.device_id left join vendors on vendors.id = orders.vendor_id left join customers on customers.id = orders.customer_id left join makes on makes.id = orders.make left join models on models.id = orders.model").
		First(&order, "orders.id = ?", id).Error

	//err := db.ORM.First(&order, "id = ?", id).Error

	return order, err
}

// UpdateOrder reads a order
func UpdateOrder(order *model.Order) (*model.Order, error) {
	// Create change info
	err := db.ORM.Model(order).Updates(order).Error
	return order, err
}

// UpdateOrderByStatus reads a order
func UpdateOrderByStatus(us *model.UpdateStatus) (*model.Joblog, error) {
	order := &model.Order{}
	customer := &model.Customer{}
	joblog := model.Joblog{}
	vendor := &model.Vendor{}
	// Create change info
	//	err := db.ORM.Model(order).Updates(order).Error
	err := db.ORM.Model(order).Where("vendor_id = ?", us.VendorID).Updates(&model.Order{Price: us.Price, Status: us.Status}).Error

	joblog.Status = us.Status
	joblog.Price = us.Price
	joblog.Comment = us.JobsComments
	joblog.ByEmail = us.ByEmail
	joblog.BySms = us.BySms
	joblog.VendorID = us.VendorID
	joblog.CustomerID = us.CustomerID
	joblog.OrderID = us.OrderID

	db.ORM.Table("customers").Where("id = ?", joblog.CustomerID).First(&customer)
	db.ORM.Table("joblogs").NewRecord(joblog)
	db.ORM.Table("joblogs").Create(&joblog)

	res := db.ORM
	res.Table("vendors").Where(model.Vendor{ID: joblog.VendorID}).Assign(&vendor).FirstOrCreate(vendor)
	res.Table("vendors").Where(model.Vendor{ID: joblog.VendorID}).Update("coin_amount", vendor.CoinAmount+50)

	res.Table("customers").Where(model.Customer{ID: joblog.CustomerID}).Assign(&customer).FirstOrCreate(customer)
	res.Table("customers").Where(model.Customer{ID: joblog.CustomerID}).Update("coin_amount", customer.CoinAmount+50)

	go email.SendLogEmail(customer.Email, joblog)

	return &joblog, err
}

// DeleteOrder deletes order with object id
func DeleteOrder(id uint) error {
	err := db.ORM.Delete(&model.Order{ID: id}).Error
	return err
}

// ReadOrders return orders after retreive with params
func ReadOrders(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Order, int, error) {
	orders := []*model.Order{}
	// reviews := []*model.Review{}
	totalCount := 0

	res := db.ORM

	results := []*Result{}
	res.Table("reviews").Select("vender_id as vender, avg(score) as average").Group("vender_id").Scan(&results)

	fmt.Println(results)
	for _, r := range results {
		fmt.Println(r.Vender, r.Average)
		res.Table("orders").Where("orders.vendor_id = ?", r.Vender).Update("rate", r.Average)
	}

	res = res.Table("orders").Select("orders.*, services.service_name, devices.device_name, vendors.username as vendor_name, customers.username as customer_name, makes.make_name, models.model_name,  devices.image as device_image").
		Joins("left join services on services.id = orders.service_id left join devices on devices.id = orders.device_id left join vendors on vendors.id = orders.vendor_id left join customers on customers.id = orders.customer_id left join makes on makes.id = orders.make left join models on models.id = orders.model")

	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	// if query != "" {
	// 	query = "%" + query + "%"
	// 	res = res.Where("vender_id LIKE ?", query)
	// }
	// get total count of collection with initial query
	res.Find(&orders).Count(&totalCount)

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
	err := res.Find(&orders).Error

	return orders, totalCount, err
}

// ReadByDate returns user
func ReadByDate(query string, offset int, count int, field string, sort int, vendorID uint, date string) (*model.TimeInfo, int, error) {
	timeinfos := &model.TimeInfo{}
	totalCount := 0

	res := db.ORM

	if date != "" {
		d, err1 := time.Parse("2006-1-2", date)
		if err1 != nil {
			fmt.Println(err1)
		}
		//curDate_pure := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Now().Location())
		curDate := d
		lastTime := timeHelper.FewDaysLaterDate(d, 7)

		fmt.Println(curDate, lastTime)
		res.Where("vendor_id = ? AND book_date >= ? AND book_date <= ?", vendorID, curDate, lastTime).Find(&timeinfos.OrderInfo)
		res.Where("vendor_id = ?", vendorID).Find(&timeinfos.WorktimeInfo)
		res.Where("vendor_id = ? AND date >= ? AND date <= ?", vendorID, curDate, lastTime).Find(&timeinfos.SpecialtimeInfo)
	}
	//get total count of collection with initial query
	// res.Find(&timeinfos).Count(&totalCount)

	// fmt.Println(timeinfos.WorktimeInfo)
	// // add page feature
	// if offset != 0 || count != 0 {
	// 	res = res.Offset(offset)
	// 	res = res.Limit(count)
	// }
	// // add sort feature
	// if field != "" && sort != 0 {
	// 	if sort > 0 {
	// 		res = res.Order(field)
	// 	} else {
	// 		res = res.Order(field + " desc")
	// 	}
	// }
	// err := res.Find(&timeinfos).Error

	return timeinfos, totalCount, nil
}

// ReadByBookdate returns user
func ReadByBookdate(query string, offset int, count int, field string, sort int, vendorID, customerID uint, date string) ([]*model.BookingInfo, int, error) {
	bookinginfos := []*model.BookingInfo{}
	totalCount := 0

	orders := []*model.Order{}
	pros := []*model.Customer{}

	patterns := []string{}
	values := []interface{}{}
	middle := []interface{}{}
	pattern := "%v = ?"
	comma := " AND "

	res := db.ORM

	res = res.Table("orders")
	res = res.Select("orders.*, vendors.company as company, services.service_name, devices.device_name, vendors.username as vendor_name, " +
		"customers.username as customer_name, devices.image as device_image, makes.make_name, models.model_name, customers.address as address, " +
		"customers.address2 as address2, customers.post_code as post_code").
		Joins("left join vendors on vendors.id = orders.vendor_id left join services on services.id = orders.service_id " +
			"left join devices on devices.id = orders.device_id left join customers on customers.id = orders.customer_id " +
			"left join makes on makes.id = orders.make left join models on models.id = orders.model")

	if vendorID != 0 {
		patterns = append(patterns, pattern)
		middle = append(middle, "orders.vendor_id")
		values = append(values, vendorID)
	}

	if customerID != 0 {
		patterns = append(patterns, pattern)
		middle = append(middle, "orders.customer_id")
		values = append(values, customerID)
	}

	whereFormat1 := strings.Join(patterns, comma)
	whereStr1 := fmt.Sprintf(whereFormat1, middle...)
	fmt.Println(whereStr1, values)

	if date != "" {
		d, err1 := time.Parse("2006-1-2", date)
		if err1 != nil {
			fmt.Println(err1)
		}
		//curDate_pure := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Now().Location())
		curDate := d
		lastDate := timeHelper.FewDaysLaterDate(d, 1)
		values = append(values, curDate)
		values = append(values, lastDate)
		res.Where(whereStr1+" AND orders.book_date >= ? AND orders.book_date < ?", values...).Find(&orders)
	} else {
		res.Where(whereStr1, values...).Find(&orders)
	}

	res1 := db.ORM
	//	res1 = res1.Table("customers").Select("customers.*").Joins("left join orders on orders.customer_id = customers.id")
	if customerID != 0 {
		res1 = res1.Where("customers.id = ?", customerID).Find(&pros)
	} else {
		res1 = res1.Find(&pros)
	}
	prop := &model.Customer{}
	for i := 0; i < len(orders); i++ {
		db.ORM.Table("customers").First(&prop, "id = ?", orders[i].CustomerID)
		bookinginfos = append(bookinginfos, &model.BookingInfo{Order: orders[i], Profile: prop})
	}

	return bookinginfos, totalCount, nil
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Order, int, error) {
	orders := []*model.Order{}
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
	res.Find(&orders).Count(&totalCount)

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
	err := res.Find(&orders).Error

	return orders, totalCount, err
}

// ReadByFilter returns user
func ReadByFilter(s, c, v, sort string, SortDirection, StatusFilter, DeviceFilter int) ([]*model.Order, int, error) {
	orders := []*model.Order{}
	totalCount := 0

	res := db.ORM

	patterns := []string{}
	values := []interface{}{}
	fields := []interface{}{}
	pattern := "%v LIKE ?"
	comma := " OR "
	if len(s) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "services.service_name")
		values = append(values, "%"+s+"%")
	}
	if len(c) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "customers.username")
		values = append(values, "%"+c+"%")
	}
	if len(v) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "vendors.username")
		values = append(values, "%"+v+"%")
	}
	whereFormat := strings.Join(patterns, comma)
	whereStr := fmt.Sprintf(whereFormat, fields...)

	// whereStr := fmt.Sprintf("%v LIKE ? OR %v LIKE ? OR %v LIKE ?", "services.service_name", "customers.username", "vendors.username")
	res = res.Table("orders").Select("orders.*, services.service_name, devices.device_name, vendors.username as vendor_name, customers.username as customer_name, devices.image as device_image").
		Joins("left join services on services.id = orders.service_id left join devices on devices.id = orders.device_id left join vendors on vendors.id = orders.vendor_id left join customers on customers.id = orders.customer_id")
	if len(s) > 0 || len(c) > 0 || len(v) > 0 {

		res = res.Where(whereStr, values...)
	}
	if StatusFilter != 0 && DeviceFilter != 0 {
		res = res.Where(fmt.Sprintf("status = ? AND device_id = ?"), StatusFilter, DeviceFilter)
	}
	if StatusFilter == 0 && DeviceFilter != 0 {
		res = res.Where(fmt.Sprintf("device_id = ?"), DeviceFilter)
	}
	if DeviceFilter == 0 && StatusFilter != 0 {
		res = res.Where(fmt.Sprintf("status = ?"), StatusFilter)
	}

	// get total count of collection with initial query
	res.Find(&orders).Count(&totalCount)

	// add page feature
	// if offset != 0 || count != 0 {
	// 	res = res.Offset(offset)
	// 	res = res.Limit(count)
	// }
	// add sort feature
	if sort != "" {
		if SortDirection == 1 {
			res = res.Order(sort)
		}
		if SortDirection == -1 {
			res = res.Order(sort + " desc")
		}
	}

	err := res.Find(&orders).Error

	return orders, totalCount, err
}

// ReadByVenderStatus returns user
func ReadByVenderStatus(sort string, SortDirection, StatusFilter, VenderFilter int) ([]*model.Order, int, error) {
	orders := []*model.Order{}
	totalCount := 0

	res := db.ORM

	res = res.Table("orders").Select("orders.*, services.service_name, devices.device_name, vendors.username as vendor_name, customers.username as customer_name, devices.image as device_image").
		Joins("left join services on services.id = orders.service_id left join devices on devices.id = orders.device_id left join vendors on vendors.id = orders.vendor_id left join customers on customers.id = orders.customer_id")

	if StatusFilter != 0 && VenderFilter != 0 {
		res = res.Where(fmt.Sprintf("status = ? AND vender_id = ?"), StatusFilter, VenderFilter)
	}
	if StatusFilter == 0 && VenderFilter != 0 {
		res = res.Where(fmt.Sprintf("vendor_id = ?"), VenderFilter)
	}
	if StatusFilter != 0 && VenderFilter == 0 {
		res = res.Where(fmt.Sprintf("status = ?"), StatusFilter)
	}

	// get total count of collection with initial query
	res.Find(&orders).Count(&totalCount)

	// add page feature
	// if offset != 0 || count != 0 {
	// 	res = res.Offset(offset)
	// 	res = res.Limit(count)
	// }
	// add sort feature
	if sort != "" {
		if SortDirection == 1 {
			res = res.Order(sort)
		}
		if SortDirection == -1 {
			res = res.Order(sort + " desc")
		}
	}

	err := res.Find(&orders).Error

	return orders, totalCount, err
}

// ReadByCustomerStatus returns user
func ReadByCustomerStatus(sort string, SortDirection, StatusFilter, CustomerFilter int) ([]*model.Order, int, error) {
	orders := []*model.Order{}
	totalCount := 0

	res := db.ORM

	res = res.Table("orders").Select("orders.*, vendors.company as company, services.service_name, devices.device_name, vendors.username as vendor_name, " +
		"customers.username as customer_name, devices.image as device_image, makes.make_name, models.model_name, customers.address as address, " +
		"customers.address2 as address2, customers.post_code as post_code").
		Joins("left join vendors on vendors.id = orders.vendor_id left join services on services.id = orders.service_id " +
			"left join devices on devices.id = orders.device_id left join customers on customers.id = orders.customer_id " +
			"left join makes on makes.id = orders.make left join models on models.id = orders.model")

	if StatusFilter != 0 && CustomerFilter != 0 {
		res = res.Where(fmt.Sprintf("status = ? AND customer_id = ?"), StatusFilter, CustomerFilter)
	}
	if StatusFilter == 0 && CustomerFilter != 0 {
		res = res.Where(fmt.Sprintf("customer_id = ?"), CustomerFilter)
	}
	if StatusFilter != 0 && CustomerFilter == 0 {
		res = res.Where(fmt.Sprintf("status = ?"), StatusFilter)
	}

	// get total count of collection with initial query
	res.Find(&orders).Count(&totalCount)

	// add page feature
	// if offset != 0 || count != 0 {
	// 	res = res.Offset(offset)
	// 	res = res.Limit(count)
	// }
	// add sort feature
	if sort != "" {
		if SortDirection == 1 {
			res = res.Order(sort)
		}
		if SortDirection == -1 {
			res = res.Order(sort + " desc")
		}
	}

	err := res.Find(&orders).Error

	return orders, totalCount, err
}

// ReadByJob returns user
func ReadByJob(jobs *model.JobsInfo) ([]*model.BookingInfo, int, error) {
	bookinginfos := []*model.BookingInfo{}
	orders := []*model.Order{}
	//	pros := []*model.Customer{}

	totalCount := 0
	patterns := []string{}
	values := []interface{}{}
	fields := []interface{}{}
	pattern := "%v = ?"
	likepattern := "%v LIKE ?"
	comma := " AND "

	res := db.ORM

	res = res.Table("orders").Select("orders.*, customers.post_code as post_code, customers.surname as customer_name, " +
		"services.service_name, devices.device_name,devices.image as device_image, makes.make_name, " +
		"models.model_name, customers.address as address, customers.address2 as address2").
		Joins("left join services on services.id = orders.service_id " +
			"left join devices on devices.id = orders.device_id left join customers on customers.id = orders.customer_id " +
			"left join makes on makes.id = orders.make left join models on models.id = orders.model")
	//fmt.Println(jobs.VendorID, jobs.ReferenceKey, jobs.CustomerName)
	if jobs.VendorID != 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "orders.vendor_id")
		values = append(values, jobs.VendorID)
	}

	if jobs.ReferenceKey != "" {
		patterns = append(patterns, likepattern)
		fields = append(fields, "orders.reference_key")
		values = append(values, jobs.ReferenceKey)
	}

	if jobs.CustomerName != "" {
		patterns = append(patterns, likepattern)
		fields = append(fields, "customers.surname")
		curStr := "%" + jobs.CustomerName + "%"
		values = append(values, curStr)
	}
	if jobs.PostCode != "" {
		patterns = append(patterns, likepattern)
		fields = append(fields, "post_code")
		curStr := "%" + jobs.PostCode + "%"
		values = append(values, curStr)
	}

	whereFormat := strings.Join(patterns, comma)

	whereStr := fmt.Sprintf(whereFormat, fields...)

	if len(whereFormat) != 0 {
		fmt.Println(whereStr, values)
		res = res.Where(whereStr, values...).Find(&orders)
	}

	fmt.Println("-----------------orders")

	fmt.Println(orders)
	customer := &model.Customer{}
	// Read Data

	for i := 0; i < len(orders); i++ {
		fmt.Println(orders[i])
		db.ORM.Table("customers").First(&customer, "id = ?", orders[i].CustomerID)
		bookinginfos = append(bookinginfos, &model.BookingInfo{Order: orders[i], Profile: customer})
	}
	fmt.Println(bookinginfos)
	//	res.Where("orders.vendor_id = ?", jobs.VendorID).Find(&bookinginfos)
	// add sort feature
	// if sort != "" {
	// 	if SortDirection == 1 {
	// 		res = res.Order(sort)
	// 	}
	// 	if SortDirection == -1 {
	// 		res = res.Order(sort + " desc")
	// 	}
	// }

	//	res.Find(&bookinginfos)

	return bookinginfos, totalCount, nil
}
