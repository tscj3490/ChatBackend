package paymentSettingService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreatePaymentSetting creates a paymentSetting
func CreatePaymentSetting(paymentSetting *model.PaymentSetting) (*model.PaymentSetting, error) {
	// Insert Data
	if err := db.ORM.Create(&paymentSetting).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&paymentSetting).Error
	return paymentSetting, err
}

// ReadPaymentSetting reads a paymentSetting
func ReadPaymentSetting(id uint) (*model.PaymentSetting, error) {
	paymentSetting := &model.PaymentSetting{}
	// Read Data
	err := db.ORM.First(&paymentSetting, "id = ?", id).Error
	return paymentSetting, err
}

// UpdatePaymentSetting reads a paymentSetting
func UpdatePaymentSetting(paymentSetting *model.PaymentSetting) (*model.PaymentSetting, error) {
	// Create change info
	err := db.ORM.Model(paymentSetting).Updates(paymentSetting).Error
	return paymentSetting, err
}

// DeletePaymentSetting deletes paymentSetting with object id
func DeletePaymentSetting(id uint) error {
	err := db.ORM.Delete(&model.PaymentSetting{ID: id}).Error
	return err
}

// ReadPaymentSettings return paymentSettings after retreive with params
func ReadPaymentSettings(query string, offset int, count int, field string, sort int, userID uint) ([]*model.PaymentSetting, int, error) {
	paymentSettings := []*model.PaymentSetting{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("userID = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("license_number LIKE ? OR location LIKE ? OR note LIKE ?", query, query, query)
	}
	// get total count of collection with initial query
	res.Find(&paymentSettings).Count(&totalCount)

	// add page feature
	if offset != 0 || count != 0 {
		res = res.Offset(offset)
		res = res.Limit(count)
	}
	// add sort feature
	if field != "" && sort != 0 {
		if sort > 0 {
			res = res.Order(field)
		} else {
			res = res.Order(field + " desc")
		}
	}
	err := res.Find(&paymentSettings).Error

	return paymentSettings, totalCount, err
}
