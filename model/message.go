package model

import "time"

// Message ...
type Message struct {
	ID         uint      `json: "id" gorm:"primary_key"`
	OrderID    uint      `json:"order_id" gorm:"column:order_id"`
	SenderID   uint      `json:"sender_id" gorm:"column:sender_id"`
	ReceiverID uint      `json:"receiver_id" gorm:"column:receiver_id"`
	Type       uint      `json:"type" gorm:"column:type"`
	Subject    string    `json:"subject" gorm:"column:subject"`
	Message    string    `json:"message" gorm:"column:message"`
	CreatedAt  time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

// TableName indicates table name of user
func (Message) TableName() string {
	return "messages"
}
