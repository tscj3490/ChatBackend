package model

import "time"

// Customers is a customers model
type Vendor struct {
	ID                 uint          `json:"id" "primary_key"`
	FirstName          string        `json:"firstname"  gorm:"column:firstname"`
	SurName            string        `json:"surname"  gorm:"column:surname"`
	Description        string        `json:"description"  gorm:"column:description"`
	Email              string        `json:"email"   gorm:"column:email"`
	UserName           string        `json:"username"  gorm:"column:username"`
	Company            string        `json:"company"  gorm:"column:company"`
	Password           string        `json:"password"  gorm:"column:password"`
	ProfilePic         string        `json:"profile_pic"   gorm:"column:profile_pic"`
	MobilePhone        string        `json:"mobile_phone" gorm:"column:mobile_phone"`
	TexPhone           string        `json:"tex_phone" gorm:"column:tex_phone"`
	Address            string        `json:"address"   gorm:"column:address"`
	Address2           string        `json:"address2"   gorm:"column:address2"`
	County             string        `json:"county"   gorm:"column:county"`
	PostCode           string        `json:"post_code"   gorm:"column:post_code"`
	PhoneCode          string        `json:"phone_code"   gorm:"column:phone_code"`
	City               string        `json:"city"   gorm:"column:city"`
	Type               int           `json:"type"   gorm:"column:type"`
	IsActive           uint          `json:"is_active" gorm:"column:is_active"`
	HasParking         uint          `json:"has_parking" gorm:"column:has_parking"`
	DropOff            uint          `json:"drop_off" gorm:"column:drop_off"`
	Collect            uint          `json:"collect" gorm:"column:collect"`
	CollectRange       uint          `json:"collect_range" gorm:"column:collect_range"`
	MaxorderDropoff    uint          `json:"maxorder_dropoff" gorm:"column:maxorder_dropoff"`
	MaxorderCollection uint          `json:"maxorder_collection" gorm:"column:maxorder_collection"`
	CoinAmount         float64       `json:"coin_amount" gorm:"column:coin_amount"`
	Rate               float32       `json:"rate" gorm:"column:rate"`
	Lat                float64       `json:"lat" gorm:"column:lat"`
	Lng                float64       `json:"lon" gorm:"column:lon"`
	CreatedAt          time.Time     `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt          time.Time     `json:"updated_at"   gorm:"column:updated_at"`
	IsOpen             int           `json:"is_open"  gorm:"-"`
	Distance           float64       `json:"distance" gorm:"-"`
	PaymentType        uint          `json:"payment_type"   gorm:"-"`
	DeviceList         []*DeviceInfo `gorm:"column:device_list"`
}

// PostVendor struct
type PostVendor struct {
	PostCode   string `json:"post_code" gorm:"column:post_code"`
	IsOpen     int    `json:"is_open" gorm:"-"`
	HasParking uint   `json:"has_parking" gorm:"column:has_parking"`
	DropOff    uint   `json:"drop_off" gorm:"column:drop_off"`
	Collect    uint   `json:"collect" gorm:"column:collect"`
	Sort       string `json:"sort" gorm:"column:sort"`
	DeviceID   []uint `json:"device_id" gorm:"column:device_id"`
}

//PublicVendors struct
type PublicVendor struct {
	*Vendor
	Password omit `json:"password,omitempty"`
}

//DeviceInfo struct
type DeviceInfo struct {
	DeviceID    uint   `json:"device_id"   gorm:"column:device_id"`
	DeviceName  string `json:"device_name"   gorm:"column:device_name"`
	DeviceImage string `json:"device_image"   gorm:"column:device_image"`
}

//WorktimeInfo struct
type WorktimeInfo struct {
	VendorID    uint           `json:"vendor_id"   gorm:"-"`
	OpenTime    []*Worktime    `json:"open_time"   gorm:"-"`
	SpecialDate []*Specialtime `json:"special_date"   gorm:"-"`
}

// SimpleVendors is struct for retrieve vendors
//PaymentType  -   0 : default(monthly)   1 : yearly    2 : 2 yearsly
type SimpleVendor struct {
	Username    string `json:"username"`
	Email       string `json:"email" description:"Vendor Email"`
	PaymentType uint   `json:"payment_type"   gorm:"-"`
}

// TableName indicates table name of user
func (Vendor) TableName() string {
	return "vendors"
}

type RegionInfo struct {
	Results []RegionResultsInfo `json:"results" gorm:"column:results"`
}

type RegionResultsInfo struct {
	Geometry GeometryInfo `json:"geometry" gorm:"column:geometry"`
}

type GeometryInfo struct {
	Location LocationInfo `json:"location" gorm:"column:location"`
}

type LocationInfo struct {
	Lat float64 `json:"lat" gorm:"column:lat"`
	Lng float64 `json:"lng" gorm:"column:lng"`
}
