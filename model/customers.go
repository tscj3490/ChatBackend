package model

import "time"

// Customer is a customers model
type Customer struct {
	ID          uint      `json:"id" gorm:"column:id"`
	FirstName   string    `json:"firstname"  gorm:"column:firstname"`
	SurName     string    `json:"surname"  gorm:"column:surname"`
	UserName    string    `json:"username"  gorm:"column:username"`
	Address     string    `json:"address"   gorm:"column:address"`
	Address2    string    `json:"address2"   gorm:"column:address2"`
	MobilePhone string    `json:"mobile_phone" gorm:"column:mobile_phone"`
	TexPhone    string    `json:"tex_phone" gorm:"column:tex_phone"`
	Email       string    `json:"email"   gorm:"column:email"`
	Password    string    `json:"password"   gorm:"column:password"`
	City        string    `json:"city"   gorm:"column:city"`
	County      string    `json:"county"   gorm:"column:county"`
	PostCode    string    `json:"post_code"   gorm:"column:post_code"`
	PhoneCode   string    `json:"phone_code"   gorm:"column:phone_code"`
	ProfilePic  string    `json:"profile_pic"   gorm:"column:profile_pic"`
	Rate        float32   `json:"rate" gorm:"column:rate"`
	CoinAmount  float32   `json:"coin_amount" gorm:"column:coin_amount"`
	CreatedAt   time.Time `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

//PublicCustomers struct
type PublicCustomer struct {
	*Customer
	Password omit `json:"password,omitempty"`
}

//PublicBook struct
type PublicBook struct {
	*Customer
	ReferenceKey omit `json:"reference_key,omitempty"`
}

// SimpleUser is struct for retrieve user
type SimpleCustomer struct {
	Username string `json:"username"`
}

// TableName indicates table name of user
func (Customer) TableName() string {
	return "customers"
}
