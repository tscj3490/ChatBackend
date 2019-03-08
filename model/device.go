package model

import "time"

// Device ...
type Device struct {
	ID          uint      `json: "id" "primary_key" description:"Object Id"`
	DeviceName  string    `json:"device_name" gorm:"column:device_name"`
	Description string    `json:"description" gorm:"column:description"`
	Image       string    `json:"image" gorm:"column:image"`
	IsDeleted   uint      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt   time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

// TableName indicates table name of user
func (Device) TableName() string {
	return "devices"
}
