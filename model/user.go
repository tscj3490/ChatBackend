package model

import "time"

// User is a user model
type User struct {
	ID         uint      `json:"id" gorm:"column:id"`
	Name       string    `json:"name"  gorm:"column:name"`
	Email      string    `json:"email"  gorm:"column:email"`
	Phone      string    `json:"phone"   gorm:"column:phone"`
	Avatar     string    `json:"avatar" gorm:"column:avatar"`
	Role       string    `json:"role" gorm:"column:role"`
	TeamID     uint      `json:"team_id" gorm:"column:team_id"`
	Code       string    `json:"code" gorm:"column:code"`
	IsVerified bool      `json:"is_verified" gorm:"column:is_verified"`
	Token      string    `json:"token" gorm:"column:token"`
	Deleted    bool      `json:"deleted"   gorm:"column:deleted"`
	CreatedAt  time.Time `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at"   gorm:"column:updated_at"`

	TeamName string `json:"team_name" gorm:"-"`
}

// ManagerInfo
type ManagerInfo struct {
	Team  *Team  `json:"team"   gorm:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// PublicUser struct.
type PublicUser struct {
	*User
}

// SimpleUser is struct for retrieve user
type SimpleUser struct {
	UserName      string `json:"UserName"`
	UserFirstName string `json:"UserFirstName" description:"User firstname"`
	UserLastName  string `json:"UserLastName" description:"User lastname"`
}

// TableName indicates table name of user
func (User) TableName() string {
	return "users"
}
