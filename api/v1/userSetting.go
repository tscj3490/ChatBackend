package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/userSettingService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitUserSettings inits userSetting CRUD apis
// @Title UserSettings
// @Description UserSettings's router group.
func InitUserSettings(parentRoute *echo.Group) {
	route := parentRoute.Group("/userSettings")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(upsertUserSetting))
	route.DELETE("/:userId", permission.AuthRequired(deleteUserSetting))

	route.GET("", permission.AuthRequired(readUserSettings))

	userSettingService.InitService()
}

// @Title createUserSetting
// @Description Create a userSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.PublicUserSetting 	"Returns created userSetting"
// @Failure 400 {object} response.BasicResponse 	"err.userSetting.bind"
// @Failure 400 {object} response.BasicResponse 	"err.userSetting.create"
// @Resource /userSettings
// @Router /userSettings [post]
func upsertUserSetting(c echo.Context) error {
	userSetting := &model.UserSetting{}
	if err := c.Bind(userSetting); err != nil {
		return response.KnownErrJSON(c, "err.userSetting.bind", err)
	}

	// create userSetting
	userSetting, err := userSettingService.UpsertUserSetting(userSetting)
	if err != nil {
		return response.KnownErrJSON(c, "err.userSetting.upsert", err)
	}

	return response.SuccessInterface(c, userSetting)
}

// @Title deleteUserSetting
// @Description Delete a userSetting.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"UserSetting ID."
// @Success 200 {object} response.BasicResponse "UserSetting is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.userSetting.bind"
// @Failure 400 {object} response.BasicResponse "err.userSetting.delete"
// @Resource /userSettings
// @Router /userSettings/{id} [delete]
func deleteUserSetting(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("userId"))
	// delete userSetting with userId
	err := userSettingService.DeleteUserSetting(uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.userSetting.delete", err)
	}
	return response.SuccessJSON(c, "UserSetting is deleted correctly.")
}

// @Title readUserSettings
// @Description Read userSettings with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   userID			form   	string  false	"Will search string."
// @Success 200 {object} UserSettingsForm 			"Returned list userSettings."
// @Failure 400 {object} response.BasicResponse "err.userSetting.read"
// @Resource /userSettings
// @Router /userSettings [get]
func readUserSettings(c echo.Context) error {
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read userSettings with params
	userSettings, err := userSettingService.ReadUserSettings(uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.userSetting.read", err)
	}

	return response.SuccessInterface(c, userSettings)
}
