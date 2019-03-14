package auth

import (
	"errors"

	"../../../model"
	"../../../service/authService"
	"../../../service/authService/permission"
	"../../../service/teamService"
	"../../response"

	"../../../config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// UserForm struct.
type UserForm struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

//VerifiedUser struct.
type VerifiedUser struct {
	Token      string `json:"token"`
	IsVerified bool   `json:"is_verified"`
}

// PublicManager
type PublicManager struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}

// Init inits authorization apis
// @Title Auth
// @Description Auth's router group.
func Init(parentRoute *echo.Group) {
	parentRoute.GET("/verifyCode", verifyCode)
	parentRoute.GET("/sendCode", sendCode)
	parentRoute.GET("/addMember", addMember)
	parentRoute.GET("/checkMember", checkMember)
	parentRoute.POST("/add/teamManager", createTeamManager)
	parentRoute.Use(middleware.JWT([]byte(config.AuthTokenKey)))
	// init admin
	// initAdmin(parentRoute)
	// init user
	// initUser(parentRoute)
	// // init vendor
	// initVendor(parentRoute)
	// // init customer
	// initCustomer(parentRoute)
	// init book
	// initBook(parentRoute)

	// parentRoute.GET("/forgotPassword", forgotPassword)
	parentRoute.GET("/inviteMember", permission.AuthRequired(inviteMember))

}

func initAdmin(parentRoute *echo.Group) {
	// admin auth
	parentRoute.POST("/admin/login", loginAdmin)
	parentRoute.POST("/admin/register", registerAdmin)
}

func initUser(parentRoute *echo.Group) {
	// user auth
	parentRoute.POST("/user/login", loginUser)
	parentRoute.POST("/user/register", registerUser)
}

func initVendor(parentRoute *echo.Group) {
	// vendor auth
	parentRoute.POST("/vendor/login", loginVendor)
	parentRoute.POST("/vendor/register", registerVendor)
}

func initCustomer(parentRoute *echo.Group) {
	// customer auth
	parentRoute.POST("/customer/login", loginCustomer)
	parentRoute.POST("/customer/register", registerCustomer)
}

func initBook(parentRoute *echo.Group) {
	// customer auth
	parentRoute.POST("/book/login", loginBook)
	//	parentRoute.POST("/book/register", registerBook)
}

// @Title forgotPassword
// @Description Forgot Password.
// @Accept  json
// @Produce	json
// @Param   username    form   string   true	"Username."
// @Param   role        form   string   true	"Client role."
// @Success 200 {object} string					"Returns result message"
// @Failure 400 {object} response.BasicResponse "err.email.read"
// @Resource /forgotPassword
// @Router /forgotPassword [post]
func forgotPassword(c echo.Context) error {
	username := c.FormValue("username")
	role := c.FormValue("role")
	// handle forgot password
	if ok := authService.ForgotPassword(username, role); !ok {
		return response.KnownErrJSON(c, "err.username.read", errors.New("Username is not existed"))
	}
	return response.SuccessJSON(c, "Server has sent email to you. Please check your email and reset password.")
}

// changePassword
func changePassword(c echo.Context) error {
	chgpass := &model.ChangePass{}
	if err := c.Bind(chgpass); err != nil {
		return response.KnownErrJSON(c, "err.changepassword.bind", err)
	}

	chgpass, err := authService.ChangePassword(chgpass)

	if err != nil {
		return response.KnownErrJSON(c, "err.changepassword.change", err)
	}

	return response.SuccessInterface(c, chgpass)
}

// createTeamManager
func createTeamManager(c echo.Context) error {
	managerInfo := &model.ManagerInfo{}

	if err := c.Bind(managerInfo); err != nil {
		return response.KnownErrJSON(c, "err.managerInfo.bind", err)
	}

	// create team
	_, err := teamService.CreateTeam(managerInfo.Team)
	if err != nil {
		return response.KnownErrJSON(c, "err.team.create", err)
	}

	// role := "manager"
	// // Generate encoded token and send it as response.
	// t, err := permission.GenerateToken(team.ID, role)
	// if err != nil {
	// 	return response.KnownErrJSON(c, "err.auth.token", err)
	// }

	_, err = authService.AddManager(managerInfo)
	if err != nil {
		return response.KnownErrJSON(c, "err.manager.add", err)
	}

	code, err := authService.SendCode(managerInfo.Phone)
	if err != nil {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	return response.SuccessInterface(c, code)
}

// @Title verifyCode
// @Description Verify code.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"User Email."
// @Param   role        form   string   true	"Client role."
// @Param   code        form   string   true	"Veryfy code."
// @Success 200 {object} {object}				"Returns token to reset password"
// @Failure 400 {object} response.BasicResponse "err.email.verify"
// @Failure 400 {object} response.BasicResponse "err.user.read"
// @Resource /verifyCode
// @Router /verifyCode [post]
func verifyCode(c echo.Context) error {
	code := c.FormValue("code")

	role := "manager"
	// check phone number with verify code
	objid, result, err := authService.VerifyCode(code)
	if result != true {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	// Generate encoded token and send it as response.
	t, err := permission.GenerateToken(objid, role)
	if err != nil {
		return response.KnownErrJSON(c, "err.auth.token", err)
	}

	return response.SuccessInterface(c, VerifiedUser{t, result})
}

func sendCode(c echo.Context) error {
	phone := c.FormValue("phone")

	code, err := authService.SendCode(phone)
	if err != nil {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	// return response.SuccessJSON(c, "Server has sent verification code to you. Please confirm verification code.")
	return response.SuccessInterface(c, code)
}

// About loged out user
func checkMember(c echo.Context) error {
	phone := c.FormValue("phone")

	code, err := authService.CheckPhone(phone)
	if err != nil {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	// return response.SuccessJSON(c, "Server has sent verification code to you. Please confirm verification code.")
	return response.SuccessInterface(c, code)
}

func inviteMember(c echo.Context) error {
	phone := c.FormValue("phone")

	code, err := authService.AddOnlyPhone(phone)
	if err != nil {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	// return response.SuccessJSON(c, "Server has sent verification code to you. Please confirm verification code.")
	return response.SuccessInterface(c, code)
}

func addMember(c echo.Context) error {
	code := c.FormValue("code")

	role := "seller"
	// check phone number with verify code
	objid, result, err := authService.VerifyCode(code)
	if result != true {
		return response.KnownErrJSON(c, "err.phone.verify", err)
	}

	// Generate encoded token and send it as response.
	t, err := permission.GenerateToken(objid, role)
	if err != nil {
		return response.KnownErrJSON(c, "err.auth.token", err)
	}

	return response.SuccessInterface(c, VerifiedUser{t, result})

	// var err error
	// code := c.FormValue("code")

	// id := c.Get("user_idx")
	// role := c.Get("user_role")
	// user, err := authService.VerifyRole(id)
	// if err != nil {
	// 	return response.KnownErrJSON(c, "err.phone.verify", err)
	// }

	// return response.SuccessJSON(c, "Server has sent verification code to you. Please confirm verification code.")
	// return response.SuccessInterface(c, user)
}
