package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/orderServiceService"
	"../response"

	"github.com/labstack/echo"
)

// InitOrderServices inits orderService CRUD apis
// @Title OrderServices
// @Description OrderServices's router group.
func InitOrderServices(parentRoute *echo.Group) {
	route := parentRoute.Group("/orderServices")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createOrderService))
	route.GET("/:id", permission.AuthRequired(readOrderService))
	route.PUT("/:id", permission.AuthRequired(updateOrderService))
	route.DELETE("/:id", permission.AuthRequired(deleteOrderService))

	route.GET("", permission.AuthRequired(readOrderServices))

	orderServiceService.InitService()
}

// @Title createOrderService
// @Description Create a orderService.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.OrderService 		"Returns created orderService"
// @Failure 400 {object} response.BasicResponse "err.orderService.bind"
// @Failure 400 {object} response.BasicResponse "err.orderService.create"
// @Resource /orderServices
// @Router /orderServices [post]
func createOrderService(c echo.Context) error {
	orderService := &model.OrderService{}
	if err := c.Bind(orderService); err != nil {
		return response.KnownErrJSON(c, "err.orderService.bind", err)
	}

	// create orderService
	orderService, err := orderServiceService.CreateOrderService(orderService)
	if err != nil {
		return response.KnownErrJSON(c, "err.orderService.create", err)
	}

	return response.SuccessInterface(c, orderService)
}

// @Title readOrderService
// @Description Read a orderService.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"OrderService ID."
// @Success 200 {object} model.OrderService 		"Returns read orderService"
// @Failure 400 {object} response.BasicResponse "err.orderService.bind"
// @Failure 400 {object} response.BasicResponse "err.orderService.read"
// @Resource /orderServices
// @Router /orderServices/{id} [get]
func readOrderService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	orderService, err := orderServiceService.ReadOrderService(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.orderService.read", err)
	}

	return response.SuccessInterface(c, orderService)
}

// @Title updateOrderService
// @Description Update a orderService.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"OrderService ID."
// @Param   avatar      	form   	string  true	"OrderService Avatar"
// @Param   firstname		form   	string  true	"OrderService Firstname"
// @Param   lastname		form   	string  true	"OrderService Lastname"
// @Param   email	    	form   	string  true	"OrderService Email"
// @Param   birth      		form   	Time   	true	"OrderService Birth"
// @Success 200 {object} model.OrderService 		"Returns read orderService"
// @Failure 400 {object} response.BasicResponse "err.orderService.bind"
// @Failure 400 {object} response.BasicResponse "err.orderService.read"
// @Resource /orderServices
// @Router /orderServices/{id} [put]
func updateOrderService(c echo.Context) error {
	orderService := &model.OrderService{}
	if err := c.Bind(orderService); err != nil {
		return response.KnownErrJSON(c, "err.orderService.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	orderService, err := orderServiceService.UpdateOrderService(orderService)
	if err != nil {
		return response.KnownErrJSON(c, "err.orderService.update", err)
	}

	orderService, _ = orderServiceService.ReadOrderService(orderService.ID)
	return response.SuccessInterface(c, orderService)
}

// @Title deleteOrderService
// @Description Delete a orderService.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"OrderService ID."
// @Success 200 {object} response.BasicResponse "OrderService is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.orderService.bind"
// @Failure 400 {object} response.BasicResponse "err.orderService.delete"
// @Resource /orderServices
// @Router /orderServices/{id} [delete]
func deleteOrderService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete orderService with id
	err := orderServiceService.DeleteOrderService(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.orderService.delete", err)
	}
	return response.SuccessJSON(c, "OrderService is deleted correctly.")
}

// @Title readOrderServices
// @Description Read orderServices with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"OrderService is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.orderService.read"
// @Resource /orderServices
// @Router /orderServices [get]
func readOrderServices(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	orderID, _ := strconv.Atoi(c.FormValue("orderId"))

	// read servicelist with params
	servicelist, total, err := orderServiceService.ReadOrderServices(query, offset, count, field, sort, uint(orderID))
	if err != nil {
		return response.KnownErrJSON(c, "err.orderService.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, servicelist})
}
