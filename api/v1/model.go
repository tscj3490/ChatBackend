package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/modelService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitModels(parentRoute *echo.Group) {
	route := parentRoute.Group("/models")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createModel))
	route.GET("/:id", permission.AuthRequired(readModel))
	route.PUT("/:id", permission.AuthRequired(updateModel))
	route.DELETE("/:id", permission.AuthRequired(deleteModel))

	route.GET("", permission.AuthRequired(readModels))
	route.GET("/field", permission.AuthRequired(readModelsByField))
	modelService.InitService()
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
func createModel(c echo.Context) error {
	model := &model.Model{}
	if err := c.Bind(model); err != nil {
		return response.KnownErrJSON(c, "err.model.bind", err)
	}

	// create model
	model, err := modelService.CreateModel(model)
	if err != nil {
		return response.KnownErrJSON(c, "err.model.create", err)
	}

	return response.SuccessInterface(c, model)
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
func readModel(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	model, err := modelService.ReadModel(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.model.read", err)
	}

	return response.SuccessInterface(c, model)
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
func updateModel(c echo.Context) error {
	model := &model.Model{}
	if err := c.Bind(model); err != nil {
		return response.KnownErrJSON(c, "err.model.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	model, err := modelService.UpdateModel(model)
	if err != nil {
		return response.KnownErrJSON(c, "err.model.update", err)
	}

	model, _ = modelService.ReadModel(model.ID)
	return response.SuccessInterface(c, model)
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
func deleteModel(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete model with id
	err := modelService.DeleteModel(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.model.delete", err)
	}
	return response.SuccessJSON(c, "Model is deleted correctly.")
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
func readModels(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read models with params
	models, total, err := modelService.ReadModels(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.model.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, models})
}

func readModelsByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read models with params
	models, total, err := modelService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.model.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, models})
}
