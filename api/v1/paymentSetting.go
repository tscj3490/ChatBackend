package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/paymentSettingService"
	"../response"

	"github.com/labstack/echo"
)

// InitPaymentSettings inits paymentSetting CRUD apis
// @Title PaymentSettings
// @Description PaymentSettings's router group.
func InitPaymentSettings(parentRoute *echo.Group) {
	route := parentRoute.Group("/paymentSettings")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createPaymentSetting))
	route.GET("/:id", permission.AuthRequired(readPaymentSetting))
	route.PUT("/:id", permission.AuthRequired(updatePaymentSetting))
	route.DELETE("/:id", permission.AuthRequired(deletePaymentSetting))

	route.GET("", permission.AuthRequired(readPaymentSettings))

	paymentSettingService.InitService()
}

// @Title createPaymentSetting
// @Description Create a paymentSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.PaymentSetting 		"Returns created paymentSetting"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.create"
// @Resource /paymentSettings
// @Router /paymentSettings [post]
func createPaymentSetting(c echo.Context) error {
	paymentSetting := &model.PaymentSetting{}
	if err := c.Bind(paymentSetting); err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.bind", err)
	}

	// create paymentSetting
	paymentSetting, err := paymentSettingService.CreatePaymentSetting(paymentSetting)
	if err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.create", err)
	}

	return response.SuccessInterface(c, paymentSetting)
}

// @Title readPaymentSetting
// @Description Read a paymentSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"PaymentSetting ID."
// @Success 200 {object} model.PaymentSetting 		"Returns read paymentSetting"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.read"
// @Resource /paymentSettings
// @Router /paymentSettings/{id} [get]
func readPaymentSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paymentSetting, err := paymentSettingService.ReadPaymentSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.read", err)
	}

	return response.SuccessInterface(c, paymentSetting)
}

// @Title updatePaymentSetting
// @Description Update a paymentSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"PaymentSetting ID."
// @Param   avatar      	form   	string  true	"PaymentSetting Avatar"
// @Param   firstname		form   	string  true	"PaymentSetting Firstname"
// @Param   lastname		form   	string  true	"PaymentSetting Lastname"
// @Param   email	    	form   	string  true	"PaymentSetting Email"
// @Param   birth      		form   	Time   	true	"PaymentSetting Birth"
// @Success 200 {object} model.PaymentSetting 		"Returns read paymentSetting"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.read"
// @Resource /paymentSettings
// @Router /paymentSettings/{id} [put]
func updatePaymentSetting(c echo.Context) error {
	paymentSetting := &model.PaymentSetting{}
	if err := c.Bind(paymentSetting); err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	paymentSetting, err := paymentSettingService.UpdatePaymentSetting(paymentSetting)
	if err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.update", err)
	}

	paymentSetting, _ = paymentSettingService.ReadPaymentSetting(paymentSetting.ID)
	return response.SuccessInterface(c, paymentSetting)
}

// @Title deletePaymentSetting
// @Description Delete a paymentSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"PaymentSetting ID."
// @Success 200 {object} response.BasicResponse "PaymentSetting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.delete"
// @Resource /paymentSettings
// @Router /paymentSettings/{id} [delete]
func deletePaymentSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete paymentSetting with id
	err := paymentSettingService.DeletePaymentSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.delete", err)
	}
	return response.SuccessJSON(c, "PaymentSetting is deleted correctly.")
}

// @Title readPaymentSettings
// @Description Read paymentSettings with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"PaymentSetting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.paymentSetting.read"
// @Resource /paymentSettings
// @Router /paymentSettings [get]
func readPaymentSettings(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read paymentSettings with params
	paymentSettings, total, err := paymentSettingService.ReadPaymentSettings(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.paymentSetting.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, paymentSettings})
}
