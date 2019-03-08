package modelService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateModel creates a model
func CreateModel(model *model.Model) (*model.Model, error) {
	// Insert Data
	res := db.ORM

	if err := res.Create(&model).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Table("models").Select("models.*, makes.make_name, devices.device_name, devices.id as device_id, devices.image as device_image").
		Joins("left join makes on makes.id = models.parent_make left join devices on devices.id = makes.parent_device").
		Last(&model).Error
	return model, err
}

// ReadDevice reads a device
func ReadModel(id uint) (*model.Model, error) {
	model := &model.Model{}
	// Read Data

	err := db.ORM.Table("models").Select("models.*, makes.make_name, devices.device_name, devices.id as device_id, devices.image as device_image").
		Joins("left join makes on makes.id = models.parent_make left join devices on devices.id = makes.parent_device").Not("is_deleted", 1).
		First(&model, "models.id = ?", id).Error

	//err := db.ORM.First(&model, "id = ?", id).Error
	return model, err
}

// UpdateModel reads a device
func UpdateModel(model *model.Model) (*model.Model, error) {
	// Create change info
	err := db.ORM.Model(model).Updates(model).Error
	return model, err
}

// DeleteDevice deletes device with object id
func DeleteModel(id uint) error {

	models := []*model.Model{}

	res := db.ORM
	res.Table("services").Where("services.parent_device = ?", id).Find(&models).Update("is_deleted", 1)

	err := res.Delete(&model.Model{ID: id}).Error
	for _, m := range models {
		fmt.Println(m.ID)
	}

	return err
}

// ReadDevices return devices after retreive with params
func ReadModels(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Model, int, error) {
	models := []*model.Model{}
	totalCount := 0

	res := db.ORM

	//	res = res.Table("makes").Select("makes.*, devices.device_name").
	//		Joins("left join devices on devices.id = makes.parent_device")

	res = res.Table("models").Select("models.*, makes.make_name, devices.device_name, devices.id as device_id, devices.image as device_image").
		Joins("left join makes on makes.id = models.parent_make left join devices on devices.id = makes.parent_device").Not("is_deleted", 1)

	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("device_name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&models).Count(&totalCount)

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
	err := res.Find(&models).Error

	return models, totalCount, err
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Model, int, error) {
	models := []*model.Model{}
	totalCount := 0

	res := db.ORM

	if query != "" {
		//query = "%" + query + "%"
		res = res.Where(fmt.Sprintf("%v LIKE ?", field), query)
	}
	// get total count of collection with initial query
	res.Find(&models).Count(&totalCount)

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
	err := res.Find(&models).Error

	return models, totalCount, err
}
