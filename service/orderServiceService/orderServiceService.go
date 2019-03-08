package orderServiceService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateOrderService creates a orderService
func CreateOrderService(orderService *model.OrderService) (*model.OrderService, error) {
	// Insert Data
	if err := db.ORM.Create(&orderService).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&orderService).Error
	return orderService, err
}

// ReadOrderService reads a orderService
func ReadOrderService(id uint) (*model.OrderService, error) {
	orderService := &model.OrderService{}
	// Read Data
	err := db.ORM.First(&orderService, "id = ?", id).Error
	return orderService, err
}

// UpdateOrderService reads a orderService
func UpdateOrderService(orderService *model.OrderService) (*model.OrderService, error) {
	// Create change info
	err := db.ORM.Model(orderService).Updates(orderService).Error
	return orderService, err
}

// DeleteOrderService deletes orderService with object id
func DeleteOrderService(id uint) error {
	err := db.ORM.Delete(&model.OrderService{ID: id}).Error
	return err
}

// ReadOrderServices return orderServices after retreive with params
func ReadOrderServices(query string, offset int, count int, field string, sort int, orderID uint) ([]*model.Order, int, error) {
	orderServices := []*model.OrderService{}
	totalCount := 0

	order := []*model.Order{}

	res := db.ORM

	res = db.ORM.Table("orders")
	res = res.Select("orders.*, vendors.company as company, services.service_name, devices.device_name, vendors.username as vendor_name, " +
		"customers.username as customer_name, devices.image as device_image, makes.make_name, models.model_name, customers.address as address, " +
		"customers.address2 as address2, customers.post_code as post_code").
		Joins("left join vendors on vendors.id = orders.vendor_id left join services on services.id = orders.service_id " +
			"left join devices on devices.id = orders.device_id left join customers on customers.id = orders.customer_id " +
			"left join makes on makes.id = orders.make left join models on models.id = orders.model")
	if query != "" {
		res = res.Where(fmt.Sprintf("%v = ?", "orders."+field), query)
	}
	res.Find(&order)

	//	res = res.Where("reference_key = ?", "ZJGA-AVH4-A8YP-AZN8")
	res.Find(&order).Count(&totalCount)
	servicelist := []uint{}
	for _, o := range order {
		db.ORM.Where("service_id = ?", o.ID).Find(&orderServices)
		for _, v := range orderServices {
			servicelist = append(servicelist, v.ID)
		}
		o.ServiceList = servicelist
		servicelist = []uint{}
	}

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
	//	err := res.Find(&orderServices).Error

	return order, totalCount, nil
}
