package model

// Company ...
type Company struct {
	ID   uint   `json:"id" "primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

// TableName indicates table name of user
func (Company) TableName() string {
	return "companies"
}
