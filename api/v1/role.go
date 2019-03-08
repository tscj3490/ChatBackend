package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/roleService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitRoles inits role CRUD apis
// @Title Roles
// @Description Roles's router group.
func InitRoles(parentRoute *echo.Group) {
	route := parentRoute.Group("/roles")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createRole))
	route.GET("/:id", permission.AuthRequired(readRole))
	route.PUT("/:id", permission.AuthRequired(updateRole))
	route.DELETE("/:id", permission.AuthRequired(deleteRole))

	route.GET("", permission.AuthRequired(readRoles))

	roleService.InitService()
}

// @Title createRole
// @Description Create a role.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Role 		"Returns created role"
// @Failure 400 {object} response.BasicResponse "err.role.bind"
// @Failure 400 {object} response.BasicResponse "err.role.create"
// @Resource /roles
// @Router /roles [post]
func createRole(c echo.Context) error {
	role := &model.Role{}
	if err := c.Bind(role); err != nil {
		return response.KnownErrJSON(c, "err.role.bind", err)
	}

	// create role
	role, err := roleService.CreateRole(role)
	if err != nil {
		return response.KnownErrJSON(c, "err.role.create", err)
	}

	return response.SuccessInterface(c, role)
}

// @Title readRole
// @Description Read a role.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Role ID."
// @Success 200 {object} model.Role 		"Returns read role"
// @Failure 400 {object} response.BasicResponse "err.role.bind"
// @Failure 400 {object} response.BasicResponse "err.role.read"
// @Resource /roles
// @Router /roles/{id} [get]
func readRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := roleService.ReadRole(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.role.read", err)
	}

	return response.SuccessInterface(c, role)
}

// @Title updateRole
// @Description Update a role.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Role ID."
// @Param   avatar      	form   	string  true	"Role Avatar"
// @Param   firstname		form   	string  true	"Role Firstname"
// @Param   lastname		form   	string  true	"Role Lastname"
// @Param   email	    	form   	string  true	"Role Email"
// @Param   birth      		form   	Time   	true	"Role Birth"
// @Success 200 {object} model.Role 		"Returns read role"
// @Failure 400 {object} response.BasicResponse "err.role.bind"
// @Failure 400 {object} response.BasicResponse "err.role.read"
// @Resource /roles
// @Router /roles/{id} [put]
func updateRole(c echo.Context) error {
	role := &model.Role{}
	if err := c.Bind(role); err != nil {
		return response.KnownErrJSON(c, "err.role.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	role, err := roleService.UpdateRole(role)
	if err != nil {
		return response.KnownErrJSON(c, "err.role.update", err)
	}

	role, _ = roleService.ReadRole(role.ID)
	return response.SuccessInterface(c, role)
}

// @Title deleteRole
// @Description Delete a role.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Role ID."
// @Success 200 {object} response.BasicResponse "Role is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.role.bind"
// @Failure 400 {object} response.BasicResponse "err.role.delete"
// @Resource /roles
// @Router /roles/{id} [delete]
func deleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete role with id
	err := roleService.DeleteRole(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.role.delete", err)
	}
	return response.SuccessJSON(c, "Role is deleted correctly.")
}

// @Title readRoles
// @Description Read roles with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Role is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.role.read"
// @Resource /roles
// @Router /roles [get]
func readRoles(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read roles with params
	roles, total, err := roleService.ReadRoles(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.role.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, roles})
}
