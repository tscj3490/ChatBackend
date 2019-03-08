package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/serviceService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitServices(parentRoute *echo.Group) {
	route := parentRoute.Group("/services")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createService))
	route.GET("/:id", permission.AuthRequired(readService))
	route.PUT("/:id", permission.AuthRequired(updateService))
	route.DELETE("/:id", permission.AuthRequired(deleteService))

	route.GET("", permission.AuthRequired(readServices))

	route.GET("/field", permission.AuthRequired(readServicesByField))

	serviceService.InitService()
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
func createService(c echo.Context) error {
	service := &model.Service{}
	if err := c.Bind(service); err != nil {
		return response.KnownErrJSON(c, "err.service.bind", err)
	}

	// create service
	service, err := serviceService.CreateService(service)
	if err != nil {
		return response.KnownErrJSON(c, "err.service.create", err)
	}

	return response.SuccessInterface(c, service)
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
func readService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	service, err := serviceService.ReadService(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.service.read", err)
	}

	return response.SuccessInterface(c, service)
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
func updateService(c echo.Context) error {
	service := &model.Service{}
	if err := c.Bind(service); err != nil {
		return response.KnownErrJSON(c, "err.service.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	service, err := serviceService.UpdateService(service)
	if err != nil {
		return response.KnownErrJSON(c, "err.service.update", err)
	}

	service, _ = serviceService.ReadService(service.ID)
	return response.SuccessInterface(c, service)
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
func deleteService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := serviceService.DeleteService(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.service.delete", err)
	}
	return response.SuccessJSON(c, "Service is deleted correctly.")
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
func readServices(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	services, total, err := serviceService.ReadServices(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.service.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, services})
}

func readServicesByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	services, total, err := serviceService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.service.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, services})
}
