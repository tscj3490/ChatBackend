package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/orderService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitOrders(parentRoute *echo.Group) {
	route := parentRoute.Group("/orders")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createOrder))
	route.GET("/:id", permission.AuthRequired(readOrder))
	route.PUT("/:id", permission.AuthRequired(updateOrder))
	route.DELETE("/:id", permission.AuthRequired(deleteOrder))

	route.GET("", permission.AuthRequired(readOrders))

	route.GET("/field", permission.AuthRequired(readOrdersByField))
	route.GET("/filter", permission.AuthRequired(readOrdersByFilter))

	route.GET("/customer", permission.AuthRequired(readOrdersByCustomerStatus))
	route.GET("/vender", permission.AuthRequired(readOrdersByVenderStatus))

	route.POST("/customer", permission.AuthRequired(CreateOrdersByCustomer))
	route.GET("/bookdate", permission.AuthRequired(readOrdersByBookdate))

	route.POST("/job", permission.AuthRequired(readOrdersByJob))
	route.PUT("/update_status", permission.AuthRequired(updateOrderByStatus))
	parentRoute.Group("/timetable").GET("/date", permission.AuthRequired(readOrdersByDate))

	orderService.InitService()
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
// func createOrder(c echo.Context) error {
// 	order := &model.Order{}
// 	if err := c.Bind(order); err != nil {
// 		return response.KnownErrJSON(c, "err.order.bind", err)
// 	}

// 	// create order
// 	order, err := orderService.CreateOrder(order)
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.order.create", err)
// 	}

// 	return response.SuccessInterface(c, order)
// }

func createOrder(c echo.Context) error {
	order := &model.Order{}
	if err := c.Bind(order); err != nil {
		return response.KnownErrJSON(c, "err.order.bind", err)
	}

	// create order
	order, err := orderService.CreateOrder(order)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.create", err)
	}

	return response.SuccessInterface(c, order)
}

// CreateOrdersByCustomer
func CreateOrdersByCustomer(c echo.Context) error {
	cbo := &model.BookingInfo{}

	// customer := &model.Customer{}
	if err := c.Bind(cbo); err != nil {
		return response.KnownErrJSON(c, "err.cbo.bind", err)
	}

	// create order
	da, err := orderService.SetOrderByCustomer(cbo)
	if err != nil {

		return response.KnownErrJSON(c, "err.cbo.create", err)
	}

	// create user with registered info
	//	order, err := orderService.CreateBookWithEmail(cbo)
	//	fmt.Println(order)
	return response.SuccessInterface(c, da)
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
func readOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := orderService.ReadOrder(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, order)
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
func updateOrder(c echo.Context) error {
	order := &model.Order{}
	if err := c.Bind(order); err != nil {
		return response.KnownErrJSON(c, "err.order.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	order, err := orderService.UpdateOrder(order)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.update", err)
	}

	order, _ = orderService.ReadOrder(order.ID)
	return response.SuccessInterface(c, order)
}

//  updateOrderByStatus
func updateOrderByStatus(c echo.Context) error {
	us := &model.UpdateStatus{}
	logs := &model.Joblog{}
	if err := c.Bind(us); err != nil {
		return response.KnownErrJSON(c, "err.update_status.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	logs, err := orderService.UpdateOrderByStatus(us)
	if err != nil {
		return response.KnownErrJSON(c, "err.logs.update", err)
	}

	//	logs, _ = orderService.ReadOrderByStatus(order.ID)
	return response.SuccessInterface(c, logs)
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
func deleteOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := orderService.DeleteOrder(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.order.delete", err)
	}
	return response.SuccessJSON(c, "Order is deleted correctly.")
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
func readOrders(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	orders, total, err := orderService.ReadOrders(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, orders})
}

func readOrdersByField(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read orders with params
	orders, total, err := orderService.ReadByField(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, orders})
}

func readOrdersByDate(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendorId"))
	date := c.FormValue("date")

	// read orders with params
	timeinfos, total, err := orderService.ReadByDate(query, offset, count, field, sort, uint(vendorID), date)
	if err != nil {
		return response.KnownErrJSON(c, "err.timeinfo.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, timeinfos})
}

func readOrdersByBookdate(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	vendorID, _ := strconv.Atoi(c.FormValue("vendorId"))
	customerID, _ := strconv.Atoi(c.FormValue("customerId"))
	date := c.FormValue("date")

	// read orders with params
	timeinfos, total, err := orderService.ReadByBookdate(query, offset, count, field, sort, uint(vendorID), uint(customerID), date)
	if err != nil {
		return response.KnownErrJSON(c, "err.timeinfo.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, timeinfos})
}

func readOrdersByFilter(c echo.Context) error {
	sort := c.FormValue("sort")
	// offset, _ := strconv.Atoi(c.FormValue("offset"))
	// count, _ := strconv.Atoi(c.FormValue("count"))
	//field := c.FormValue("field")
	SortDirection, _ := strconv.Atoi(c.FormValue("SortDirection"))
	// userID, _ := strconv.Atoi(c.FormValue("userId"))

	// SearchContent := c.FormValue("SearchContent")
	// SearchType := c.FormValue("SearchType")
	StatusFilter, _ := strconv.Atoi(c.FormValue("StatusFilter"))
	DeviceFilter, _ := strconv.Atoi(c.FormValue("DeviceFilter"))

	S := c.FormValue("s_name")
	C := c.FormValue("c_name")
	V := c.FormValue("v_name")

	// read orders with params
	orders, total, err := orderService.ReadByFilter(S, C, V, sort, SortDirection, StatusFilter, DeviceFilter)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, orders})
}

func readOrdersByVenderStatus(c echo.Context) error {
	sort := c.FormValue("sort")
	SortDirection, _ := strconv.Atoi(c.FormValue("SortDirection"))

	StatusFilter, _ := strconv.Atoi(c.FormValue("StatusFilter"))
	VenderFilter, _ := strconv.Atoi(c.FormValue("VenderFilter"))

	// read orders with params
	orders, total, err := orderService.ReadByVenderStatus(sort, SortDirection, StatusFilter, VenderFilter)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, orders})
}

func readOrdersByCustomerStatus(c echo.Context) error {
	sort := c.FormValue("sort")
	SortDirection, _ := strconv.Atoi(c.FormValue("SortDirection"))

	StatusFilter, _ := strconv.Atoi(c.FormValue("StatusFilter"))
	CustomerFilter, _ := strconv.Atoi(c.FormValue("CustomerFilter"))

	// read orders with params
	orders, total, err := orderService.ReadByCustomerStatus(sort, SortDirection, StatusFilter, CustomerFilter)
	if err != nil {
		return response.KnownErrJSON(c, "err.order.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, orders})
}

func readOrdersByJob(c echo.Context) error {
	jobs := &model.JobsInfo{}
	//	customers := []model.Customer{}
	if err := c.Bind(jobs); err != nil {
		return response.KnownErrJSON(c, "err.jobs.bind", err)
	}
	// read customers with params
	bookinginfos, total, err := orderService.ReadByJob(jobs)
	if err != nil {
		return response.KnownErrJSON(c, "err.customers.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, bookinginfos})
}
