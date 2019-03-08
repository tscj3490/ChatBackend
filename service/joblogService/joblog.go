package joblogService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateJoblog creates a joblog
func CreateJoblog(joblog *model.Joblog) (*model.Joblog, error) {
	// Insert Data
	if err := db.ORM.Create(&joblog).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&joblog).Error
	return joblog, err
}

// ReadJoblog reads a joblog
func ReadJoblog(id uint) (*model.Joblog, error) {
	joblog := &model.Joblog{}
	// Read Data
	err := db.ORM.First(&joblog, "id = ?", id).Error
	return joblog, err
}

// UpdateJoblog reads a joblog
func UpdateJoblog(joblog *model.Joblog) (*model.Joblog, error) {
	// Create change info
	err := db.ORM.Model(joblog).Updates(joblog).Error
	return joblog, err
}

// DeleteJoblog deletes joblog with object id
func DeleteJoblog(id uint) error {
	err := db.ORM.Delete(&model.Joblog{ID: id}).Error
	return err
}

// ReadJoblogs return joblogs after retreive with params
func ReadJoblogs(query string, offset int, count int, field string, sort int, orderID uint) ([]*model.Joblog, int, error) {
	joblogs := []*model.Joblog{}
	totalCount := 0

	res := db.ORM
	if orderID != 0 {
		res = res.Where("order_id = ?", orderID)
	}
	// if query != "" {
	// 	query = "%" + query + "%"
	// 	res = res.Where("license_number LIKE ? OR description LIKE ?", query, query)
	// }
	// get total count of collection with initial query
	res.Find(&joblogs).Count(&totalCount)

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
	err := res.Find(&joblogs).Error

	return joblogs, totalCount, err
}
