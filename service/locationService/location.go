package locationService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateLocation creates a location
func CreateLocation(location *model.Location) (*model.Location, error) {
	// Insert Data
	if err := db.ORM.Create(&location).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&location).Error
	return location, err
}

// ReadLocation reads a location
func ReadLocation(id uint) (*model.Location, error) {
	location := &model.Location{}
	// Read Data
	err := db.ORM.First(&location, "id = ?", id).Error
	return location, err
}

// UpdateLocation reads a location
func UpdateLocation(location *model.Location) (*model.Location, error) {
	// Create change info
	err := db.ORM.Model(location).Updates(location).Error
	return location, err
}

// DeleteLocation deletes location with object id
func DeleteLocation(id uint) error {
	err := db.ORM.Delete(&model.Location{ID: id}).Error
	return err
}

// ReadLocations return locations after retreive with params
func ReadLocations(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Location, int, error) {
	locations := []*model.Location{}
	totalCount := 0

	res := db.ORM
	if userID != 0 {
		res = res.Where("user_id = ?", userID)
	}
	// get total count of collection with initial query
	res.Find(&locations).Count(&totalCount)

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
	err := res.Find(&locations).Error

	return locations, totalCount, err
}
