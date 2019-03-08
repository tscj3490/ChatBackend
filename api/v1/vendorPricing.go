package v1

import (
	"fmt"
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/authService/vendorPricingService"
	"../response"

	"github.com/labstack/echo"
)

// InitUsers inits user CRUD apis
// @Title Users
// @Description Users's router group.
func InitVendorPricing(parentRoute *echo.Group) {
	route := parentRoute.Group("/vendorPricing")
	//route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("/insert", permission.AuthRequired(createVendorPricing))
	route.GET("/:id", permission.AuthRequired(readVendorPricing))
	route.PUT("", permission.AuthRequired(updateVendorPricing))
	route.DELETE("/:id", permission.AuthRequired(deleteVendorPricing))

	route.GET("", permission.AuthRequired(readVendorsPricing))

	route.GET("/field", permission.AuthRequired(readVendorsPricingByField))

	vendorPricingService.InitService()
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
func createVendorPricing(c echo.Context) error {
	vendorPricing := &model.VendorPricing{}
	if err := c.Bind(vendorPricing); err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.bind", err)
	}

	// create vendor
	vendorPricing, err := vendorPricingService.CreateVendorPricing(vendorPricing)

	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.create", err)
	}

	return response.SuccessInterface(c, vendorPricing)
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
func readVendorPricing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, uint(id))

	vendorPricing, err := vendorPricingService.ReadVendorPricing(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.read", err)
	}

	return response.SuccessInterface(c, vendorPricing)
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
func updateVendorPricing(c echo.Context) error {
	vendorPricing := &model.VendorPricing{}
	if err := c.Bind(vendorPricing); err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	vendorPricing, err := vendorPricingService.UpdateVendorPricing(vendorPricing)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.udpate", err)
	}

	vendorPricing, _ = vendorPricingService.ReadVendorPricing(vendorPricing.ID)

	return response.SuccessInterface(c, vendorPricing)
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
func deleteVendorPricing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete user with id
	err := vendorPricingService.DeleteVendorPricing(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.delete", err)
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

func readVendorsPricing(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))

	// read vendors with params
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("vendorId"))

	vendorsPricing, total, err := vendorPricingService.ReadVendorsPricing(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsPricing})
}

func readVendorsPricingByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendor_id"))
	deviceID, _ := strconv.Atoi(c.FormValue("device_id"))

	// read vendors with params
	vendorsPricing, total, err := vendorPricingService.ReadByField(query, offset, count, field, sort, uint(vendorID), uint(deviceID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorPricing.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsPricing})
}
