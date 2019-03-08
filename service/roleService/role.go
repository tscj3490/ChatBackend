package roleService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateRole creates a role
func CreateRole(role *model.Role) (*model.Role, error) {
	// Insert Data
	if err := db.ORM.Create(&role).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&role).Error
	return role, err
}

// ReadRole reads a role
func ReadRole(id uint) (*model.Role, error) {
	role := &model.Role{}
	// Read Data
	err := db.ORM.First(&role, "id = ?", id).Error
	return role, err
}

// UpdateRole reads a role
func UpdateRole(role *model.Role) (*model.Role, error) {
	// Create change info
	err := db.ORM.Model(role).Updates(role).Error
	return role, err
}

// DeleteRole deletes role with object id
func DeleteRole(id uint) error {
	err := db.ORM.Delete(&model.Role{ID: id}).Error
	return err
}

// ReadRoles return roles after retreive with params
func ReadRoles(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Role, int, error) {
	roles := []*model.Role{}
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
	res.Find(&roles).Count(&totalCount)

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
	err := res.Find(&roles).Error

	return roles, totalCount, err
}
