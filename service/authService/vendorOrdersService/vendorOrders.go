package vendorOrdersService

import (
	"errors"
	"fmt"
	"time"

	"../../../db"
	"../../../model"
	"../../../util/timeHelper"
	//	"../../../util/crypto"
)

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateVendorOrders(vendorOrders *model.VendorOrders) (*model.VendorOrders, error) {
	// check duplicate username
	v := &model.VendorOrders{}
	if res := db.ORM.Where("vendor_id = ?", vendorOrders.VendorID).First(&v).RecordNotFound(); !res {
		err := errors.New(fmt.Sprintf("%v", vendorOrders.VendorID) + " is already registered")
		return nil, err
	}

	// Insert Data
	if err := db.ORM.Create(&vendorOrders).Error; err != nil {
		return nil, err
	}
	return vendorOrders, nil
}

// ReadUser reads a user
func ReadVendorOrders(id uint) (*model.VendorOrders, error) {
	vendorOrders := &model.VendorOrders{}
	// Read Data
	err := db.ORM.First(&vendorOrders, "id = ?", id).Error
	return vendorOrders, err
}

// UpdateUser reads a user
func UpdateVendorOrders(vendorOrders *model.VendorOrders) (*model.VendorOrders, error) {
	// Create change info
	err := db.ORM.Model(vendorOrders).Updates(vendorOrders).Error
	return vendorOrders, err
}

// DeleteUser deletes user with object id
func DeleteVendorOrders(id uint) error {
	vendorOrders := &model.VendorOrders{}

	err := db.ORM.Where("id = ?", id).Delete(vendorOrders).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadVendorsOrders(query string, offset int, count int, field string, sort int, vendorID uint) ([]*model.VendorOrders, int, error) {
	vendorsOrders := []*model.VendorOrders{}
	totalCount := 0

	res := db.ORM

	if vendorID != 0 {
		res = res.Where("vendor_id = ?", vendorID)
	}
	if query != "" {
		query = "%" + query + "%"
		//		res = res.Where("license_number LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&vendorsOrders).Count(&totalCount)

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
	err := res.Find(&vendorsOrders).Error

	return vendorsOrders, totalCount, err
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, vendor_id uint) ([]*model.VendorOrders, int, error) {
	vendorsOrders := []*model.VendorOrders{}
	totalCount := 0

	res := db.ORM
	if vendor_id != 0 {
		res = res.Where(fmt.Sprintf("vendor_id = %v", vendor_id))
	}
	//	if query != "" {
	//		query = "%" + query + "%"
	//		res = res.Where(fmt.Sprintf("%v LIKE ? OR %v LIKE ?", field, field), query, query)
	//	}
	// get total count of collection with initial query
	res.Find(&vendorsOrders).Count(&totalCount)

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
	err := res.Find(&vendorsOrders).Error

	return vendorsOrders, totalCount, err
}

// SetVendorOrder
func SetVendorOrder(vendor *model.Vendor) (*model.VendorOrders, error) {
	vo := &model.VendorOrders{}
	vo.VendorID = vendor.ID
	vo.PlanType = vendor.PaymentType
	vo.OrderDate = time.Now()
	switch vo.PlanType {
	case 0:
		vo.RenewDate = timeHelper.FewDaysLaterDate(vo.OrderDate, 30)
	case 1:
		vo.RenewDate = timeHelper.FewDaysLaterDate(vo.OrderDate, 365)
	case 2:
		vo.RenewDate = timeHelper.FewDaysLaterDate(vo.OrderDate, 365*2)
	default:
		vo.RenewDate = timeHelper.FewDaysLaterDate(vo.OrderDate, 30)
	}
	vo.Status = 1 //completed
	fmt.Println(vo.PlanType, vo.OrderDate, vo.RenewDate)
	db.ORM.NewRecord(&vo)
	err := db.ORM.Create(vo).Error

	return vo, err
}
