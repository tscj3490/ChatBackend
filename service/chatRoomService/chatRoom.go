package chatRoomService

import (
	"strings"

	"../../db"
	"../../model"
	"../../util/notification"
)

// InitService inits service
func InitService() {

}

// CreateChatRoom creates a chatRoom
func CreateChatRoom(chatRoom *model.ChatRoom) (*model.ChatRoom, error) {
	// Insert Data
	if err := db.ORM.Create(&chatRoom).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&chatRoom).Error
	return chatRoom, err
}

// ReadChatRoom reads a chatRoom
func ReadChatRoom(id uint) (*model.ChatRoom, error) {
	chatRoom := &model.ChatRoom{}
	// Read Data
	err := db.ORM.First(&chatRoom, "id = ?", id).Error
	return chatRoom, err
}

// UpdateChatRoom reads a chatRoom
func UpdateChatRoom(chatRoom *model.ChatRoom) (*model.ChatRoom, error) {
	// Create change info
	err := db.ORM.Model(chatRoom).Updates(chatRoom).Error
	return chatRoom, err
}

// DeleteChatRoom deletes chatRoom with object id
func DeleteChatRoom(id uint) error {
	err := db.ORM.Delete(&model.ChatRoom{ID: id}).Error
	return err
}

// SendMessage
func SendMessage(msgInfo *model.SendMessageInfo, senderName string) (*model.ChatRoom, error) {
	chatRoom := &model.ChatRoom{}
	var err error
	var body string

	if msgInfo.Message.Text != "" {
		body = msgInfo.Message.Text
	} else {
		body = "Image or Contact"
		if msgInfo.Message.Image != "" {
			body = "Image"
		}
		if msgInfo.Message.Contact != "" {
			body = "Contact"
		}
	}
	err = db.ORM.First(&chatRoom, "id = ?", msgInfo.GroupID).Error

	if err == nil {
		ids := strings.Split(chatRoom.UserIDs, ",")

		for _, id := range ids {
			user := &model.User{}
			err = db.ORM.First(&user, "id = ?", id).Error

			if err != nil {
				// continue
			} else {
				if user != nil {
					if senderName == user.Name {
						continue
					}
					if senderName == "" {
						senderName = "Unnamed"
					}
					if user.PushToken != "" {
						go notification.SendNotification(user.PushToken, senderName, body)
					}
				}
			}
		}
	}
	return chatRoom, nil
}

// ReadChatRooms return chatRooms after retreive with params
func ReadChatRooms(query string, offset int, count int, field string, sort int, userID uint) ([]*model.ChatRoom, int, error) {
	chatRooms := []*model.ChatRoom{}
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
	res.Find(&chatRooms).Count(&totalCount)

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
	err := res.Find(&chatRooms).Error

	return chatRooms, totalCount, err
}

// ReadChatRoomsByGroupId return chatRooms after retreive with params
func ReadChatRoomsByGroupId(id string) ([]*model.ChatRoom, int, error) {
	chatRooms := []*model.ChatRoom{}
	totalCount := 0

	res := db.ORM
	if id != "" {
		res = res.Where("id = ?", id)
	}

	// get total count of collection with initial query
	res.Find(&chatRooms).Count(&totalCount)

	err := res.Find(&chatRooms).Error

	return chatRooms, totalCount, err
}

// ReadChatRoomsByUserId return chatRooms after retreive with params
func ReadChatRoomsByUserId(id string) ([]*model.ChatRoom, int, error) {
	chatRooms := []*model.ChatRoom{}
	totalCount := 0

	res := db.ORM

	groupList := []*model.ChatRoom{}

	res.Find(&chatRooms).Count(&totalCount)

	for _, chatroom := range chatRooms {
		ids := strings.Split(chatroom.UserIDs, ",")
		for _, i := range ids {
			if i != id {
				continue
			} else {
				groupList = append(groupList, chatroom)
			}
		}
	}

	return groupList, totalCount, nil
}
