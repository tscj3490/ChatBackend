package customersService

import (
	"errors"
	"fmt"

	"../../../db"
	"../../../model"

	"../../../util/random"
	"golang.org/x/crypto/bcrypt"
)

type Result struct {
	Customer string
	Average  float32
}

// InitService inits service
func InitService() {

}

// CreateUser creates a user
func CreateCustomer(customer *model.Customer) (*model.Customer, error) {
	// check duplicate username
	//	u := &model.Customers{}
	if res := db.ORM.Where("username = ?", customer.UserName).First(&customer).RecordNotFound(); !res {
		err := errors.New(customer.UserName + " is already registered")
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 10)
	if err != nil {
		return nil, err
	}
	customer.Password = string(password)

	// Insert Data
	if err := db.ORM.Create(&customer).Error; err != nil {
		return nil, err
	}
	return customer, err

}

// CreateUser creates a user
func CreateCustomerWithEmail(customer *model.Customer) (*model.Customer, error) {
	// check duplicate username
	v := &model.Customer{}

	//	customer.UserName = customer.Email
	if res := db.ORM.Where("username = ?", customer.UserName).First(&v).RecordNotFound(); !res {
		err := errors.New(customer.UserName + " is already registered")
		return nil, err
	}

	verifyCode := random.GenerateRandomString(8)
	//	password, err := bcrypt.GenerateFromPassword([]byte(verifyCode), 8)
	//	if err != nil {
	//		return nil, err
	//	}
	customer.Password = string(verifyCode)
	customer.CoinAmount = 50.0
	// email send
	fmt.Println(verifyCode)
	fmt.Println(customer.Email)
	fmt.Println(customer.UserName)
	// go email.SendForgotEmail(customer.Email, customer.UserName, verifyCode)

	// Insert Data
	if err := db.ORM.Create(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

// ReadUser reads a user
func ReadCustomer(id uint) (*model.Customer, error) {
	customer := &model.Customer{}
	// Read Data
	err := db.ORM.First(&customer, "id = ?", id).Error
	return customer, err
}

// UpdateUser reads a user
func UpdateCustomer(customer *model.Customer) (*model.Customer, error) {
	// Create change info
	err := db.ORM.Model(customer).Updates(customer).Error
	return customer, err
}

// DeleteUser deletes user with object id
func DeleteCustomer(id uint) error {
	customer := &model.Customer{}

	err := db.ORM.Where("id = ?", id).Delete(customer).Error
	return err
}

// ReadUsers return users after retreive with params
func ReadCustomers(query string, offset int, count int, field string, sort int) ([]*model.Customer, int, error) {
	customers := []*model.Customer{}
	totalCount := 0

	res := db.ORM

	results := []*Result{}
	res.Table("reviews").Select("customer_id as customer, avg(score) as average").Group("customer_id").Where("type = ?", 0).Scan(&results)

	fmt.Println(results)
	customerIds := []string{}
	for _, r := range results {
		fmt.Println(r.Customer, r.Average)
		customerIds = append(customerIds, r.Customer)
		res.Table("customers").Where("id = ?", r.Customer).Update("rate", r.Average)
	}

	dbb := res.Table("customers")
	//	dbb := res.Table("customers").Where("customers.id IN (?)", customerIds)
	// if query != "" {
	// 	query = "%" + query + "%"
	// 	res = res.Where("username LIKE ? OR fullname LIKE ? OR email LIKE ?", query, query, query)
	// }
	// get total count of collection with initial query
	dbb.Find(&customers).Count(&totalCount)

	// add page feature
	if offset != 0 || count != 0 {
		dbb = dbb.Offset(offset)
		dbb = dbb.Limit(count)
	}
	// add sort feature
	if field != "" && sort != 0 {
		if sort > 0 {
			dbb = dbb.Order(field)
		} else {
			dbb = dbb.Order(field + " desc")
		}
	}
	err := dbb.Find(&customers).Error

	return customers, totalCount, err
}

// ReadUserByUsername returns user
func ReadCustomerByUsername(username string) (*model.Customer, error) {
	customer := &model.Customer{}
	fmt.Println("username:", username)

	res := db.ORM.Where("username = ?", username).First(&customer).RecordNotFound()
	fmt.Println("customer:", customer)
	fmt.Println("res:", res)
	if res {
		return nil, errors.New("User doesnot exist")
	}
	return customer, nil
}

// ReadUserByUsername returns user
func ReadBookByEmail(email string) (*model.Customer, error) {
	book := &model.Customer{}
	fmt.Println("email:", email)

	res := db.ORM.Where("email = ?", email).First(&book).RecordNotFound()
	fmt.Println("book:", book)
	fmt.Println("res:", res)
	if res {
		return nil, errors.New("User doesnot exist")
	}
	return book, nil
}

// ReadByField returns user
func ReadByField(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Customer, int, error) {
	customers := []*model.Customer{}
	totalCount := 0

	res := db.ORM
	// if userID != 0 {
	// 	res = res.Where("userID = ?", userID)
	// }

	if query != "" {
		query = "%" + query + "%"
		res = res.Where(fmt.Sprintf("%v LIKE ? OR %v LIKE ?", field, field), query, query)
	}
	// get total count of collection with initial query
	res.Find(&customers).Count(&totalCount)

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
	err := res.Find(&customers).Error

	return customers, totalCount, err
}
