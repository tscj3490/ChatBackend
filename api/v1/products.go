package v1

import (
	"fmt"
	"strconv"

	"../../service/authService/adminService"
	"../response"

	"../../config"
	"../../model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// InitAdmins inits admin CRUD apis
// @Title Admins
// @Description Admins's router group.
func InitProducts(parentRoute *echo.Group) {
	route := parentRoute.Group("/products")
	//route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	//route.POST("", createProduct)
	route.GET("/:id", readProduct)
	//route.PUT("/:id", updateProduct)
	//route.DELETE("/:id", deleteProduct)

	route.GET("", readProducts)

	adminService.InitService()
}

// @Title createAdmin
// @Description Create a admin.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   email       	form   	string  true	"Admin Email."
// @Param   password		form   	string 	true	"Admin Password."
// @Success 200 {object} model.PublicAdmin 		"Returns created admin"
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.create"
// @Resource /admins
// @Router /admins [post]
// func createAdmin(c echo.Context) error {
// 	admin := &model.Admin{}
// 	if err := c.Bind(admin); err != nil {
// 		return response.KnownErrJSON(c, "err.admin.bind", err)
// 	}

// 	// create admin
// 	admin, err := adminService.CreateAdmin(admin)
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.admin.create", err)
// 	}

// 	publicAdmin := &model.PublicAdmin{Admin: admin}
// 	return response.SuccessInterface(c, publicAdmin)
// }

// @Title readAdmin
// @Description Read a admin.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Admin ID."
// @Success 200 {object} model.PublicAdmin 		"Returns read admin"
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.read"
// @Resource /admins
// @Router /admins/{id} [get]

// @Title updateAdmin
// @Description Update a admin.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"Admin ID."
// @Param   avatar      	form   	string  true	"Admin Avatar"
// @Param   firstname		form   	string  true	"Admin Firstname"
// @Param   lastname		form   	string  true	"Admin Lastname"
// @Param   email	    	form   	string  true	"Admin Email"
// @Param   birth      		form   	Time   	true	"Admin Birth"
// @Success 200 {object} model.PublicAdmin 		"Returns read admin"
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.read"
// @Resource /admins
// @Router /admins/{id} [put]
// func updateAdmin(c echo.Context) error {
// 	admin := &model.Admin{}
// 	if err := c.Bind(admin); err != nil {
// 		return response.KnownErrJSON(c, "err.admin.bind", err)
// 	}

// 	// id, _ := strconv.Atoi(c.Param("id"))
// 	admin, err := adminService.UpdateAdmin(admin)
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.admin.update", err)
// 	}

// 	admin, _ = adminService.ReadAdmin(admin.ID)
// 	publicAdmin := &model.PublicAdmin{Admin: admin}
// 	return response.SuccessInterface(c, publicAdmin)
// }

// @Title deleteAdmin
// @Description Delete a admin.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"Admin ID."
// @Success 200 {object} response.BasicResponse "Admin is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.delete"
// @Resource /admins
// @Router /admins/{id} [delete]
// func deleteAdmin(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	// delete admin with id
// 	err := adminService.DeleteAdmin(uint(id))
// 	if err != nil {
// 		return response.KnownErrJSON(c, "err.admin.delete", err)
// 	}
// 	return response.SuccessJSON(c, "Admin is deleted correctly.")
// }

// @Title readAdmins
// @Description Read admins with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"Admin is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.admin.read"
// @Resource /admins
// @Router /admins [get]
func readProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := gorm.Open("mysql", config.MysqlDSL())

	var product model.Product
	db.First(&product, "id = ?", id)
	// read admins with params
	fmt.Printf("%+v", product)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.read", err)
	}

	return response.SuccessInterface(c, product)
}

func readProducts(c echo.Context) error {
	fmt.Println("Here")

	db, err := gorm.Open("mysql", config.MysqlDSL())

	var products []model.Product
	db.Find(&products)
	// read admins with params
	fmt.Printf("%+v", products)
	fmt.Printf("%#v", products)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.read", err)
	}

	return response.SuccessInterface(c, &ListForm{0, products})
}
