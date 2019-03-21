package model

import "time"

// Admin model
type Admin struct {
	ID        uint      `json:"id" gorm:"primary_key" description:"Object ID"`
	Name      string    `json:"name" description:"Name"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role"`
	Enabled   bool      `json:"enabled"`
	Deleted   bool      `json:"deleted"`
	LastIP    string    `json:"lastIP"`
	CreatedAt time.Time `json:"createdAt" description:"Created date"`
	UpdatedAt time.Time `json:"updatedAt" description:"Updated date. This field will be updated when any update operation will be occured"`
}

// PublicAdmin struct.
type PublicAdmin struct {
	*Admin
	Password omit `json:"password,omitempty"`
}

// ChangePass struct
type ChangePass struct {
	// Role    string `json:"role" gorm:"-"`
	Email   string `json:"email" gorm:"-"`
	OldPass string `json:"old_password" gorm:"-"`
	NewPass string `json:"new_password" gorm:"-"`
}

// TableName indicates table name of user
func (Admin) TableName() string {
	return "admins"
}
