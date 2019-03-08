package productService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// ReadBlacklists return blacklists after retreive with params
func ReadProducts() ([]model.Product, error) {
	var products []model.Product

	res := db.ORM
	// get total count of collection with initial query
	err := res.Find(&products).Error

	return products, err
}
