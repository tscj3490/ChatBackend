package makeService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateMake creates a make
func CreateMake(make *model.Make) (*model.Make, error) {
	// Insert Data
	res := db.ORM
	if err := res.Create(&make).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Table("makes").Select("makes.*, devices.device_name, devices.image as device_image").
		Joins("left join devices on devices.id = makes.parent_device").
		Last(&make).Error
	return make, err
}

// ReadMake reads a make
func ReadMake(id uint) (*model.Make, error) {
	make := &model.Make{}
	// Read Data
	err := db.ORM.Table("makes").Select("makes.*, devices.device_name, devices.image as device_image").
		Joins("left join devices on devices.id = makes.parent_device").Not("is_deleted", 1).
		First(&make, "makes.id = ?", id).Error
	//	err := db.ORM.First(&make, "id = ?", id).Error
	return make, err
}

// UpdateMake reads a device
func UpdateMake(make *model.Make) (*model.Make, error) {
	// Create change info
	err := db.ORM.Model(make).Updates(make).Error
	return make, err
}

// DeleteMake deletes make with object id
func DeleteMake(id uint) error {

	makes := []*model.Make{}

	res := db.ORM
	res.Table("models").Where("models.parent_make = ?", id).Find(&makes).Update("is_deleted", 1)

	fmt.Println("models_id : ", id)
	err := res.Delete(&model.Make{ID: id}).Error

	for _, m := range makes {
		fmt.Println(m.ID)
	}

	return err
}

// ReadDevices return devices after retreive with params
func ReadMakes(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Make, int, error) {
	makes := []*model.Make{}
	totalCount := 0

	res := db.ORM

	res = res.Table("makes").Select("makes.*, devices.device_name, devices.image, devices.image as device_image").
		Joins("left join devices on devices.id = makes.parent_device").Not("is_deleted", 1)

	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("make_name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&makes).Count(&totalCount)

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
	err := res.Find(&makes).Error

	return makes, totalCount, err
}

// get list by parent_id
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Make, int, error) {
	makes := []*model.Make{}
	totalCount := 0

	res := db.ORM

	if query != "" {
		//query = "%" + query + "%"
		res = res.Where(fmt.Sprintf("%v LIKE ?", field), query)
	}
	// get total count of collection with initial query
	res.Find(&makes).Count(&totalCount)

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
	err := res.Find(&makes).Error

	return makes, totalCount, err
}
