package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/messageService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitMessages(parentRoute *echo.Group) {
	route := parentRoute.Group("/messages")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createMessage))
	route.GET("/:id", permission.AuthRequired(readMessage))
	route.PUT("/:id", permission.AuthRequired(updateMessage))
	route.DELETE("/:id", permission.AuthRequired(deleteMessage))

	route.GET("", permission.AuthRequired(readMessages))

	messageService.InitService()
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
func createMessage(c echo.Context) error {
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		return response.KnownErrJSON(c, "err.message.bind", err)
	}

	// create message
	message, err := messageService.CreateMessage(message)
	if err != nil {
		return response.KnownErrJSON(c, "err.message.create", err)
	}

	return response.SuccessInterface(c, message)
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
func readMessage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	message, err := messageService.ReadMessage(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.message.read", err)
	}

	return response.SuccessInterface(c, message)
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
func updateMessage(c echo.Context) error {
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		return response.KnownErrJSON(c, "err.message.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	message, err := messageService.UpdateMessage(message)
	if err != nil {
		return response.KnownErrJSON(c, "err.message.update", err)
	}

	message, _ = messageService.ReadMessage(message.ID)
	return response.SuccessInterface(c, message)
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
func deleteMessage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete device with id
	err := messageService.DeleteMessage(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.message.delete", err)
	}
	return response.SuccessJSON(c, "Message is deleted correctly.")
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
func readMessages(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read messages with params
	messages, total, err := messageService.ReadMessages(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.message.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, messages})
}
