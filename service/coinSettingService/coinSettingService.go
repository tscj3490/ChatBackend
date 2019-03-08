package coinSettingService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateCoinSetting creates a coinSetting
func CreateCoinSetting(coinSetting *model.CoinSetting) (*model.CoinSetting, error) {
	// Insert Data
	if err := db.ORM.Create(&coinSetting).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&coinSetting).Error
	return coinSetting, err
}

// ReadCoinSetting reads a coinSetting
func ReadCoinSetting(id uint) (*model.CoinSetting, error) {
	coinSetting := &model.CoinSetting{}
	// Read Data
	err := db.ORM.First(&coinSetting, "id = ?", id).Error
	return coinSetting, err
}

// UpdateCoinSetting reads a coinSetting
func UpdateCoinSetting(coinSetting *model.CoinSetting) (*model.CoinSetting, error) {
	// Create change info
	err := db.ORM.Model(coinSetting).Updates(coinSetting).Error
	return coinSetting, err
}

// DeleteCoinSetting deletes coinSetting with object id
func DeleteCoinSetting(id uint) error {
	err := db.ORM.Delete(&model.CoinSetting{ID: id}).Error
	return err
}

// ReadCoinSettings return coinSettings after retreive with params
func ReadCoinSettings(query string, offset int, count int, field string, sort int, userID uint) ([]*model.CoinSetting, int, error) {
	coinSettings := []*model.CoinSetting{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("user_id = ?", userID)
	}
	// if query != "" {
	// 	query = "%" + query + "%"
	// 	res = res.Where("license_number LIKE ? OR description LIKE ?", query, query)
	// }
	// get total count of collection with initial query
	res.Find(&coinSettings).Count(&totalCount)

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
	err := res.Find(&coinSettings).Error

	return coinSettings, totalCount, err
}
