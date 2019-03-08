package model

import "time"

// OrderService ...
type OrderService struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	OrderID   uint      `json:"order_id" gorm:"order_id"`
	ServiceID uint      `json:"service_id" gorm:"service_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (OrderService) TableName() string {
	return "order_service"
}
