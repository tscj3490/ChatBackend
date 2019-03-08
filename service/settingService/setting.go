package settingService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateSetting creates a setting
func CreateSetting(setting *model.Setting) (*model.Setting, error) {
	// Insert Data
	if err := db.ORM.Create(&setting).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&setting).Error
	return setting, err
}

// ReadSetting reads a setting
func ReadSetting(id uint) (*model.Setting, error) {
	setting := &model.Setting{}
	// Read Data
	err := db.ORM.First(&setting, "id = ?", id).Error
	return setting, err
}

// UpdateSetting reads a setting
func UpdateSetting(setting *model.Setting) (*model.Setting, error) {
	// Create change info
	err := db.ORM.Model(setting).Updates(setting).Error
	return setting, err
}

// DeleteSetting deletes setting with object id
func DeleteSetting(id uint) error {
	err := db.ORM.Delete(&model.Setting{ID: id}).Error
	return err
}

// ReadSettings return settings after retreive with params
func ReadSettings(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Setting, int, error) {
	settings := []*model.Setting{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("user_id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("name LIKE ? OR description LIKE ?", query, query)
	}
	// get total count of collection with initial query
	res.Find(&settings).Count(&totalCount)

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
	err := res.Find(&settings).Error

	return settings, totalCount, err
}
