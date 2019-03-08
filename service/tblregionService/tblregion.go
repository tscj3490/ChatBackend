package tblregionService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateTblregion creates a tblregion
func CreateTblregion(tblregion *model.Tblregion) (*model.Tblregion, error) {
	// Insert Data
	if err := db.ORM.Create(&tblregion).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&tblregion).Error
	return tblregion, err
}

// ReadTblregion reads a tblregion
func ReadTblregion(id uint) (*model.Tblregion, error) {
	tblregion := &model.Tblregion{}
	// Read Data
	err := db.ORM.First(&tblregion, "id = ?", id).Error
	return tblregion, err
}

// UpdateTblregion reads a tblregion
func UpdateTblregion(tblregion *model.Tblregion) (*model.Tblregion, error) {
	// Create change info
	err := db.ORM.Model(tblregion).Updates(tblregion).Error
	return tblregion, err
}

// DeleteTblregion deletes tblregion with object id
func DeleteTblregion(id uint) error {
	err := db.ORM.Delete(&model.Tblregion{ID: id}).Error
	return err
}

// ReadTblregions return transactions after retreive with params
func ReadTblregions(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Tblregion, int, error) {
	transactions := []*model.Tblregion{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("vendor_id LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&transactions).Count(&totalCount)

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
	err := res.Find(&transactions).Error

	return transactions, totalCount, err
}
