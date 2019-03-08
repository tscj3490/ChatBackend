package transactionService

import (
	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateTransaction creates a transaction
func CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	// Insert Data
	if err := db.ORM.Create(&transaction).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&transaction).Error
	return transaction, err
}

// ReadTransaction reads a transaction
func ReadTransaction(id uint) (*model.Transaction, error) {
	transaction := &model.Transaction{}
	// Read Data
	err := db.ORM.First(&transaction, "id = ?", id).Error
	return transaction, err
}

// UpdateTransaction reads a transaction
func UpdateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	// Create change info
	err := db.ORM.Model(transaction).Updates(transaction).Error
	return transaction, err
}

// DeleteTransaction deletes transaction with object id
func DeleteTransaction(id uint) error {
	err := db.ORM.Delete(&model.Transaction{ID: id}).Error
	return err
}

// ReadTransactions return transactions after retreive with params
func ReadTransactions(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Transaction, int, error) {
	transactions := []*model.Transaction{}
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
