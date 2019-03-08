package v1

import (
	"fmt"
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/authService/vendorOrdersService"
	"../response"

	"github.com/labstack/echo"
)

// InitUsers inits user CRUD apis
// @Title Users
// @Description Users's router group.
func InitVendorOrders(parentRoute *echo.Group) {
	route := parentRoute.Group("/vendorOrders")
	//route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("/insert", permission.AuthRequired(createVendorOrders))
	route.GET("/:id", permission.AuthRequired(readVendorOrders))
	route.PUT("", permission.AuthRequired(updateVendorOrders))
	route.DELETE("/:id", permission.AuthRequired(deleteVendorOrders))

	route.GET("", permission.AuthRequired(readVendorsOrders))

	route.GET("/field", permission.AuthRequired(readVendorsOrdersByField))

	vendorOrdersService.InitService()
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
func createVendorOrders(c echo.Context) error {
	vendorOrders := &model.VendorOrders{}
	if err := c.Bind(vendorOrders); err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.bind", err)
	}
	fmt.Println("****************")
	// create vendor
	vendorOrders, err := vendorOrdersService.CreateVendorOrders(vendorOrders)

	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.create", err)
	}

	return response.SuccessInterface(c, vendorOrders)
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
func readVendorOrders(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, uint(id))

	vendorOrders, err := vendorOrdersService.ReadVendorOrders(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.read", err)
	}

	return response.SuccessInterface(c, vendorOrders)
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
func updateVendorOrders(c echo.Context) error {
	vendorOrders := &model.VendorOrders{}
	if err := c.Bind(vendorOrders); err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	vendorOrders, err := vendorOrdersService.UpdateVendorOrders(vendorOrders)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.udpate", err)
	}

	vendorOrders, _ = vendorOrdersService.ReadVendorOrders(vendorOrders.ID)

	return response.SuccessInterface(c, vendorOrders)
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
func deleteVendorOrders(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete user with id
	err := vendorOrdersService.DeleteVendorOrders(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.delete", err)
	}
	return response.SuccessJSON(c, "Vendor is deleted correctly.")
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

func readVendorsOrders(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))

	// read vendors with params
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	vendorsOrders, total, err := vendorOrdersService.ReadVendorsOrders(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsOrders})
}

func readVendorsOrdersByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendor_id"))

	// read vendors with params
	vendorsOrders, total, err := vendorOrdersService.ReadByField(query, offset, count, field, sort, uint(vendorID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorOrders.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsOrders})
}
