package model

import "time"

// Customers is a customers model
// PlanType : 1 - after a month  2 - after a year  3 - after 2 year
// Status :   1 - completed      2 - pending       3 - fail
type VendorOrders struct {
	ID        uint      `json:"id" "primary_key"`
	VendorID  uint      `json:"vendor_id"  gorm:"column:vendor_id"`
	PlanType  uint      `json:"plan_type"  gorm:"column:plan_type"`
	OrderDate time.Time `json:"order_date"   gorm:"column:order_date"`
	RenewDate time.Time `json:"renew_date"   gorm:"column:renew_date"`
	Status    int       `json:"status"  gorm:"column:status"`
	CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

// TableName indicates table name of user
func (VendorOrders) TableName() string {
	return "vendor_orders"
}
