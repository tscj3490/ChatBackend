package model

import "time"

// Worktime ...
type Worktime struct {
	ID        uint      `json: "id" "primary_key" description:"Object Id"`
	VendorID  uint      `json:"vendor_id" gorm:"column:vendor_id"`
	DayOfWeek uint      `json:"day_of_week" gorm:"column:day_of_week"`
	StartTime int       `json:"start_time" gorm:"column:start_time"`
	CloseTime int       `json:"close_time" gorm:"column:close_time"`
	Type      uint      `json:"type" gorm:"column:type"`
	CreatedAt time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

type WorkTimeOrder struct {
	Worktime *Worktime `json:"worktime"`
	Order    *Order    `json:"order"`
}

// TableName indicates table name of Worktime
func (Worktime) TableName() string {
	return "worktimes"
}
