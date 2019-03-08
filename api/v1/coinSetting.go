package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/coinSettingService"
	"../response"

	"github.com/labstack/echo"
)

// InitCoinSettings inits coinSetting CRUD apis
// @Title CoinSettings
// @Description CoinSettings's router group.
func InitCoinSettings(parentRoute *echo.Group) {
	route := parentRoute.Group("/coinSettings")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createCoinSetting))
	route.GET("/:id", permission.AuthRequired(readCoinSetting))
	route.PUT("/:id", permission.AuthRequired(updateCoinSetting))
	route.DELETE("/:id", permission.AuthRequired(deleteCoinSetting))

	route.GET("", permission.AuthRequired(readCoinSettings))

	coinSettingService.InitService()
}

// @Title createCoinSetting
// @Description Create a coinSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.CoinSetting 		"Returns created coinSetting"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.create"
// @Resource /coinSettings
// @Router /coinSettings [post]
func createCoinSetting(c echo.Context) error {
	coinSetting := &model.CoinSetting{}
	if err := c.Bind(coinSetting); err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.bind", err)
	}

	// create coinSetting
	coinSetting, err := coinSettingService.CreateCoinSetting(coinSetting)
	if err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.create", err)
	}

	return response.SuccessInterface(c, coinSetting)
}

// @Title readCoinSetting
// @Description Read a coinSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"CoinSetting ID."
// @Success 200 {object} model.CoinSetting 		"Returns read coinSetting"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.read"
// @Resource /coinSettings
// @Router /coinSettings/{id} [get]
func readCoinSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	coinSetting, err := coinSettingService.ReadCoinSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.read", err)
	}

	return response.SuccessInterface(c, coinSetting)
}

// @Title updateCoinSetting
// @Description Update a coinSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"CoinSetting ID."
// @Param   avatar      	form   	string  true	"CoinSetting Avatar"
// @Param   firstname		form   	string  true	"CoinSetting Firstname"
// @Param   lastname		form   	string  true	"CoinSetting Lastname"
// @Param   email	    	form   	string  true	"CoinSetting Email"
// @Param   birth      		form   	Time   	true	"CoinSetting Birth"
// @Success 200 {object} model.CoinSetting 		"Returns read coinSetting"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.read"
// @Resource /coinSettings
// @Router /coinSettings/{id} [put]
func updateCoinSetting(c echo.Context) error {
	coinSetting := &model.CoinSetting{}
	if err := c.Bind(coinSetting); err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	coinSetting, err := coinSettingService.UpdateCoinSetting(coinSetting)
	if err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.update", err)
	}

	coinSetting, _ = coinSettingService.ReadCoinSetting(coinSetting.ID)
	return response.SuccessInterface(c, coinSetting)
}

// @Title deleteCoinSetting
// @Description Delete a coinSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"CoinSetting ID."
// @Success 200 {object} response.BasicResponse "CoinSetting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.coinSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.coinSetting.delete"
// @Resource /coinSettings
// @Router /coinSettings/{id} [delete]
func deleteCoinSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete coinSetting with id
	err := coinSettingService.DeleteCoinSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.delete", err)
	}
	return response.SuccessJSON(c, "CoinSetting is deleted correctly.")
}

// @Title readCoinSettings
// @Description Read coinSettings with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"CoinSetting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.coinSetting.read"
// @Resource /coinSettings
// @Router /coinSettings [get]
func readCoinSettings(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read coinSettings with params
	coinSettings, total, err := coinSettingService.ReadCoinSettings(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.coinSetting.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, coinSettings})
}
