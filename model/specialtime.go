package model

import "time"

// Specialtime ...
// kind   1 - not talking    2 -  close day
// type   1 - drop           2 -  collect
type Specialtime struct {
	ID        uint      `json: "id" "primary_key" description:"Object Id"`
	Type      int       `json:"type" gorm:"column:type"`
	Kind      int       `json:"kind" gorm:"column:kind"`
	Date      time.Time `json:"date" gorm:"column:date"`
	HourStart uint      `json:"hour_start" gorm:"column:hour_start"`
	HourEnd   uint      `json:"hour_end" gorm:"column:hour_end"`
	TimeRange uint      `json:"time_range" gorm:"column:time_range"`
	VendorID  uint      `json:"vendor_id" gorm:"column:vendor_id"`
	CreatedAt time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

// TableName indicates table name of Worktime
func (Specialtime) TableName() string {
	return "specialtimes"
}
