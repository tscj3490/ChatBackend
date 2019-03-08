package model

import "time"

// Review ...
type Review struct {
	ID         uint      `json: "id" "primary_key" description:"Object Id"`
	VendorID   uint      `json:"vendor_id" gorm:"column:vendor_id"`
	CustomerID uint      `json:"customer_id" gorm:"column:customer_id"`
	OrderID    uint      `json:"order_id" gorm:"column:order_id"`
	Type       uint      `json:"type" gorm:"column:type"`
	Score      uint      `json:"score" gorm:"column:score"`
	Message    string    `json:"message"   gorm:"column:message"`
	CreatedAt  time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

// TableName indicates table name of Review
func (Review) TableName() string {
	return "reviews"
}
