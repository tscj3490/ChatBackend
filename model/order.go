package model

import "time"

// Order ...
type Order struct {
	ID           uint   `json: "id" "primary_key" description:"Object Id"`
	VendorID     uint   `json:"vendor_id" gorm:"column:vendor_id"`
	CustomerID   uint   `json:"customer_id" gorm:"column:customer_id"`
	ServiceID    uint   `json:"service_id" gorm:"column:service_id"`
	Description  string `json:"description" gorm:"column:description"`
	Status       uint   `json:"status" gorm:"column:status"`
	Type         int    `json:"type"   gorm:"column:type"`
	ByVendor     int    `json:"byvendor"   gorm:"column:byvendor"`
	ReferenceKey string `json:"reference_key" gorm:"column:reference_key"`
	Price        uint   `json:"price" gorm:"column:price"`
	IsOpen       uint   `json:"is_open" gorm:"column:is_open"`
	DeviceID     uint   `json:"device_id" gorm:"column:device_id"`
	Make         uint   `json:"make" gorm:"column:make"`
	Model        uint   `json:"model" gorm:"column:model"`
	//	Service     uint      `json:"service" gorm:"column:service"`
	BookDate  time.Time `json:"book_date"   gorm:"column:book_date"`
	CreatedAt time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt"   gorm:"column:updated_at"`

	Company      string `json:"company" gorm:"-"`
	ServiceName  string `json:"servicename" gorm:"-"`
	DeviceName   string `json:"devicename" gorm:"-"`
	VendorName   string `json:"vendorname" gorm:"-"`
	CustomerName string `json:"customername" gorm:"-"`
	MakeName     string `json:"makename" gorm:"-"`
	ModelName    string `json:"modelname" gorm:"-"`
	DeviceImage  string `json:"deviceimage" gorm:"-"`
	Address      string `json:"address" gorm:"-"`
	Address2     string `json:"address2" gorm:"-"`
	PostCode     string `json:"post_code" gorm:"-"`
	ServiceList  []uint `json:"servicelist" gorm:"-"`
}

//BookingInfo ...
type BookingInfo struct {
	Profile *Customer `json:"profile"`
	Order   *Order    `json:"order"`
}

// Job search ...
type JobsInfo struct {
	VendorID     uint   `json:"vendor_id" gorm:"-"`
	ReferenceKey string `json:"reference_key" gorm:"-"`
	CustomerName string `json:"customer_name" gorm:"-"`
	PostCode     string `json:"post_code" gorm:"-"`
}

// UpdateStatus ...
//EmailSms : 0 - not send  1 - email sending   2 - sms sending
type UpdateStatus struct {
	VendorID     uint   `json:"vendor_id" gorm:"-"`
	OrderID      uint   `json:"order_id" gorm:"-"`
	CustomerID   uint   `json:"customer_id" gorm:"-"`
	Status       uint   `json:"status" gorm:"-"`
	Price        uint   `json:"price" gorm:"-"`
	JobsComments string `json:"jobs_comments" gorm:"-"`
	ByEmail      uint   `json:"byemail" gorm:"-"`
	BySms        uint   `json:"bysms" gorm:"-"`
}

// TableName indicates table name of Order
func (Order) TableName() string {
	return "orders"
}

// TimeInfo ...
type TimeInfo struct {
	OrderInfo       []*Order       `gorm:"column:order_info"`
	WorktimeInfo    []*Worktime    `gorm:"column:worktime_info"`
	SpecialtimeInfo []*Specialtime `gorm:"column:specialtime_info"`
}
