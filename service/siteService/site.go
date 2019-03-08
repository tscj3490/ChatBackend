package siteService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateSite creates a site
func CreateSite(site *model.Site) (*model.Site, error) {
	// Insert Data
	if err := db.ORM.Create(&site).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&site).Error
	return site, err
}

// ReadSite reads a site
func ReadSite(id uint) (*model.Site, error) {
	site := &model.Site{}
	// Read Data
	err := db.ORM.First(&site, "id = ?", id).Error
	return site, err
}

// UpdateSite reads a site
func UpdateSite(site *model.Site) (*model.Site, error) {
	// Create change info
	err := db.ORM.Model(site).Updates(site).Error
	return site, err
}

// DeleteSite deletes site with object id
func DeleteSite(id uint) error {
	err := db.ORM.Delete(&model.Site{ID: id}).Error
	return err
}

// ReadSites return sites after retreive with params
func ReadSites(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Site, int, error) {
	sites := []*model.Site{}
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
	res.Find(&sites).Count(&totalCount)

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
	err := res.Find(&sites).Error

	return sites, totalCount, err
}
