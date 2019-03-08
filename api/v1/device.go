package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/deviceService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitDevices(parentRoute *echo.Group) {
	route := parentRoute.Group("/devices")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createDevice))
	route.GET("/:id", permission.AuthRequired(readDevice))
	route.PUT("/:id", permission.AuthRequired(updateDevice))
	route.DELETE("/:id", permission.AuthRequired(deleteDevice))

	route.GET("", permission.AuthRequired(readDevices))

	deviceService.InitService()
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
func createDevice(c echo.Context) error {
	device := &model.Device{}
	if err := c.Bind(device); err != nil {
		return response.KnownErrJSON(c, "err.device.bind", err)
	}

	// create device
	device, err := deviceService.CreateDevice(device)
	if err != nil {
		return response.KnownErrJSON(c, "err.device.create", err)
	}

	return response.SuccessInterface(c, device)
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
func readDevice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	device, err := deviceService.ReadDevice(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.device.read", err)
	}

	return response.SuccessInterface(c, device)
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
func updateDevice(c echo.Context) error {
	device := &model.Device{}
	if err := c.Bind(device); err != nil {
		return response.KnownErrJSON(c, "err.device.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	device, err := deviceService.UpdateDevice(device)
	if err != nil {
		return response.KnownErrJSON(c, "err.device.update", err)
	}

	device, _ = deviceService.ReadDevice(device.ID)
	return response.SuccessInterface(c, device)
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
func deleteDevice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete device with id
	err := deviceService.DeleteDevice(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.device.delete", err)
	}
	return response.SuccessJSON(c, "Device is deleted correctly.")
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
func readDevices(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	devices, total, err := deviceService.ReadDevices(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.device.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, devices})
}
