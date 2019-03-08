package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/tblregionService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description tblregion's router group.
func InitTblregions(parentRoute *echo.Group) {
	route := parentRoute.Group("/tblregions")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createTblregion))
	route.GET("/:id", permission.AuthRequired(readTblregion))
	route.PUT("/:id", permission.AuthRequired(updateTblregion))
	route.DELETE("/:id", permission.AuthRequired(deleteTblregion))

	route.GET("", permission.AuthRequired(readTblregions))

	//	route.POST("/by/order", permission.AuthRequired(createTblregionByOrder))

	tblregionService.InitService()
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
func createTblregion(c echo.Context) error {
	tblregion := &model.Tblregion{}
	if err := c.Bind(tblregion); err != nil {
		return response.KnownErrJSON(c, "err.tblregion.bind", err)
	}

	// create tblregion
	tblregion, err := tblregionService.CreateTblregion(tblregion)
	if err != nil {
		return response.KnownErrJSON(c, "err.tblregion.create", err)
	}

	return response.SuccessInterface(c, tblregion)
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
func readTblregion(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tblregion, err := tblregionService.ReadTblregion(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.tblregion.read", err)
	}

	return response.SuccessInterface(c, tblregion)
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
func updateTblregion(c echo.Context) error {
	tblregion := &model.Tblregion{}
	if err := c.Bind(tblregion); err != nil {
		return response.KnownErrJSON(c, "err.tblregion.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	tblregion, err := tblregionService.UpdateTblregion(tblregion)
	if err != nil {
		return response.KnownErrJSON(c, "err.tblregion.update", err)
	}

	tblregion, _ = tblregionService.ReadTblregion(tblregion.ID)
	return response.SuccessInterface(c, tblregion)
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
func deleteTblregion(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := tblregionService.DeleteTblregion(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.tblregion.delete", err)
	}
	return response.SuccessJSON(c, "tblregion is deleted correctly.")
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
func readTblregions(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read tblregions with params
	tblregions, total, err := tblregionService.ReadTblregions(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.tblregion.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, tblregions})
}
