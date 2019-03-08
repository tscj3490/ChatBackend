package model

// Tblregion ...

type Tblregion struct {
	ID            uint   `json: "id" "primary_key" description:"Object Id"`
	PostCode      string `json:"postcode" gorm:"column:postcode"`
	Eastings      int    `json:"eastings" gorm:"column:eastings"`
	Nothings      int    `json:"nothings" gorm:"column:nothings"`
	Latitude      string `json:"latitude" gorm:"column:latitude"`
	Longitude     string `json:"longitude" gorm:"column:longitude"`
	Town          string `json:"town" gorm:"column:town"`
	Region        string `json:"region" gorm:"column:region"`
	County        string `json:"county" gorm:"column:county"`
	CountryString string `json:"country_string" gorm:"column:country_string"`
}

// TableName indicates table name of Worktime
func (Tblregion) TableName() string {
	return "tblregions"
}
