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

// SendMessageInfo
type SendMessageInfo struct {
	GroupID uint `json:"groupId"`
	Message *Msg `json:"message"`
}

// Msg
type Msg struct {
	Text      string    `json:"text"`
	Image     string    `json:"image"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"createdAt"`
}

// TableName indicates table name of user
func (ChatRoom) TableName() string {
	return "chat_room"
}
