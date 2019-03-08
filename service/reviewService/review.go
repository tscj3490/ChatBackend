package reviewService

import (
	"fmt"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateReview creates a review
func CreateReview(review *model.Review) (*model.Review, error) {
	// Insert Data
	if err := db.ORM.Create(&review).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&review).Error
	return review, err
}

// ReadReview reads a review
func ReadReview(id uint) (*model.Review, error) {
	review := &model.Review{}
	// Read Data
	err := db.ORM.First(&review, "id = ?", id).Error
	return review, err
}

// UpdateReview reads a review
func UpdateReview(review *model.Review) (*model.Review, error) {
	// Create change info
	err := db.ORM.Model(review).Updates(review).Error
	return review, err
}

// DeleteDevice deletes review with object id
func DeleteReview(id uint) error {
	err := db.ORM.Delete(&model.Review{ID: id}).Error
	return err
}

// ReadReviews return reviews after retreive with params
func ReadReviews(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Review, int, error) {
	reviews := []*model.Review{}
	totalCount := 0

	for _, m := range reviews {
		fmt.Println(m.ID)
	}

	res := db.ORM
	if userID != 0 {
		res = res.Where("id = ?", userID)
	}
	if query != "" {
		query = "%" + query + "%"
		res = res.Where("customer_id LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&reviews).Count(&totalCount)

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
	err := res.Find(&reviews).Error

	return reviews, totalCount, err
}
