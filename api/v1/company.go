package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/companyService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitCompanies inits company CRUD apis
// @Title Companies
// @Description Companies's router group.
func InitCompanies(parentRoute *echo.Group) {
	route := parentRoute.Group("/companies")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createCompany))
	route.GET("/:id", permission.AuthRequired(readCompany))
	route.PUT("/:id", permission.AuthRequired(updateCompany))
	route.DELETE("/:id", permission.AuthRequired(deleteCompany))

	route.GET("", permission.AuthRequired(readCompanies))

	companyService.InitService()
}

// @Title createCompany
// @Description Create a company.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.Company 		"Returns created company"
// @Failure 400 {object} response.BasicResponse "err.company.bind"
// @Failure 400 {object} response.BasicResponse "err.company.create"
// @Resource /companies
// @Router /companies [post]
func createCompany(c echo.Context) error {
	company := &model.Company{}
	if err := c.Bind(company); err != nil {
		return response.KnownErrJSON(c, "err.company.bind", err)
	}

	// create company
	company, err := companyService.CreateCompany(company)
	if err != nil {
		return response.KnownErrJSON(c, "err.company.create", err)
	}

	return response.SuccessInterface(c, company)
}

// @Title readCompany
// @Description Read a company.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Company ID."
// @Success 200 {object} model.Company 		"Returns read company"
// @Failure 400 {object} response.BasicResponse "err.company.bind"
// @Failure 400 {object} response.BasicResponse "err.company.read"
// @Resource /companies
// @Router /companies/{id} [get]
func readCompany(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := companyService.ReadCompany(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.company.read", err)
	}

	return response.SuccessInterface(c, company)
}

// @Title updateCompany
// @Description Update a company.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Company ID."
// @Param   avatar      	form   	string  true	"Company Avatar"
// @Param   firstname		form   	string  true	"Company Firstname"
// @Param   lastname		form   	string  true	"Company Lastname"
// @Param   email	    	form   	string  true	"Company Email"
// @Param   birth      		form   	Time   	true	"Company Birth"
// @Success 200 {object} model.Company 		"Returns read company"
// @Failure 400 {object} response.BasicResponse "err.company.bind"
// @Failure 400 {object} response.BasicResponse "err.company.read"
// @Resource /companies
// @Router /companies/{id} [put]
func updateCompany(c echo.Context) error {
	company := &model.Company{}
	if err := c.Bind(company); err != nil {
		return response.KnownErrJSON(c, "err.company.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	company, err := companyService.UpdateCompany(company)
	if err != nil {
		return response.KnownErrJSON(c, "err.company.update", err)
	}

	company, _ = companyService.ReadCompany(company.ID)
	return response.SuccessInterface(c, company)
}

// @Title deleteCompany
// @Description Delete a company.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Company ID."
// @Success 200 {object} response.BasicResponse "Company is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.company.bind"
// @Failure 400 {object} response.BasicResponse "err.company.delete"
// @Resource /companies
// @Router /companies/{id} [delete]
func deleteCompany(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete company with id
	err := companyService.DeleteCompany(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.company.delete", err)
	}
	return response.SuccessJSON(c, "Company is deleted correctly.")
}

// @Title readCompanies
// @Description Read companies with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Company is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.company.read"
// @Resource /companies
// @Router /companies [get]
func readCompanies(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read companies with params
	companies, total, err := companyService.ReadCompanies(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.company.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, companies})
}
