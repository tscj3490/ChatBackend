package model

// Role ...
type Role struct {
	ID   uint   `json:"id" "primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

// TableName indicates table name of user
func (Role) TableName() string {
	return "roles"
}
