package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/reminderService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitReminders inits reminder CRUD apis
// @Title Reminders
// @Description Reminders's router group.
func InitReminders(parentRoute *echo.Group) {
	route := parentRoute.Group("/reminder")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createReminder))
	route.GET("/:id", permission.AuthRequired(readReminder))
	route.GET("/groupId", permission.AuthRequired(readReminderByGroupId))
	route.PUT("/:id", permission.AuthRequired(updateReminder))
	route.DELETE("/:id", permission.AuthRequired(deleteReminder))

	route.GET("", permission.AuthRequired(readReminders))

	reminderService.InitService()
}

// @Title createReminder
// @Description Create a reminder.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Reminder 		"Returns created reminder"
// @Failure 400 {object} response.BasicResponse "err.reminder.bind"
// @Failure 400 {object} response.BasicResponse "err.reminder.create"
// @Resource /reminders
// @Router /reminders [post]
func createReminder(c echo.Context) error {
	reminder := &model.Reminder{}
	if err := c.Bind(reminder); err != nil {
		return response.KnownErrJSON(c, "err.reminder.bind", err)
	}

	// create reminder
	reminder, err := reminderService.CreateReminder(reminder)
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.create", err)
	}

	return response.SuccessInterface(c, reminder)
}

// @Title readReminder
// @Description Read a reminder.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Reminder ID."
// @Success 200 {object} model.Reminder 		"Returns read reminder"
// @Failure 400 {object} response.BasicResponse "err.reminder.bind"
// @Failure 400 {object} response.BasicResponse "err.reminder.read"
// @Resource /reminders
// @Router /reminders/{id} [get]
func readReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	reminder, err := reminderService.ReadReminder(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.read", err)
	}

	return response.SuccessInterface(c, reminder)
}

// @Title updateReminder
// @Description Update a reminder.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Reminder ID."
// @Param   avatar      	form   	string  true	"Reminder Avatar"
// @Param   firstname		form   	string  true	"Reminder Firstname"
// @Param   lastname		form   	string  true	"Reminder Lastname"
// @Param   email	    	form   	string  true	"Reminder Email"
// @Param   birth      		form   	Time   	true	"Reminder Birth"
// @Success 200 {object} model.Reminder 		"Returns read reminder"
// @Failure 400 {object} response.BasicResponse "err.reminder.bind"
// @Failure 400 {object} response.BasicResponse "err.reminder.read"
// @Resource /reminders
// @Router /reminders/{id} [put]
func updateReminder(c echo.Context) error {
	reminder := &model.Reminder{}
	if err := c.Bind(reminder); err != nil {
		return response.KnownErrJSON(c, "err.reminder.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	reminder, err := reminderService.UpdateReminder(reminder)
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.update", err)
	}

	reminder, _ = reminderService.ReadReminder(reminder.ID)
	return response.SuccessInterface(c, reminder)
}

// @Title deleteReminder
// @Description Delete a reminder.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Reminder ID."
// @Success 200 {object} response.BasicResponse "Reminder is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.reminder.bind"
// @Failure 400 {object} response.BasicResponse "err.reminder.delete"
// @Resource /reminders
// @Router /reminders/{id} [delete]
func deleteReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete reminder with id
	err := reminderService.DeleteReminder(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.delete", err)
	}
	return response.SuccessJSON(c, "Reminder is deleted correctly.")
}

// @Title readReminders
// @Description Read reminders with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Reminder is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.reminder.read"
// @Resource /reminders
// @Router /reminders [get]
func readReminders(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read reminders with params
	reminders, total, err := reminderService.ReadReminders(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, reminders})
}

// readReminderByGroupId
func readReminderByGroupId(c echo.Context) error {
	id := c.FormValue("id")

	// read reminders with params
	reminders, total, err := reminderService.ReadReminderByGroupId(id)
	if err != nil {
		return response.KnownErrJSON(c, "err.reminder.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, reminders})
}
