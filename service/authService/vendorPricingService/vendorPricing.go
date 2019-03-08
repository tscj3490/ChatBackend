package vendorPricingService

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
func CreateVendorPricing(vendorPricing *model.VendorPricing) (*model.VendorPricing, error) {
	// Insert Data
	if err := db.ORM.Create(&vendorPricing).Error; err != nil {
		return nil, err
	}
	return vendorPricing, nil
}

// ReadUser reads a user
func ReadVendorPricing(id uint) (*model.VendorPricing, error) {
	vendorPricing := &model.VendorPricing{}
	// Read Data
	err := db.ORM.First(&vendorPricing, "id = ?", id).Error
	return vendorPricing, err
}

// UpdateUser reads a user
func UpdateVendorPricing(vendorPricing *model.VendorPricing) (*model.VendorPricing, error) {
	res := db.ORM
	res = res.Table("vendor_pricing").Select("vendor_pricing.*, services.service_name as service_name, devices.device_name as device_name").
		Joins("left join services on services.id = vendor_pricing.service_id left join devices on devices.id = vendor_pricing.device_id")
	// Create change info
	err := db.ORM.Model(vendorPricing).Updates(vendorPricing).Error
	return vendorPricing, err
}

// DeleteUser deletes user with object id
func DeleteVendorPricing(id uint) error {
	vendorPricing := &model.VendorPricing{}

	err := db.ORM.Where("id = ?", id).Delete(vendorPricing).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadVendorsPricing(query string, offset int, count int, field string, sort int, vendorID uint) ([]*model.VendorPricing, int, error) {
	vendorsPricing := []*model.VendorPricing{}
	totalCount := 0

	res := db.ORM

	res = res.Table("vendor_pricing").Select("vendor_pricing.*, services.service_name as service_name, devices.device_name as device_name").
		Joins("left join services on services.id = vendor_pricing.service_id left join devices on devices.id = vendor_pricing.device_id")
	if vendorID != 0 {
		res = res.Where("vendor_pricing.vendor_id = ?", vendorID)
	}

	// get total count of collection with initial query
	res.Find(&vendorsPricing).Count(&totalCount)

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
	err := res.Find(&vendorsPricing).Error

	return vendorsPricing, totalCount, err
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, VendorID, DeviceID uint) ([]*model.VendorPricing, int, error) {
	vendorsPricing := []*model.VendorPricing{}
	totalCount := 0

	res := db.ORM

	// if vendor_id != 0 {
	// 	res = res.Where(fmt.Sprintf("vendor_id = %v", VendorID))
	// }
	//	if query != "" {
	//		query = "%" + query + "%"
	//		res = res.Where(fmt.Sprintf("%v LIKE ? OR %v LIKE ?", field, field), query, query)
	//	}
	// get total count of collection with initial query
	res = res.Table("vendor_pricing").Select("vendor_pricing.*, services.service_name as service_name").Joins("left join services on services.id = vendor_pricing.service_id")

	if VendorID != 0 && DeviceID != 0 {
		res = res.Where(fmt.Sprintf("vendor_pricing.vendor_id = ? AND vendor_pricing.device_id = ?"), VendorID, DeviceID)
	}
	if VendorID != 0 {
		res = res.Where(fmt.Sprintf("vendor_pricing.device_id = ?"), DeviceID)
	}
	if DeviceID != 0 {
		res = res.Where(fmt.Sprintf("vendor_pricing.vendor_id = ?"), VendorID)
	}
	res.Find(&vendorsPricing).Count(&totalCount)

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
	err := res.Find(&vendorsPricing).Error

	return vendorsPricing, totalCount, err
}
