package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/settingService"
	"../response"

	"github.com/labstack/echo"
)

// InitSettings inits setting CRUD apis
// @Title Settings
// @Description Settings's router group.
func InitSettings(parentRoute *echo.Group) {
	route := parentRoute.Group("/settings")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createSetting))
	route.GET("/:id", permission.AuthRequired(readSetting))
	route.PUT("/:id", permission.AuthRequired(updateSetting))
	route.DELETE("/:id", permission.AuthRequired(deleteSetting))

	route.GET("", permission.AuthRequired(readSettings))

	settingService.InitService()
}

// @Title createSetting
// @Description Create a setting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Setting 		"Returns created setting"
// @Failure 400 {object} response.BasicResponse "err.setting.bind"
// @Failure 400 {object} response.BasicResponse "err.setting.create"
// @Resource /settings
// @Router /settings [post]
func createSetting(c echo.Context) error {
	setting := &model.Setting{}
	if err := c.Bind(setting); err != nil {
		return response.KnownErrJSON(c, "err.setting.bind", err)
	}

	// create setting
	setting, err := settingService.CreateSetting(setting)
	if err != nil {
		return response.KnownErrJSON(c, "err.setting.create", err)
	}

	return response.SuccessInterface(c, setting)
}

// @Title readSetting
// @Description Read a setting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Setting ID."
// @Success 200 {object} model.Setting 		"Returns read setting"
// @Failure 400 {object} response.BasicResponse "err.setting.bind"
// @Failure 400 {object} response.BasicResponse "err.setting.read"
// @Resource /settings
// @Router /settings/{id} [get]
func readSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	setting, err := settingService.ReadSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.setting.read", err)
	}

	return response.SuccessInterface(c, setting)
}

// @Title updateSetting
// @Description Update a setting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Setting ID."
// @Param   avatar      	form   	string  true	"Setting Avatar"
// @Param   firstname		form   	string  true	"Setting Firstname"
// @Param   lastname		form   	string  true	"Setting Lastname"
// @Param   email	    	form   	string  true	"Setting Email"
// @Param   birth      		form   	Time   	true	"Setting Birth"
// @Success 200 {object} model.Setting 		"Returns read setting"
// @Failure 400 {object} response.BasicResponse "err.setting.bind"
// @Failure 400 {object} response.BasicResponse "err.setting.read"
// @Resource /settings
// @Router /settings/{id} [put]
func updateSetting(c echo.Context) error {
	setting := &model.Setting{}
	if err := c.Bind(setting); err != nil {
		return response.KnownErrJSON(c, "err.setting.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	setting, err := settingService.UpdateSetting(setting)
	if err != nil {
		return response.KnownErrJSON(c, "err.setting.update", err)
	}

	setting, _ = settingService.ReadSetting(setting.ID)
	return response.SuccessInterface(c, setting)
}

// @Title deleteSetting
// @Description Delete a setting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Setting ID."
// @Success 200 {object} response.BasicResponse "Setting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.setting.bind"
// @Failure 400 {object} response.BasicResponse "err.setting.delete"
// @Resource /settings
// @Router /settings/{id} [delete]
func deleteSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete setting with id
	err := settingService.DeleteSetting(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.setting.delete", err)
	}
	return response.SuccessJSON(c, "Setting is deleted correctly.")
}

// @Title readSettings
// @Description Read settings with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Setting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.setting.read"
// @Resource /settings
// @Router /settings [get]
func readSettings(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read settings with params
	settings, total, err := settingService.ReadSettings(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.setting.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, settings})
}
