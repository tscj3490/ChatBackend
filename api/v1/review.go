package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/reviewService"
	"../response"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

// InitDevices inits device CRUD apis
// @Title Devices
// @Description Devices's router group.
func InitReviews(parentRoute *echo.Group) {
	route := parentRoute.Group("/reviews")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createReview))
	route.GET("/:id", permission.AuthRequired(readReview))
	route.PUT("/:id", permission.AuthRequired(updateReview))
	route.DELETE("/:id", permission.AuthRequired(deleteReview))

	route.GET("", permission.AuthRequired(readReviews))

	reviewService.InitService()
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
func createReview(c echo.Context) error {
	review := &model.Review{}
	if err := c.Bind(review); err != nil {
		return response.KnownErrJSON(c, "err.review.bind", err)
	}

	// create review
	review, err := reviewService.CreateReview(review)
	if err != nil {
		return response.KnownErrJSON(c, "err.review.create", err)
	}

	return response.SuccessInterface(c, review)
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
func readReview(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	review, err := reviewService.ReadReview(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.review.read", err)
	}

	return response.SuccessInterface(c, review)
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
func updateReview(c echo.Context) error {
	review := &model.Review{}
	if err := c.Bind(review); err != nil {
		return response.KnownErrJSON(c, "err.review.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	review, err := reviewService.UpdateReview(review)
	if err != nil {
		return response.KnownErrJSON(c, "err.review.update", err)
	}

	review, _ = reviewService.ReadReview(review.ID)
	return response.SuccessInterface(c, review)
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
func deleteReview(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete order with id
	err := reviewService.DeleteReview(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.review.delete", err)
	}
	return response.SuccessJSON(c, "Review is deleted correctly.")
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
func readReviews(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read devices with params
	reviews, total, err := reviewService.ReadReviews(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.review.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, reviews})
}
