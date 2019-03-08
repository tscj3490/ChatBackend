package messageService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateMessage creates a message
func CreateMessage(message *model.Message) (*model.Message, error) {
	// Insert Data
	if err := db.ORM.Create(&message).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&message).Error
	return message, err
}

// ReadMessage reads a message
func ReadMessage(id uint) (*model.Message, error) {
	message := &model.Message{}
	// Read Data
	err := db.ORM.First(&message, "id = ?", id).Error
	return message, err
}

// UpdateMessage reads a message
func UpdateMessage(message *model.Message) (*model.Message, error) {
	// Create change info
	err := db.ORM.Model(message).Updates(message).Error
	return message, err
}

// DeleteMessage deletes make with object id
func DeleteMessage(id uint) error {
	err := db.ORM.Delete(&model.Message{ID: id}).Error
	return err
}

// ReadMessages return messages after retreive with params
func ReadMessages(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Message, int, error) {
	messages := []*model.Message{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("order_id LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&messages).Count(&totalCount)

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
	err := res.Find(&messages).Error

	return messages, totalCount, err
}
