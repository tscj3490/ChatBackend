package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/specialtimeService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description specialtime's router group.
func InitSpecialtimes(parentRoute *echo.Group) {
	route := parentRoute.Group("/specialtimes")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createSpecialtime))
	route.GET("/:id", permission.AuthRequired(readSpecialtime))
	route.PUT("/:id", permission.AuthRequired(updateSpecialtime))
	route.DELETE("/:id", permission.AuthRequired(deleteSpecialtime))

	route.GET("", permission.AuthRequired(readSpecialtimes))

	//	route.POST("/by/order", permission.AuthRequired(createSpecialtimeByOrder))

	route.GET("/date", permission.AuthRequired(readSpecialtimeByDate))

	specialtimeService.InitService()
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
func createSpecialtime(c echo.Context) error {
	specialtime := &model.Specialtime{}
	if err := c.Bind(specialtime); err != nil {
		return response.KnownErrJSON(c, "err.specialtime.bind", err)
	}

	// create specialtime
	specialtime, err := specialtimeService.CreateSpecialtime(specialtime)
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.create", err)
	}

	return response.SuccessInterface(c, specialtime)
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
func readSpecialtime(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	specialtime, err := specialtimeService.ReadSpecialtime(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.read", err)
	}

	return response.SuccessInterface(c, specialtime)
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
func updateSpecialtime(c echo.Context) error {
	specialtime := &model.Specialtime{}
	if err := c.Bind(specialtime); err != nil {
		return response.KnownErrJSON(c, "err.specialtime.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	specialtime, err := specialtimeService.UpdateSpecialtime(specialtime)
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.update", err)
	}

	specialtime, _ = specialtimeService.ReadSpecialtime(specialtime.ID)
	return response.SuccessInterface(c, specialtime)
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
func deleteSpecialtime(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := specialtimeService.DeleteSpecialtime(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.delete", err)
	}
	return response.SuccessJSON(c, "specialtime is deleted correctly.")
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
func readSpecialtimes(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read specialtimes with params
	specialtimes, total, err := specialtimeService.ReadSpecialtimes(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, specialtimes})
}

func readSpecialtimeByDate(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))
	date := c.FormValue("date")

	// read specialtimes with params
	specialtimes, total, err := specialtimeService.ReadSpecialtimesByDate(query, offset, count, field, sort, uint(userID), date)
	if err != nil {
		return response.KnownErrJSON(c, "err.specialtime.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, specialtimes})
}

// func createSpecialtimeByOrder(c echo.Context) error {
// 	wo := &model.Specialtime{}
// 	if err := c.Bind(wo); err != nil {
// 		return response.KnownErrJSON(c, "err.specialtime.bind", err)
// 	}

// 	// create specialtime
// 	specialtime, err := specialtimeService.CreateSpecialtime(wo.Specialtime)
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.specialtime.create", err)
// 	}

// 	order, err := orderService.CreateOrder(wo.Order)
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.order.create", err)
// 	}

// 	//return response.SuccessInterface(c, specialtime)
// 	//return response.SuccessBindTwoInterface(c, specialtime, order)
// 	return response.SuccessInterface(c, &model.Specialtime{specialtime, order})
// }
