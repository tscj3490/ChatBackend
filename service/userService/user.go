package userService

import (
	"errors"
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateUser(user *model.User) (*model.User, error) {
	// check duplicate name
	var err error
	u := &model.User{}
	if res := db.ORM.Where("phone = ?", user.Phone).First(&u).RecordNotFound(); !res {
		// err := errors.New(user.Name + " is already registered")
		fmt.Println("=====", "omit")
		db.ORM.Omit("phone").Updates(user)
		return nil, err
	}

	// Insert Data
	if err = db.ORM.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, err
}

// ReadUser reads a user
func ReadUser(id uint) (*model.User, error) {
	user := &model.User{}
	// Read Data
	err := db.ORM.First(&user, "id = ?", id).Error
	return user, err
}

// UpdateUser reads a user
func UpdateUser(user *model.User) (*model.User, error) {

	// Create change info
	err := db.ORM.Model(user).Updates(user).Error
	return user, err
}

// UpdateProfile
func UpdateProfile(user *model.User, id uint) (*model.User, error) {

	// Create change info
	err := db.ORM.Model(user).Where("id = ?", id).Updates(user).Error
	return user, err
}

// DeleteUser deletes user with object id
func DeleteUser(id uint) error {
	err := db.ORM.Delete(&model.User{ID: id}).Error
	return err
}

// ReadUsersByTeamID
func ReadUsersByTeamID(teamID uint, offset int, count int, field string, sort int) ([]*model.User, int, error) {
	users := []*model.User{}
	totalCount := 0

	res := db.ORM
	if teamID != 0 {
		res = res.Where("team_id = ?", teamID)
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

// ReadUsers return users after retreive with params
func ReadUsers(query string, offset int, count int, field string, sort int) ([]*model.User, int, error) {
	users := []*model.User{}
	totalCount := 0

	res := db.ORM
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("name LIKE ?", query)
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
	if res := db.ORM.Where("name = ?", username).First(&user).RecordNotFound(); !res {
		return nil, errors.New("Same username is existed already")
	}
	return user, nil
}
