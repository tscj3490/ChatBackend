package reminderService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateReminder creates a reminder
func CreateReminder(reminder *model.Reminder) (*model.Reminder, error) {
	// Insert Data
	if err := db.ORM.Create(&reminder).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&reminder).Error
	return reminder, err
}

// ReadReminder reads a reminder
func ReadReminder(id uint) (*model.Reminder, error) {
	reminder := &model.Reminder{}
	// Read Data
	err := db.ORM.First(&reminder, "id = ?", id).Error
	return reminder, err
}

// UpdateReminder reads a reminder
func UpdateReminder(reminder *model.Reminder) (*model.Reminder, error) {
	// Create change info
	err := db.ORM.Model(reminder).Updates(reminder).Error
	return reminder, err
}

// DeleteReminder deletes reminder with object id
func DeleteReminder(id uint) error {
	err := db.ORM.Delete(&model.Reminder{ID: id}).Error
	return err
}

// ReadReminders return reminders after retreive with params
func ReadReminders(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Reminder, int, error) {
	reminders := []*model.Reminder{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("user_id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("name LIKE ? OR address1 LIKE ? OR address2 LIKE ? OR city LIKE ? OR description LIKE ?", query, query, query, query, query)
	}
	// get total count of collection with initial query
	res.Find(&reminders).Count(&totalCount)

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
	err := res.Find(&reminders).Error

	return reminders, totalCount, err
}

// ReadReminderByGroupId
func ReadReminderByGroupId(id string) ([]*model.Reminder, int, error) {
	reminders := []*model.Reminder{}
	totalCount := 0

	res := db.ORM
	if id != "" {
		res = res.Where("group_id = ?", id)
	}
	// get total count of collection with initial query
	res.Find(&reminders).Count(&totalCount)

	err := res.Find(&reminders).Error

	return reminders, totalCount, err
}
