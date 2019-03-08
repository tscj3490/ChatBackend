package v1

import (
	"fmt"
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/makeService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitMakes(parentRoute *echo.Group) {
	route := parentRoute.Group("/makes")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createMake)) 
	route.GET("/:id", permission.AuthRequired(readMake))
	route.PUT("/:id", permission.AuthRequired(updateMake))
	route.DELETE("/:id", permission.AuthRequired(deleteMake))

	route.GET("", permission.AuthRequired(readMakes))
	route.GET("/field", permission.AuthRequired(readMakesByField))

	makeService.InitService()
}

// @Title createDevice
// @Description Create a device.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Device 		"Returns created device"
// @Failure 400 {object} response.BasicResponse "err.device.bind"
// @Failure 400 {object} response.BasicResponse "err.device.create"
// @Resource /devices
// @Router /devices [post]
func createMake(c echo.Context) error { 
	make := &model.Make{}
	if err := c.Bind(make); err != nil {
		fmt.Println("ddd", make, err)
		return response.KnownErrJSON(c, "err.make.bind", err)
	}

	// create make
	make, err := makeService.CreateMake(make)
	fmt.Println("sss", make, err)
	if err != nil {
		return response.KnownErrJSON(c, "err.make.create", err)
	}

	return response.SuccessInterface(c, make)
}

// @Title readDevice
// @Description Read a device.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Device ID."
// @Success 200 {object} model.Device 		"Returns read device"
// @Failure 400 {object} response.BasicResponse "err.device.bind"
// @Failure 400 {object} response.BasicResponse "err.device.read"
// @Resource /devices
// @Router /devices/{id} [get]
func readMake(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	make, err := makeService.ReadMake(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.make.read", err)
	}

	return response.SuccessInterface(c, make)
}

// @Title updateDevice
// @Description Update a device.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Device ID."
// @Param   avatar      	form   	string  true	"Device Avatar"
// @Param   firstname		form   	string  true	"Device Firstname"
// @Param   lastname		form   	string  true	"Device Lastname"
// @Param   email	    	form   	string  true	"Device Email"
// @Param   birth      		form   	Time   	true	"Device Birth"
// @Success 200 {object} model.Device 		"Returns read device"
// @Failure 400 {object} response.BasicResponse "err.device.bind"
// @Failure 400 {object} response.BasicResponse "err.device.read"
// @Resource /devices
// @Router /devices/{id} [put]
func updateMake(c echo.Context) error {
	make := &model.Make{}
	if err := c.Bind(make); err != nil {
		return response.KnownErrJSON(c, "err.make.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	make, err := makeService.UpdateMake(make)
	if err != nil {
		return response.KnownErrJSON(c, "err.make.update", err)
	}

	make, _ = makeService.ReadMake(make.ID)
	return response.SuccessInterface(c, make)
}

// @Title deleteDevice
// @Description Delete a device.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Device ID."
// @Success 200 {object} response.BasicResponse "Device is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.device.bind"
// @Failure 400 {object} response.BasicResponse "err.device.delete"
// @Resource /devices
// @Router /devices/{id} [delete]
func deleteMake(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete device with id
	err := makeService.DeleteMake(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.make.delete", err)
	}
	return response.SuccessJSON(c, "Make is deleted correctly.")
}

// @Title readDevices
// @Description Read devices with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Device is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.device.read"
// @Resource /devices
// @Router /devices [get]
func readMakes(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	makes, total, err := makeService.ReadMakes(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.make.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, makes})
}

func readMakesByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	makes, total, err := makeService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.make.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, makes})
}
