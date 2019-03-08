package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"../../../config"
	"../../../model"
	"../../../service/authService/adminService"
	"../../../service/authService/permission"
	"../../response"

	"github.com/labstack/echo"
)

// @Title loginAdmin
// @Description Login a admin.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"Admin Email."
// @Param   password	form   string 	true	"Admin Password."
// @Success 200 {object} UserForm 				"Returns login admin"
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.incorrect"
// @Failure 400 {object} response.BasicResponse "err.admin.token"
// @Resource /admin/login
// @Router /admin/login [post]
func loginAdmin(c echo.Context) error {
	admin := &model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return response.KnownErrJSON(c, "err.admin.bind", err)
	}

	// check admin crediential
	a, err := adminService.ReadAdminByUsername(admin.Username)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.incorrect", errors.New("Incorrect email or password"))
	}
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(admin.Password))
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.incorrect", errors.New("Incorrect password"))
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(a.ID, config.RoleAdmin)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public admin
	publicAdmin := &model.PublicAdmin{Admin: a}
	return response.SuccessInterface(c, UserForm{t, publicAdmin})
}

// @Title registerAdmin
// @Description Register a admin.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"Admin Email."
// @Param   password	form   string 	true	"Admin Password."
// @Success 200 {object} UserForm				"Returns registered admin"
// @Failure 400 {object} response.BasicResponse "err.admin.bind"
// @Failure 400 {object} response.BasicResponse "err.admin.exist"
// @Failure 400 {object} response.BasicResponse "err.admin.create"
// @Failure 400 {object} response.BasicResponse "err.admin.token"
// @Resource /admin/register
// @Router /admin/register [post]
func registerAdmin(c echo.Context) error {
	admin := &model.Admin{}
	if err := c.Bind(admin); err != nil {
		return response.KnownErrJSON(c, "err.admin.bind", err)
	}

	// check existed username
	if _, err := adminService.ReadAdminByUsername(admin.Username); err == nil {
		return response.KnownErrJSON(c, "err.admin.exist",
			errors.New("Same username is existed. Please input other username"))
	}

	// create admin with registered info
	admin, err := adminService.CreateAdmin(admin)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.create", err)
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(admin.ID, config.RoleAdmin)
	if err != nil {
		return response.KnownErrJSON(c, "err.admin.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public admin
	publicAdmin := &model.PublicAdmin{Admin: admin}
	return response.SuccessInterface(c, UserForm{t, publicAdmin})
}
