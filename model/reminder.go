package model

import "time"

// Reminder is strucut for store position for every member
type Reminder struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CreatorID   string    `json:"creatorId" gorm:"creator_id"`
	Title       string    `json:"title" gorm:"title"`
	Description string    `json:"description" gorm:"description"`
	MeetingTime string    `json:"meeting_time" gorm:"meeting_time"`
	GroupID     string    `json:"groupId" gorm:"groupId"`
	Deleted     bool      `json:"deleted" gorm:"deleted"`
	CreatedAt   time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"updatedAt"`
}

// TableName indicates table name of user
func (Reminder) TableName() string {
	return "reminder"
}
