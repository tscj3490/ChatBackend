package companyService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateCompany creates a company
func CreateCompany(company *model.Company) (*model.Company, error) {
	// Insert Data
	if err := db.ORM.Create(&company).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&company).Error
	return company, err
}

// ReadCompany reads a company
func ReadCompany(id uint) (*model.Company, error) {
	company := &model.Company{}
	// Read Data
	err := db.ORM.First(&company, "id = ?", id).Error
	return company, err
}

// UpdateCompany reads a company
func UpdateCompany(company *model.Company) (*model.Company, error) {
	// Create change info
	err := db.ORM.Model(company).Updates(company).Error
	return company, err
}

// DeleteCompany deletes company with object id
func DeleteCompany(id uint) error {
	err := db.ORM.Delete(&model.Company{ID: id}).Error
	return err
}

// ReadCompanies return companies after retreive with params
func ReadCompanies(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Company, int, error) {
	companies := []*model.Company{}
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
	res.Find(&companies).Count(&totalCount)

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
	err := res.Find(&companies).Error

	return companies, totalCount, err
}
