package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/teamService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitTeams inits team CRUD apis
// @Title Teams
// @Description Teams's router group.
func InitTeams(parentRoute *echo.Group) {
	route := parentRoute.Group("/teams")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createTeam))
	route.GET("/:id", permission.AuthRequired(readTeam))
	route.PUT("/:id", permission.AuthRequired(updateTeam))
	route.DELETE("/:id", permission.AuthRequired(deleteTeam))

	route.GET("", permission.AuthRequired(readTeams))

	teamService.InitService()
}

// @Title createTeam
// @Description Create a team.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Team 		"Returns created team"
// @Failure 400 {object} response.BasicResponse "err.team.bind"
// @Failure 400 {object} response.BasicResponse "err.team.create"
// @Resource /teams
// @Router /teams [post]
func createTeam(c echo.Context) error {
	team := &model.Team{}
	if err := c.Bind(team); err != nil {
		return response.KnownErrJSON(c, "err.team.bind", err)
	}

	// create team
	team, err := teamService.CreateTeam(team)
	if err != nil {
		return response.KnownErrJSON(c, "err.team.create", err)
	}

	return response.SuccessInterface(c, team)
}

// @Title readTeam
// @Description Read a team.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Team ID."
// @Success 200 {object} model.Team 		"Returns read team"
// @Failure 400 {object} response.BasicResponse "err.team.bind"
// @Failure 400 {object} response.BasicResponse "err.team.read"
// @Resource /teams
// @Router /teams/{id} [get]
func readTeam(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	team, err := teamService.ReadTeam(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.team.read", err)
	}

	return response.SuccessInterface(c, team)
}

// @Title updateTeam
// @Description Update a team.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Team ID."
// @Param   avatar      	form   	string  true	"Team Avatar"
// @Param   firstname		form   	string  true	"Team Firstname"
// @Param   lastname		form   	string  true	"Team Lastname"
// @Param   email	    	form   	string  true	"Team Email"
// @Param   birth      		form   	Time   	true	"Team Birth"
// @Success 200 {object} model.Team 		"Returns read team"
// @Failure 400 {object} response.BasicResponse "err.team.bind"
// @Failure 400 {object} response.BasicResponse "err.team.read"
// @Resource /teams
// @Router /teams/{id} [put]
func updateTeam(c echo.Context) error {
	team := &model.Team{}
	if err := c.Bind(team); err != nil {
		return response.KnownErrJSON(c, "err.team.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	team, err := teamService.UpdateTeam(team)
	if err != nil {
		return response.KnownErrJSON(c, "err.team.update", err)
	}

	team, _ = teamService.ReadTeam(team.ID)
	return response.SuccessInterface(c, team)
}

// @Title deleteTeam
// @Description Delete a team.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Team ID."
// @Success 200 {object} response.BasicResponse "Team is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.team.bind"
// @Failure 400 {object} response.BasicResponse "err.team.delete"
// @Resource /teams
// @Router /teams/{id} [delete]
func deleteTeam(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete team with id
	err := teamService.DeleteTeam(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.team.delete", err)
	}
	return response.SuccessJSON(c, "Team is deleted correctly.")
}

// @Title readTeams
// @Description Read teams with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Team is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.team.read"
// @Resource /teams
// @Router /teams [get]
func readTeams(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read teams with params
	teams, total, err := teamService.ReadTeams(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.team.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, teams})
}
