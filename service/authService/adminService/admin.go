package adminService

import (
	"errors"

	"../../../db"
	"../../../model"

	"golang.org/x/crypto/bcrypt"
)

// InitService inits service
func InitService() {

}

// CreateAdmin creates a admin
func CreateAdmin(admin *model.Admin) (*model.Admin, error) {
	// check duplicate email
	a := &model.Admin{}
	if !db.ORM.Where("email = ?", admin.Email).First(&a).RecordNotFound() {
		err := errors.New(admin.Email + " is already registered")
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return nil, err
	}
	admin.Password = string(password)

	// Insert Data
	if err := db.ORM.Create(&admin).Error; err != nil {
		return nil, err
	}
	return admin, err
}

// ReadAdmin reads a admin
func ReadAdmin(id uint) (*model.Admin, error) {
	admin := &model.Admin{}
	// Read Data
	err := db.ORM.First(&admin, "id = ?", id).Error
	return admin, err
}

// UpdateAdmin reads a admin
func UpdateAdmin(admin *model.Admin) (*model.Admin, error) {
	// Create change info
	err := db.ORM.Model(admin).Updates(admin).Error
	return admin, err
}

// DeleteAdmin deletes admin with object id
func DeleteAdmin(id uint) error {
	err := db.ORM.Delete(&model.Admin{ID: id}).Error
	return err
}

// ReadAdmins return admins after retreive with params
func ReadAdmins(query string, offset int, count int, field string, sort int) ([]*model.Admin, int, error) {
	admins := []*model.Admin{}
	totalCount := 0

	res := db.ORM
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&admins).Count(&totalCount)

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
	err := res.Find(&admins).Error

	return admins, totalCount, err
}

// ReadAdminByEmail returns admin
func ReadAdminByEmail(email string) (*model.Admin, error) {
	admin := &model.Admin{}
	res := db.ORM.Table("admins").Where("email = ?", email).First(&admin).RecordNotFound()
	if res {
		return nil, errors.New("Admin is not found")
	}
	return admin, nil
}
