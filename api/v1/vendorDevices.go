package v1

import (
	"fmt"
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/authService/vendorDevicesService"
	"../response"

	"github.com/labstack/echo"
)

// InitUsers inits user CRUD apis
// @Title Users
// @Description Users's router group.
func InitVendorDevices(parentRoute *echo.Group) {
	route := parentRoute.Group("/vendorDevices")
	//route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("/insert", permission.AuthRequired(createVendorDevices))
	route.GET("/:id", permission.AuthRequired(readVendorDevices))
	route.PUT("", permission.AuthRequired(updateVendorDevices))
	route.DELETE("/:id", permission.AuthRequired(deleteVendorDevices))

	route.GET("", permission.AuthRequired(readVendorsDevices))

	route.GET("/field", permission.AuthRequired(readVendorsDevicesByField))

	vendorDevicesService.InitService()
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
func createVendorDevices(c echo.Context) error {
	vendorDevices := &model.VendorDevices{}
	if err := c.Bind(vendorDevices); err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.bind", err)
	}

	// create vendor
	vendorDevices, err := vendorDevicesService.CreateVendorDevices(vendorDevices)

	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.create", err)
	}

	return response.SuccessInterface(c, vendorDevices)
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
func readVendorDevices(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, uint(id))

	vendorDevices, err := vendorDevicesService.ReadVendorDevices(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.read", err)
	}

	return response.SuccessInterface(c, vendorDevices)
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
func updateVendorDevices(c echo.Context) error {
	vendorDevices := &model.VendorDevices{}
	if err := c.Bind(vendorDevices); err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	vendorDevices, err := vendorDevicesService.UpdateVendorDevices(vendorDevices)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.udpate", err)
	}

	vendorDevices, _ = vendorDevicesService.ReadVendorDevices(vendorDevices.ID)

	return response.SuccessInterface(c, vendorDevices)
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
func deleteVendorDevices(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete user with id
	err := vendorDevicesService.DeleteVendorDevices(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.delete", err)
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

func readVendorsDevices(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))

	// read vendors with params
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("vendorId"))

	vendorsDevices, total, err := vendorDevicesService.ReadVendorsDevices(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsDevices})
}

func readVendorsDevicesByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendor_id"))

	// read vendors with params
	vendorsDevices, total, err := vendorDevicesService.ReadByField(query, offset, count, field, sort, uint(vendorID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendorDevices.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendorsDevices})
}
