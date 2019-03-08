package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/orderService"
	"../../service/worktimeService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Worktime's router group.
func InitWorktimes(parentRoute *echo.Group) {
	route := parentRoute.Group("/worktimes")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createWorktime))
	route.GET("/:id", permission.AuthRequired(readWorktime))
	route.PUT("/:id", permission.AuthRequired(updateWorktime))
	route.DELETE("/:id", permission.AuthRequired(deleteWorktime))

	route.GET("", permission.AuthRequired(readWorktimes))

	route.POST("/by/order", permission.AuthRequired(createWorktimeByOrder))

	route.GET("/date", permission.AuthRequired(readWorktimeByDate))

	worktimeService.InitService()
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
func createWorktime(c echo.Context) error {
	worktime := &model.Worktime{}
	if err := c.Bind(worktime); err != nil {
		return response.KnownErrJSON(c, "err.worktime.bind", err)
	}

	// create worktime
	worktime, err := worktimeService.CreateWorktime(worktime)
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.create", err)
	}

	return response.SuccessInterface(c, worktime)
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
func readWorktime(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	worktime, err := worktimeService.ReadWorktime(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.read", err)
	}

	return response.SuccessInterface(c, worktime)
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
func updateWorktime(c echo.Context) error {
	worktime := &model.Worktime{}
	if err := c.Bind(worktime); err != nil {
		return response.KnownErrJSON(c, "err.worktime.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	worktime, err := worktimeService.UpdateWorktime(worktime)
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.update", err)
	}

	worktime, _ = worktimeService.ReadWorktime(worktime.ID)
	return response.SuccessInterface(c, worktime)
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
func deleteWorktime(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := worktimeService.DeleteWorktime(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.delete", err)
	}
	return response.SuccessJSON(c, "Worktime is deleted correctly.")
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
func readWorktimes(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read worktimes with params
	worktimes, total, err := worktimeService.ReadWorktimes(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, worktimes})
}

func readWorktimeByDate(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))
	date := c.FormValue("date")

	// read worktimes with params
	worktimes, total, err := worktimeService.ReadWorktimesByDate(query, offset, count, field, sort, uint(userID), date)
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, worktimes})
}

func createWorktimeByOrder(c echo.Context) error {
	wo := &model.WorkTimeOrder{}
	if err := c.Bind(wo); err != nil {
		return response.KnownErrJSON(c, "err.worktime.bind", err)
	}

	// create worktime
	worktime, err := worktimeService.CreateWorktime(wo.Worktime)
	if err != nil {
		return response.KnownErrJSON(c, "err.worktime.create", err)
	}

	order, err := orderService.CreateOrder(wo.Order)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.create", err)
	}

	//return response.SuccessInterface(c, worktime)
	//return response.SuccessBindTwoInterface(c, worktime, order)
	return response.SuccessInterface(c, &model.WorkTimeOrder{worktime, order})
}
