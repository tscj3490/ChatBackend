package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/siteService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitSites inits site CRUD apis
// @Title Sites
// @Description Sites's router group.
func InitSites(parentRoute *echo.Group) {
	route := parentRoute.Group("/sites")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createSite))
	route.GET("/:id", permission.AuthRequired(readSite))
	route.PUT("/:id", permission.AuthRequired(updateSite))
	route.DELETE("/:id", permission.AuthRequired(deleteSite))

	route.GET("", permission.AuthRequired(readSites))

	siteService.InitService()
}

// @Title createSite
// @Description Create a site.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Site 		"Returns created site"
// @Failure 400 {object} response.BasicResponse "err.site.bind"
// @Failure 400 {object} response.BasicResponse "err.site.create"
// @Resource /sites
// @Router /sites [post]
func createSite(c echo.Context) error {
	site := &model.Site{}
	if err := c.Bind(site); err != nil {
		return response.KnownErrJSON(c, "err.site.bind", err)
	}

	// create site
	site, err := siteService.CreateSite(site)
	if err != nil {
		return response.KnownErrJSON(c, "err.site.create", err)
	}

	return response.SuccessInterface(c, site)
}

// @Title readSite
// @Description Read a site.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Site ID."
// @Success 200 {object} model.Site 		"Returns read site"
// @Failure 400 {object} response.BasicResponse "err.site.bind"
// @Failure 400 {object} response.BasicResponse "err.site.read"
// @Resource /sites
// @Router /sites/{id} [get]
func readSite(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	site, err := siteService.ReadSite(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.site.read", err)
	}

	return response.SuccessInterface(c, site)
}

// @Title updateSite
// @Description Update a site.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Site ID."
// @Param   avatar      	form   	string  true	"Site Avatar"
// @Param   firstname		form   	string  true	"Site Firstname"
// @Param   lastname		form   	string  true	"Site Lastname"
// @Param   email	    	form   	string  true	"Site Email"
// @Param   birth      		form   	Time   	true	"Site Birth"
// @Success 200 {object} model.Site 		"Returns read site"
// @Failure 400 {object} response.BasicResponse "err.site.bind"
// @Failure 400 {object} response.BasicResponse "err.site.read"
// @Resource /sites
// @Router /sites/{id} [put]
func updateSite(c echo.Context) error {
	site := &model.Site{}
	if err := c.Bind(site); err != nil {
		return response.KnownErrJSON(c, "err.site.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	site, err := siteService.UpdateSite(site)
	if err != nil {
		return response.KnownErrJSON(c, "err.site.update", err)
	}

	site, _ = siteService.ReadSite(site.ID)
	return response.SuccessInterface(c, site)
}

// @Title deleteSite
// @Description Delete a site.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Site ID."
// @Success 200 {object} response.BasicResponse "Site is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.site.bind"
// @Failure 400 {object} response.BasicResponse "err.site.delete"
// @Resource /sites
// @Router /sites/{id} [delete]
func deleteSite(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete site with id
	err := siteService.DeleteSite(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.site.delete", err)
	}
	return response.SuccessJSON(c, "Site is deleted correctly.")
}

// @Title readSites
// @Description Read sites with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Site is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.site.read"
// @Resource /sites
// @Router /sites [get]
func readSites(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read sites with params
	sites, total, err := siteService.ReadSites(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.site.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, sites})
}
