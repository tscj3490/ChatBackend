package model

import "time"

// Customers is a customers model
type VendorDevices struct {
	ID         uint      `json:"id" "primary_key"`
	VendorID   uint      `json:"vendor_id"  gorm:"column:vendor_id"`
	DeviceID   uint      `json:"device_id"  gorm:"column:device_id"`
	CreatedAt  time.Time `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at"   gorm:"column:updated_at"`
	DeviceName string    `json:"device_name"  gorm:"-"`
}

// TableName indicates table name of user
func (VendorDevices) TableName() string {
	return "vendor_devices"
}
