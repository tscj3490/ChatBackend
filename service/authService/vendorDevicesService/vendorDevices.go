package vendorDevicesService

import (
	"fmt"

	"../../../db"
	"../../../model"
	//	"../../../util/crypto"
)

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateVendorDevices(vendorDevices *model.VendorDevices) (*model.VendorDevices, error) {
	// check duplicate username
	// v := &model.VendorDevices{}
	// if res := db.ORM.Where("vendor_id = ?", vendorDevices.VendorID).First(&v).RecordNotFound(); !res {
	// 	err := errors.New(fmt.Sprintf("%v", vendorDevices.VendorID) + " is already registered")
	// 	return nil, err
	// }

	// Insert Data
	if err := db.ORM.Create(&vendorDevices).Error; err != nil {
		return nil, err
	}
	return vendorDevices, nil
}

// ReadUser reads a user
func ReadVendorDevices(id uint) (*model.VendorDevices, error) {
	vendorDevices := &model.VendorDevices{}
	// Read Data
	err := db.ORM.First(&vendorDevices, "id = ?", id).Error
	return vendorDevices, err
}

// UpdateUser reads a user
func UpdateVendorDevices(vendorDevices *model.VendorDevices) (*model.VendorDevices, error) {
	// Create change info
	err := db.ORM.Model(vendorDevices).Updates(vendorDevices).Error
	return vendorDevices, err
}

// DeleteUser deletes user with object id
func DeleteVendorDevices(id uint) error {
	vendorDevices := &model.VendorDevices{}

	err := db.ORM.Where("id = ?", id).Delete(vendorDevices).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadVendorsDevices(query string, offset int, count int, field string, sort int, vendorID uint) ([]*model.VendorDevices, int, error) {
	vendorsDevices := []*model.VendorDevices{}
	totalCount := 0

	res := db.ORM

	res = res.Table("vendor_devices").Select("vendor_devices.*, devices.device_name as device_name").
		Joins("left join devices on vendor_devices.device_id = devices.id")
	if vendorID != 0 {
		res = res.Where("vendor_devices.vendor_id = ?", vendorID)
	}
	if query != "" {
		query = "%" + query + "%"
		//		res = res.Where("license_number LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&vendorsDevices).Count(&totalCount)

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
	err := res.Find(&vendorsDevices).Error

	return vendorsDevices, totalCount, err
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, vendor_id uint) ([]*model.VendorDevices, int, error) {
	vendorsDevices := []*model.VendorDevices{}
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
	res.Find(&vendorsDevices).Count(&totalCount)

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
	err := res.Find(&vendorsDevices).Error

	return vendorsDevices, totalCount, err
}
