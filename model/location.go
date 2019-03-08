package model

import "time"

// Location ...
type Location struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	UserNumber uint      `json:"userNumber" gorm:"user_number"`
	DeviceID   uint      `json:"deviceId" gorm:"device_id"`
	Lat        float64   `json:"lat"`
	Lng        float64   `json:"lng"`
	Deleted    bool      `json:"deleted"`
	SentAt     time.Time `json:"sentAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (Location) TableName() string {
	return "tbl_location"
}
