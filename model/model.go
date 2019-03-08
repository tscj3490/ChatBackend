package model

import "time"

// Model ...
type Model struct {
	ID          uint      `json: "id" "primary_key" description:"Object Id"`
	ModelName   string    `json:"model_name" gorm:"column:model_name"`
	ParentMake  uint      `json:"parent_make" gorm:"column:parent_make"`
	Description string    `json:"description" gorm:"column:description"`
	Image       string    `json:"image" gorm:"column:image"`
	IsDeleted   uint      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt   time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt"   gorm:"column:updated_at"`

	MakeName    string `json:"make_name"  gorm:"-"`
	DeviceName  string `json:"device_name"  gorm:"-"`
	DeviceID    uint   `json:"device_id"  gorm:"-"`
	DeviceImage string `json:"device_image" gorm:"-"`
}

// TableName indicates table name of user
func (Model) TableName() string {
	return "models"
}
