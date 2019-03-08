package model

import "time"

// Joblog ...
type Joblog struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	OrderID    uint      `json:"order_id" gorm:"order_id"`
	CustomerID uint      `json:"customer_id" gorm:"customer_id"`
	VendorID   uint      `json:"vendor_id" gorm:"vendor_id"`
	Status     uint      `json:"status" gorm:"column:status"`
	Price      uint      `json:"price" gorm:"column:price"`
	Comment    string    `json:"comment" gorm:"column:comment"`
	ByEmail    uint      `json:"byemail" gorm:"column:byemail"`
	BySms      uint      `json:"bysms" gorm:"column:bysms"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (Joblog) TableName() string {
	return "joblogs"
}
