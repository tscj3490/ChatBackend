package auth

import (
	"errors"
	"fmt"

	"../../../config"
	"../../../model"
	"../../../service/authService/customersService"
	"../../../service/authService/permission"
	"../../response"

	"github.com/labstack/echo"
)

// @Title loginCustomer
// @Description Login a customer.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"Customer Email."
// @Param   password	form   string 	true	"Customer Password."
// @Success 200 {object} CustomerForm 				"Returns login Customer"
// @Failure 400 {object} response.BasicResponse "err.customer.bind"
// @Failure 400 {object} response.BasicResponse "err.customer.incorrect"
// @Failure 400 {object} response.BasicResponse "err.customer.token"
// @Resource /customer/login
// @Router /customer/login [post]
func loginCustomer(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(&customer); err != nil {
		return response.KnownErrJSON(c, "err.customer.bind", err)
	}
	//	fmt.Printf("%+v", customer)
	// check user crediential
	u, err := customersService.ReadCustomerByUsername(customer.UserName)
	fmt.Println(u)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.incorrect", errors.New("Incorrect username"))
	}
	fmt.Println(customer.UserName)
	fmt.Println(customer.Email)
	fmt.Println(customer.Password)
	fmt.Println(u.Password)
	//	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(customer.Password))
	if u.Password != customer.Password {
		return response.KnownErrJSON(c, "err.customer.incorrect", errors.New("Incorrect password"))
	}
	//	if err != nil {
	//		return response.KnownErrJSON(c, "err.customer.incorrect", errors.New("Incorrect password"))
	//	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(u.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public customer
	publicCustomer := &model.PublicCustomer{Customer: u}
	return response.SuccessInterface(c, UserForm{t, publicCustomer})
}

// @Title registerUser
// @Description Register a user.
// @Accept  json
// @Produce	json
// @Param   firstname	form   string   true	"User Firstname."
// @Param   lastname   	form   string   true	"User Lastname."
// @Param   email       form   string   true	"User Email."
// @Param   password	form   string 	true	"User Password."
// @Success 200 {object} UserForm				"Returns registered user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.exist"
// @Failure 400 {object} response.BasicResponse "err.user.create"
// @Failure 400 {object} response.BasicResponse "err.user.token"
// @Resource /user/register
// @Router /user/register [post]
func registerCustomer(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		return response.KnownErrJSON(c, "err.customer.bind", err)
	}

	// check existed email
	if _, err := customersService.ReadCustomerByUsername(customer.UserName); err == nil {
		return response.KnownErrJSON(c, "err.customer.exist",
			errors.New("Same username is existed. Please input other username"))
	}

	// create user with registered info
	customer, err := customersService.CreateCustomerWithEmail(customer)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.create", err)
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(customer.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public user
	publicCustomer := &model.PublicCustomer{Customer: customer}
	return response.SuccessInterface(c, UserForm{t, publicCustomer})
}

// @Title loginBook
// @Description Login a Book.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"Book Email."
// @Param   password	form   string 	true	"Book Password."
// @Success 200 {object} BookForm 				"Returns login Book"
// @Failure 400 {object} response.BasicResponse "err.Book.bind"
// @Failure 400 {object} response.BasicResponse "err.Book.incorrect"
// @Failure 400 {object} response.BasicResponse "err.Book.token"
// @Resource /Book/login
// @Router /Book/login [post]
func loginBook(c echo.Context) error {
	book := &model.Customer{}
	if err := c.Bind(&book); err != nil {
		return response.KnownErrJSON(c, "err.book.bind", err)
	}
	//	fmt.Printf("%+v", book)
	// check user crediential
	u, err := customersService.ReadBookByEmail(book.Email)
	fmt.Println(u)
	if err != nil {
		return response.KnownErrJSON(c, "err.book.incorrect", errors.New("Incorrect email"))
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(u.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.book.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public book
	publicBook := &model.PublicBook{Customer: u}
	return response.SuccessInterface(c, UserForm{t, publicBook})
}
