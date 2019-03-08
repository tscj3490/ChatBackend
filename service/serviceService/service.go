package serviceService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateService creates a service
func CreateService(service *model.Service) (*model.Service, error) {
	// Insert Data
	res := db.ORM

	if err := res.Create(&service).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Table("services").Select("services.*, devices.device_name, vendors.username as vendor_name, devices.image as device_image").
		Joins("left join devices on devices.id = services.parent_device left join vendors on vendors.id = services.vendor_id").
		Last(&service).Error
	return service, err
}

// ReadService reads a service
func ReadService(id uint) (*model.Service, error) {
	service := &model.Service{}

	// Read Data
	err := db.ORM.Table("services").Select("services.*, devices.device_name, vendors.username as vendor_name, devices.image as device_image").
		Joins("left join devices on devices.id = services.parent_device left join vendors on vendors.id = services.vendor_id").Not("is_deleted", 1).
		First(&service, "services.id = ?", id).Error

	return service, err
}

// UpdateReview reads a service
func UpdateService(service *model.Service) (*model.Service, error) {
	// Create change info

	err := db.ORM.Model(service).Updates(service).Error
	return service, err
}

// DeleteService deletes service with object id
func DeleteService(id uint) error {
	err := db.ORM.Delete(&model.Service{ID: id}).Error
	return err
}

// ReadServices return reviews after retreive with params
func ReadServices(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Service, int, error) {
	services := []*model.Service{}
	totalCount := 0

	res := db.ORM

	res = res.Table("services").Select("services.*, devices.device_name, vendors.username as vendor_name, devices.image as device_image").
		Joins("left join devices on devices.id = services.parent_device left join vendors on vendors.id = services.vendor_id").Not("is_deleted", 1)

	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("service_name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&services).Count(&totalCount)

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
	err := res.Find(&services).Error

	return services, totalCount, err
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Service, int, error) {
	services := []*model.Service{}
	totalCount := 0

	res := db.ORM

	if query != "" {
		//query = "%" + query + "%"
		res = res.Where(fmt.Sprintf("%v LIKE ?", field), query)
	}
	// get total count of collection with initial query
	res.Find(&services).Count(&totalCount)

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
	err := res.Find(&services).Error

	return services, totalCount, err
}
