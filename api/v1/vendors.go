package v1

import (
	"fmt"
	"log"
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/authService/vendorsService"
	"../response"

	"github.com/labstack/echo"
)

// InitUsers inits user CRUD apis
// @Title Users
// @Description Users's router group.
func InitVendors(parentRoute *echo.Group) {
	route := parentRoute.Group("/vendors")
	//route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("/insert", permission.AuthRequired(createVendor))
	route.GET("/:id", permission.AuthRequired(readVendor))
	route.PUT("", permission.AuthRequired(updateVendor))
	route.DELETE("/:id", permission.AuthRequired(deleteVendor))

	route.GET("", permission.AuthRequired(readVendors))

	route.GET("/field", permission.AuthRequired(readVendorsByField))

	route.POST("/filter", permission.AuthRequired(readVendorsByFilter))
	route.GET("/customer", permission.AuthRequired(readVendorsByCustomer))

	route.POST("/worktime", permission.AuthRequired(createVendorByWorking))
	route.GET("/worktime", permission.AuthRequired(readVendorByWorking))

	vendorsService.InitService()
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
func createVendor(c echo.Context) error {
	vendor := &model.Vendor{}
	if err := c.Bind(vendor); err != nil {
		return response.KnownErrJSON(c, "err.vendor.bind", err)
	}

	log.Printf("%+v", vendor)
	// create vendor
	vendor, err := vendorsService.CreateVendor(vendor)

	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.create", err)
	}

	publicVendor := &model.PublicVendor{Vendor: vendor}
	return response.SuccessInterface(c, publicVendor)
}

// createVendorByWorking
func createVendorByWorking(c echo.Context) error {
	wkinfo := &model.WorktimeInfo{}

	if err := c.Bind(wkinfo); err != nil {
		return response.KnownErrJSON(c, "err.worktime.bind", err)
	}

	// create vendor
	wkinfo, err := vendorsService.CreateVendorByWorking(wkinfo)

	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.create", err)
	}

	return response.SuccessInterface(c, wkinfo)
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
func readVendor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, uint(id))

	vendor, err := vendorsService.ReadVendor(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.read", err)
	}
	// publicVendor := &model.PublicVendor{Vendor: vendor}
	return response.SuccessInterface(c, vendor)
}

//readVendorByWorking
func readVendorByWorking(c echo.Context) error {
	vendorID, _ := strconv.Atoi(c.FormValue("vendorId"))
	fmt.Println(vendorID)
	wktime, err := vendorsService.ReadVendorByWorking(uint(vendorID))
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.read", err)
	}
	// publicVendor := &model.PublicVendor{Vendor: vendor}
	return response.SuccessInterface(c, wktime)
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
func updateVendor(c echo.Context) error {
	vendor := &model.Vendor{}
	if err := c.Bind(vendor); err != nil {
		return response.KnownErrJSON(c, "err.vendor.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	vendor, err := vendorsService.UpdateVendor(vendor)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.udpate", err)
	}

	// vendor, _ = vendorsService.ReadVendor(vendor.ID)
	// publicVendor := &model.PublicVendor{Vendor: vendor}
	return response.SuccessInterface(c, vendor)
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
func deleteVendor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete user with id
	err := vendorsService.DeleteVendor(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.delete", err)
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

func readVendors(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))

	// read vendors with params
	vendors, total, err := vendorsService.ReadVendors(query, offset, count, field, sort)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendors})
}

func readVendorsByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read vendors with params
	vendors, total, err := vendorsService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, vendors})
}

func readVendorsByFilter(c echo.Context) error {
	postval := &model.PostVendor{}
	vendors := []*model.Vendor{}
	if err := c.Bind(postval); err != nil {
		return response.KnownErrJSON(c, "err.postval.bind", err)
	}
	fmt.Println("%+v", postval)
	// read vendors with params
	vendors, err := vendorsService.ReadByFilter(postval)

	if err != nil {
		return response.KnownErrJSON(c, "err.vendors.read", err)
	}

	return response.SuccessInterface(c, vendors)
}

func readVendorsByCustomer(c echo.Context) error {
	count, _ := strconv.Atoi(c.FormValue("count"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendorId"))

	// read customers with params
	customers, total, err := vendorsService.ReadByCustomer(count, uint(vendorID))
	if err != nil {
		return response.KnownErrJSON(c, "err.customers.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, customers})
}
