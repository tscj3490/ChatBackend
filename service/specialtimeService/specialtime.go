package specialtimeService

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

// func ReadspecialtimeWithOrder(specialtime *model.specialtime, book_time *model.Orders) (*model.specialtime, error) {
// {
// 	if err := db.ORM.Join().Error; err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

// func JoinspecialtimeFromOrder(specialtime *model.specialtime, book_time *model.Orders)
// {
// 	if err := db.ORM.Join(&specialtime, &book_time).Error; err != nil {
// 		return nil, err
// 	}
// 	err := db.ORM.
// }
// Createspecialtime creates a specialtime
func CreateSpecialtime(specialtime *model.Specialtime) (*model.Specialtime, error) {
	// Insert Data
	if err := db.ORM.Create(&specialtime).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&specialtime).Error
	return specialtime, err
}

// ReadSpecialtime reads a specialtime
func ReadSpecialtime(id uint) (*model.Specialtime, error) {
	specialtime := &model.Specialtime{}
	// Read Data
	err := db.ORM.First(&specialtime, "id = ?", id).Error
	return specialtime, err
}

// UpdateSpecialtime reads a specialtime
func UpdateSpecialtime(specialtime *model.Specialtime) (*model.Specialtime, error) {
	// Create change info
	err := db.ORM.Model(specialtime).Updates(specialtime).Error
	return specialtime, err
}

// DeleteSpecialtime deletes specialtime with object id
func DeleteSpecialtime(id uint) error {
	err := db.ORM.Delete(&model.Specialtime{ID: id}).Error
	return err
}

// ReadSpecialtimes return specialtimes after retreive with params
func ReadSpecialtimes(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Specialtime, int, error) {
	specialtimes := []*model.Specialtime{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	// get total count of collection with initial query
	res.Find(&specialtimes).Count(&totalCount)

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
	err := res.Find(&specialtimes).Error

	return specialtimes, totalCount, err
}

// ReadSpecialtimesByDate return specialtimes after retreive with params
func ReadSpecialtimesByDate(query string, offset int, count int, field string, sort int, userID uint, date string) ([]*model.Specialtime, int, error) {
	specialtimes := []*model.Specialtime{}
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
		res = res.Where("date >= ? AND date <= ?", curDate, lastTime)
	}

	// get total count of collection with initial query
	res.Find(&specialtimes).Count(&totalCount)

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
	err := res.Find(&specialtimes).Error

	return specialtimes, totalCount, err
}
