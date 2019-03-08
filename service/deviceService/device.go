package deviceService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateDevice creates a device
func CreateDevice(device *model.Device) (*model.Device, error) {
	// Insert Data
	if err := db.ORM.Create(&device).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&device).Error
	return device, err
}

// ReadDevice reads a device
func ReadDevice(id uint) (*model.Device, error) {
	device := &model.Device{}
	// Read Data
	err := db.ORM.First(&device, "id = ?", id).Error
	return device, err
}

// UpdateDevice reads a device
func UpdateDevice(device *model.Device) (*model.Device, error) {
	// Create change info
	err := db.ORM.Model(device).Updates(device).Error
	return device, err
}

// DeleteDevice deletes device with object id
func DeleteDevice(id uint) error {
	devices := []*model.Device{}

	res := db.ORM
	res.Table("devices").Where("id = ?", id).Update("is_deleted", 1)
	res.Table("makes").Where("makes.parent_device = ?", id).Find(&devices).Update("is_deleted", 1)

	fmt.Println("makes_id : ", id)
	err := res.Delete(&model.Device{ID: id}).Error
	for _, m := range devices {
		fmt.Println(m.ID)
	}
	return err
}

// ReadDevices return devices after retreive with params
func ReadDevices(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Device, int, error) {
	devices := []*model.Device{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("device_name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&devices).Count(&totalCount)

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
	err := res.Find(&devices).Error

	return devices, totalCount, err
}
