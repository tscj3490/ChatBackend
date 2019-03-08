package userSettingService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// InitUserSetting inits user setting when create user
func InitUserSetting(userID uint) error {
return nil
}

// UpsertUserSetting upserts a userSetting
func UpsertUserSetting(userSetting *model.UserSetting) (*model.UserSetting, error) {
	// check duplicate userSettingname
	setting := &model.UserSetting{}
	if db.ORM.Where("user_id = ? AND code = ?", userSetting.UserID, userSetting.Code).First(&setting).RecordNotFound() {
		if err := db.ORM.Create(&setting).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.ORM.Model(userSetting).Where("user_id = ? AND code = ?", userSetting.UserID, userSetting.Code).Updates(userSetting).Error; err != nil {
			return nil, err
		}
	}
	return userSetting, nil
}

// DeleteUserSetting deletes userSetting with object id
func DeleteUserSetting(userID uint) error {
	err := db.ORM.Where("user_id", userID).Delete(model.UserSetting{}).Error
	return err
}

// ReadUserSettings return userSettings after retreive with params
func ReadUserSettings(userID uint) ([]*model.UserSetting, error) {
	userSettings := []*model.UserSetting{}
	err := db.ORM.Where("user_id = ?", userID).Find(&userSettings).Error
	return userSettings, err
}
