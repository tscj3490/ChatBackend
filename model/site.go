package model

import "time"

// Site is strucut for store position for every member
type Site struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UserID      string    `json:"userId" gorm:"user_id"`
	Name        string    `json:"name"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Type        int       `json:"type"`
	Description string    `json:"description"`
	Lat         float64   `json:"lat"`
	Lng         float64   `json:"lng"`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (Site) TableName() string {
	return "tbl_site"
}
