package model

import "time"

// Service ...
type Service struct {
	ID           uint      `json: "id" "primary_key" description:"Object Id"`
	ServiceName  string    `json:"service_name" gorm:"column:service_name"`
	VendorID     uint      `json:"vendor_id" gorm:"column:vendor_id"`
	ParentDevice uint      `json:"parent_device" gorm:"column:parent_device"`
	Description  string    `json:"description" gorm:"column:description"`
	Image        string    `json:"image" gorm:"column:image"`
	Price        int       `json:"price" gorm:"column:price"`
	IsDeleted    uint      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt    time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt"   gorm:"column:updated_at"`

	DeviceName  string `json:"device_name" gorm:"-"`
	VendorName  string `json:"vendor_name" gorm:"-"`
	DeviceImage string `json:"device_image" gorm:"-"`
}

// TableName indicates table name of Service
func (Service) TableName() string {
	return "services"
}
