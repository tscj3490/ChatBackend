package v1

import (
	"strconv"

	"../../model"
	"../../service/authService/permission"
	"../../service/joblogService"
	"../response"

	"github.com/labstack/echo"
)

// InitJoblogs inits joblog CRUD apis
// @Title Joblogs
// @Description Joblogs's router group.
func InitJoblogs(parentRoute *echo.Group) {
	route := parentRoute.Group("/joblogs")
	//	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createJoblog))
	route.GET("/:id", permission.AuthRequired(readJoblog))
	route.PUT("/:id", permission.AuthRequired(updateJoblog))
	route.DELETE("/:id", permission.AuthRequired(deleteJoblog))

	route.GET("", permission.AuthRequired(readJoblogs))

	joblogService.InitService()
}

// @Title createJoblog
// @Description Create a joblog.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Joblog 		"Returns created joblog"
// @Failure 400 {object} response.BasicResponse "err.joblog.bind"
// @Failure 400 {object} response.BasicResponse "err.joblog.create"
// @Resource /joblogs
// @Router /joblogs [post]
func createJoblog(c echo.Context) error {
	joblog := &model.Joblog{}
	if err := c.Bind(joblog); err != nil {
		return response.KnownErrJSON(c, "err.joblog.bind", err)
	}

	// create joblog
	joblog, err := joblogService.CreateJoblog(joblog)
	if err != nil {
		return response.KnownErrJSON(c, "err.joblog.create", err)
	}

	return response.SuccessInterface(c, joblog)
}

// @Title readJoblog
// @Description Read a joblog.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Joblog ID."
// @Success 200 {object} model.Joblog 		"Returns read joblog"
// @Failure 400 {object} response.BasicResponse "err.joblog.bind"
// @Failure 400 {object} response.BasicResponse "err.joblog.read"
// @Resource /joblogs
// @Router /joblogs/{id} [get]
func readJoblog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	joblog, err := joblogService.ReadJoblog(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.joblog.read", err)
	}

	return response.SuccessInterface(c, joblog)
}

// @Title updateJoblog
// @Description Update a joblog.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Joblog ID."
// @Param   avatar      	form   	string  true	"Joblog Avatar"
// @Param   firstname		form   	string  true	"Joblog Firstname"
// @Param   lastname		form   	string  true	"Joblog Lastname"
// @Param   email	    	form   	string  true	"Joblog Email"
// @Param   birth      		form   	Time   	true	"Joblog Birth"
// @Success 200 {object} model.Joblog 		"Returns read joblog"
// @Failure 400 {object} response.BasicResponse "err.joblog.bind"
// @Failure 400 {object} response.BasicResponse "err.joblog.read"
// @Resource /joblogs
// @Router /joblogs/{id} [put]
func updateJoblog(c echo.Context) error {
	joblog := &model.Joblog{}
	if err := c.Bind(joblog); err != nil {
		return response.KnownErrJSON(c, "err.joblog.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	joblog, err := joblogService.UpdateJoblog(joblog)
	if err != nil {
		return response.KnownErrJSON(c, "err.joblog.update", err)
	}

	joblog, _ = joblogService.ReadJoblog(joblog.ID)
	return response.SuccessInterface(c, joblog)
}

// @Title deleteJoblog
// @Description Delete a joblog.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Joblog ID."
// @Success 200 {object} response.BasicResponse "Joblog is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.joblog.bind"
// @Failure 400 {object} response.BasicResponse "err.joblog.delete"
// @Resource /joblogs
// @Router /joblogs/{id} [delete]
func deleteJoblog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("orderId"))
	// delete joblog with id
	err := joblogService.DeleteJoblog(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.joblog.delete", err)
	}
	return response.SuccessJSON(c, "Joblog is deleted correctly.")
}

// @Title readJoblogs
// @Description Read joblogs with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Joblog is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.joblog.read"
// @Resource /joblogs
// @Router /joblogs [get]
func readJoblogs(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("orderId"))

	// read joblogs with params
	joblogs, total, err := joblogService.ReadJoblogs(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.joblog.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, joblogs})
}
