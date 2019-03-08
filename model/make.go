package model

import "time"

// Make ...
type Make struct {
	ID           uint      `json: "id" "primary_key" description:"Object Id"`
	MakeName     string    `json:"make_name" gorm:"column:make_name"`
	ParentDevice uint      `json:"parent_device" gorm:"column:parent_device"`
	Description  string    `json:"description" gorm:"column:description"`
	Image        string    `json:"image" gorm:"column:image"`
	IsDeleted    uint      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt    time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt"   gorm:"column:updated_at"`

	DeviceName  string `json:"device_name" gorm:"-"`
	DeviceImage string `json:"device_image" gorm:"-"`
}

// TableName indicates table name of user
func (Make) TableName() string {
	return "makes"
}
