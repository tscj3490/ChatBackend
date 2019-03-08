package worktimeService

import (
	"fmt"
	"time"

	"../../db"
	"../../model"
	"../../util/timeHelper"
)

// InitService inits service
func InitService() {

}

// func ReadWorktimeWithOrder(worktime *model.Worktime, book_time *model.Orders) (*model.Worktime, error) {
// {
// 	if err := db.ORM.Join().Error; err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

// func JoinWorktimeFromOrder(worktime *model.Worktime, book_time *model.Orders)
// {
// 	if err := db.ORM.Join(&worktime, &book_time).Error; err != nil {
// 		return nil, err
// 	}
// 	err := db.ORM.
// }
// CreateWorktime creates a worktime
func CreateWorktime(worktime *model.Worktime) (*model.Worktime, error) {
	// Insert Data
	if err := db.ORM.Create(&worktime).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&worktime).Error
	return worktime, err
}

// ReadWorktime reads a worktime
func ReadWorktime(id uint) (*model.Worktime, error) {
	worktime := &model.Worktime{}
	// Read Data
	err := db.ORM.First(&worktime, "id = ?", id).Error
	return worktime, err
}

// UpdateWorktime reads a worktime
func UpdateWorktime(worktime *model.Worktime) (*model.Worktime, error) {
	// Create change info
	err := db.ORM.Model(worktime).Updates(worktime).Error
	return worktime, err
}

// DeleteWorktime deletes worktime with object id
func DeleteWorktime(id uint) error {
	err := db.ORM.Delete(&model.Worktime{ID: id}).Error
	return err
}

// ReadWorktimes return worktimes after retreive with params
func ReadWorktimes(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Worktime, int, error) {
	worktimes := []*model.Worktime{}
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
	res.Find(&worktimes).Count(&totalCount)

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
	err := res.Find(&worktimes).Error

	return worktimes, totalCount, err
}

// ReadWorktimesByDate return worktimes after retreive with params
func ReadWorktimesByDate(query string, offset int, count int, field string, sort int, userID uint, date string) ([]*model.Worktime, int, error) {
	worktimes := []*model.Worktime{}
	totalCount := 0

	res := db.ORM

	if date != "" {
		d, err1 := time.Parse("2006-01-02", date)
		if err1 != nil {
			fmt.Println(err1)
		}
		//curDate_pure := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Now().Location())
		curDate := d
		lastTime := timeHelper.FewDaysLaterDate(d, 7)

		fmt.Println(curDate, lastTime)
		res = res.Where("start_time >= ? AND start_time <= ?", curDate, lastTime)
	}

	// get total count of collection with initial query
	res.Find(&worktimes).Count(&totalCount)

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
	err := res.Find(&worktimes).Error

	return worktimes, totalCount, err
}
