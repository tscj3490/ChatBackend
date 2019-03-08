package model

import "time"

// UserSetting struct.
type UserSetting struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"userId" gorm:"user_id"`
	Code      int       `json:"code"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (UserSetting) TableName() string {
	return "tbl_user_setting"
}
