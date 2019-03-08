package userService

import (
	"errors"

	"../../../db"
	"../../../model"
	// "golang.org/x/crypto/bcrypt"
)

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateUser(user *model.User) (*model.User, error) {
	// check duplicate username
	var err error
	u := &model.User{}
	if res := db.ORM.Where("phone = ?", user.Phone).First(&u).RecordNotFound(); !res {
		// err := errors.New(user.Name + " is already registered")
		db.ORM.Table("users").Where("phone = ?", user.Phone).Omit("phone").Updates(user)
		return user, err
	}

	// password, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), 8)
	// if err != nil {
	// 	return nil, err
	// }
	// user.UserPassword = string(password)

	// Insert Data
	if err := db.ORM.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ReadUser reads a user
func ReadUser(id uint) (*model.User, error) {
	user := &model.User{}
	// Read Data
	err := db.ORM.Table("users").Select("users.*, teams.name as team_name").
		Joins("left join teams on teams.id = users.team_id").
		First(&user, "users.id = ?", id).Error
	// err := db.ORM.First(&user, "id = ?", id).Error
	return user, err
}

// UpdateUser reads a user
func UpdateUser(user *model.User) (*model.User, error) {
	// Create change info
	err := db.ORM.Model(user).Updates(user).Error
	return user, err
}

// DeleteUser deletes user with object id
func DeleteUser(id uint) error {
	err := db.ORM.Delete(&model.User{ID: id}).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadUsers(query string, offset int, count int, field string, sort int) ([]*model.User, int, error) {
	users := []*model.User{}
	totalCount := 0

	res := db.ORM

	res = res.Table("users").Select("users.*, teams.name as team_name").
		Joins("left join teams on teams.id = users.team_id")

	if query != "" {
		query = "%" + query + "%"
		res = res.Where("users.name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&users).Count(&totalCount)

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
	err := res.Find(&users).Error

	return users, totalCount, err
}

// ReadUserByUsername returns user
func ReadUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	res := db.ORM.Where("name = ?", username).First(&user).RecordNotFound()
	if res {
		return nil, errors.New("User doesnot exist")
	}
	return user, nil
}
