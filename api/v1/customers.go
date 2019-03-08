package v1

import (
	"fmt"
	"strconv"

	"../../model"
	"../../service/authService/customersService"
	"../../service/authService/permission"
	"../response"

	"github.com/labstack/echo"
)

// InitUsers inits user CRUD apis
// @Title Users
// @Description Users's router group.
func InitCustomers(parentRoute *echo.Group) {
	route := parentRoute.Group("/customers")
	// route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("/insert", permission.AuthRequired(createCustomer))
	route.GET("/:id", permission.AuthRequired(readCustomer)) //
	//	route.GET("/:id", permission.AuthRequired(readUser))
	route.PUT("", permission.AuthRequired(updateCustomer))
	route.DELETE("/:id", permission.AuthRequired(deleteCustomer))

	route.GET("", permission.AuthRequired(readCustomers))
	route.GET("/field", permission.AuthRequired(readCustomersByField))

	customersService.InitService()
}

// @Title createUser
// @Description Create a user.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   email       	form   	string  true	"User Email."
// @Param   password		form   	string 	true	"User Password."
// @Success 200 {object} model.PublicUser 		"Returns created user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.create"
// @Resource /users
// @Router /users [post]
func createCustomer(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		return response.KnownErrJSON(c, "err.customer.bind", err)
	}

	// create user
	customer, err := customersService.CreateCustomer(customer)
	//customer, err := gorm.Open("mysql", config.MysqlDSL())

	//if err := customer.Create(customer).Error; err != nil {
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.create", err)
	}

	//return response.SuccessInterface(c, &ListForm{len(publicCustomer), publicCustomer})
	publicCustomer := &model.PublicCustomer{Customer: customer}
	return response.SuccessInterface(c, publicCustomer)
}

// @Title readUser
// @Description Read a user.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"User ID."
// @Success 200 {object} model.PublicUser 		"Returns read user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /users
// @Router /users/{id} [get]
func readCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, uint(id))
	customer, err := customersService.ReadCustomer(uint(id))
	//db, err := gorm.Open("mysql", config.MysqlDSL())
	//var publicCustomer model.Customer
	//	db.First(&publicCustomer, "username = ?", "Bing")
	//db.First(&publicCustomer, "id = ?", id)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.read", err)
	}
	publicCustomer := &model.PublicCustomer{Customer: customer}
	return response.SuccessInterface(c, publicCustomer)
}

// @Title updateUser
// @Description Update a user.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"User ID."
// @Param   avatar      	form   	string  true	"User Avatar"
// @Param   firstname		form   	string  true	"User Firstname"
// @Param   lastname		form   	string  true	"User Lastname"
// @Success 200 {object} model.PublicUser 		"Returns read user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /users
// @Router /users/{id} [put]
func updateCustomer(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		return response.KnownErrJSON(c, "err.customer.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	customer, err := customersService.UpdateCustomer(customer)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.udpate", err)
	}

	customer, _ = customersService.ReadCustomer(customer.ID)
	publicCustomer := &model.PublicCustomer{Customer: customer}
	return response.SuccessInterface(c, publicCustomer)
}

// @Title deleteUser
// @Description Delete a user.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"User ID."
// @Success 200 {object} response.BasicResponse "User is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.delete"
// @Resource /users
// @Router /users/{id} [delete]
func deleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete user with id
	err := customersService.DeleteCustomer(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.delete", err)
	}
	return response.SuccessJSON(c, "Customer is deleted correctly.")
}

// @Title readUsers
// @Description Read users with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} UsersForm 				"Returned list users."
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /users
// @Router /users [get]

func readCustomers(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))

	// read vendors with params
	customers, total, err := customersService.ReadCustomers(query, offset, count, field, sort)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, customers})
}

func readCustomersByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read vendors with params
	customers, total, err := customersService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, customers})
}
