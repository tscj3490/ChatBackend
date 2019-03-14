package model

// Team ...
type Team struct {
	ID          uint   `json:"id" "primary_key"`
	Name        string `json:"name" gorm:"column:name"`
	CompanyID   int64  `json:"company_id" gorm:"column:company_id"`
	Logo        string `json:"logo" gorm:"column:logo"`
	CompanyName string `json:"company_name" gorm:"column:company_name"`
}

// TableName indicates table name of user
func (Team) TableName() string {
	return "teams"
}
