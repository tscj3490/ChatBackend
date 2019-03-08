package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/transactionService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitTransactions(parentRoute *echo.Group) {
	route := parentRoute.Group("/transactions")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createTransaction))
	route.GET("/:id", permission.AuthRequired(readTransaction))
	route.PUT("/:id", permission.AuthRequired(updateTransaction))
	route.DELETE("/:id", permission.AuthRequired(deleteTransaction))

	route.GET("", permission.AuthRequired(readTransactions))

	transactionService.InitService()
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
func createTransaction(c echo.Context) error {
	transaction := &model.Transaction{}
	if err := c.Bind(transaction); err != nil {
		return response.KnownErrJSON(c, "err.transaction.bind", err)
	}

	// create transaction
	transaction, err := transactionService.CreateTransaction(transaction)
	if err != nil {
		return response.KnownErrJSON(c, "err.transaction.create", err)
	}

	return response.SuccessInterface(c, transaction)
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
func readTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := transactionService.ReadTransaction(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.transaction.read", err)
	}

	return response.SuccessInterface(c, transaction)
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
func updateTransaction(c echo.Context) error {
	transaction := &model.Transaction{}
	if err := c.Bind(transaction); err != nil {
		return response.KnownErrJSON(c, "err.transaction.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := transactionService.UpdateTransaction(transaction)
	if err != nil {
		return response.KnownErrJSON(c, "err.transaction.update", err)
	}

	transaction, _ = transactionService.ReadTransaction(transaction.ID)
	return response.SuccessInterface(c, transaction)
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
func deleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := transactionService.DeleteTransaction(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.transaction.delete", err)
	}
	return response.SuccessJSON(c, "Transaction is deleted correctly.")
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
func readTransactions(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	transactions, total, err := transactionService.ReadTransactions(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.transaction.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, transactions})
}
