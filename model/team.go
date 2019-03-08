package model

// Team ...
type Team struct {
	ID        uint   `json:"id" "primary_key"`
	Name      string `json:"name" gorm:"column:name"`
	CompanyID int64  `json:"company_id" gorm:"column:company_id"`

	CompanyName string `json:"company_name" gorm:"-"`
}

// TableName indicates table name of user
func (Team) TableName() string {
	return "teams"
}
