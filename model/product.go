package model

// Blacklist ...
type Product struct {
	ID          int64  `json:"ID" gorm:"column:ID"`
	ImageURL    string `json:"imgUrl" gorm:"column:imgUrl"`
	Product     string `json:"Product" gorm:"column:Product"`
	Description string `json:"Description" gorm:"column:Description"`
}

// TableName indicates table name of user
func (Product) TableName() string {
	return "products"
}
