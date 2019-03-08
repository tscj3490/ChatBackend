package model

import "time"

// Transaction ...
type Transaction struct {
	ID        uint      `json: "id" "primary_key" description:"Object Id"`
	VendorID  uint      `json:"vendor_id" gorm:"column:vendor_id"`
	PlanType  uint      `json:"plan_type" gorm:"column:plan_type"`
	Status    uint      `json:"status" gorm:"column:status"`
	Price     float32   `json:"price" gorm:"column:price"`
	CreatedAt time.Time `json:"createdAt"   gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt"   gorm:"column:updated_at"`
}

// TableName indicates table name of Transaction
func (Transaction) TableName() string {
	return "transactions"
}
