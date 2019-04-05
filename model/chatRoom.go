package model

import "time"

// ChatRoom is strucut for store position for every member
type ChatRoom struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"name"`
	UserIDs   string    `json:"userIds" gorm:"user_ids"`
	Deleted   bool      `json:"deleted" gorm:"deleted"`
	CreatedAt time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updatedAt"`
}

// TableName indicates table name of user
func (ChatRoom) TableName() string {
	return "chat_room"
}
