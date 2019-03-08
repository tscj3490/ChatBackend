package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/locationService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitLocations inits location CRUD apis
// @Title Locations
// @Description Locations's router group.
func InitLocations(parentRoute *echo.Group) {
	route := parentRoute.Group("/locations")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createLocation))
	route.GET("/:id", permission.AuthRequired(readLocation))
	route.PUT("/:id", permission.AuthRequired(updateLocation))
	route.DELETE("/:id", permission.AuthRequired(deleteLocation))

	route.GET("", permission.AuthRequired(readLocations))

	locationService.InitService()
}

// @Title createLocation
// @Description Create a location.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Location 		"Returns created location"
// @Failure 400 {object} response.BasicResponse "err.location.bind"
// @Failure 400 {object} response.BasicResponse "err.location.create"
// @Resource /locations
// @Router /locations [post]
func createLocation(c echo.Context) error {
	location := &model.Location{}
	if err := c.Bind(location); err != nil {
		return response.KnownErrJSON(c, "err.location.bind", err)
	}

	// create location
	location, err := locationService.CreateLocation(location)
	if err != nil {
		return response.KnownErrJSON(c, "err.location.create", err)
	}

	return response.SuccessInterface(c, location)
}

// @Title readLocation
// @Description Read a location.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Location ID."
// @Success 200 {object} model.Location 		"Returns read location"
// @Failure 400 {object} response.BasicResponse "err.location.bind"
// @Failure 400 {object} response.BasicResponse "err.location.read"
// @Resource /locations
// @Router /locations/{id} [get]
func readLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := locationService.ReadLocation(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.location.read", err)
	}

	return response.SuccessInterface(c, location)
}

// @Title updateLocation
// @Description Update a location.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Location ID."
// @Param   avatar      	form   	string  true	"Location Avatar"
// @Param   firstname		form   	string  true	"Location Firstname"
// @Param   lastname		form   	string  true	"Location Lastname"
// @Param   email	    	form   	string  true	"Location Email"
// @Param   birth      		form   	Time   	true	"Location Birth"
// @Success 200 {object} model.Location 		"Returns read location"
// @Failure 400 {object} response.BasicResponse "err.location.bind"
// @Failure 400 {object} response.BasicResponse "err.location.read"
// @Resource /locations
// @Router /locations/{id} [put]
func updateLocation(c echo.Context) error {
	location := &model.Location{}
	if err := c.Bind(location); err != nil {
		return response.KnownErrJSON(c, "err.location.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	location, err := locationService.UpdateLocation(location)
	if err != nil {
		return response.KnownErrJSON(c, "err.location.update", err)
	}

	location, _ = locationService.ReadLocation(location.ID)
	return response.SuccessInterface(c, location)
}

// @Title deleteLocation
// @Description Delete a location.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Location ID."
// @Success 200 {object} response.BasicResponse "Location is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.location.bind"
// @Failure 400 {object} response.BasicResponse "err.location.delete"
// @Resource /locations
// @Router /locations/{id} [delete]
func deleteLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete location with id
	err := locationService.DeleteLocation(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.location.delete", err)
	}
	return response.SuccessJSON(c, "Location is deleted correctly.")
}

// @Title readLocations
// @Description Read locations with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Location is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.location.read"
// @Resource /locations
// @Router /locations [get]
func readLocations(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read locations with params
	locations, total, err := locationService.ReadLocations(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.location.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, locations})
}
