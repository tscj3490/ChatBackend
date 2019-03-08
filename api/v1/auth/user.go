package auth

import (
	"errors"

	"../../../config"
	"../../../model"
	"../../../service/authService/permission"
	"../../../service/authService/userService"
	"../../response"

	"github.com/labstack/echo"
)

// @Title loginUser
// @Description Login a user.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"User Email."
// @Param   password	form   string 	true	"User Password."
// @Success 200 {object} UserForm 				"Returns login user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.incorrect"
// @Failure 400 {object} response.BasicResponse "err.user.token"
// @Resource /user/login
// @Router /user/login [post]
func loginUser(c echo.Context) error {
	// user := &model.User{}
	// if err := c.Bind(&user); err != nil {
	// 	return response.KnownErrJSON(c, "err.user.bind", err)
	// }
	// // check user crediential
	// u, err := userService.ReadUserByUsername(user.Name)
	// if err != nil {
	// 	return response.KnownErrJSON(c, "err.user.incorrect", errors.New("Incorrect username"))
	// }
	// err = bcrypt.CompareHashAndPassword([]byte(u.UserPassword), []byte(user.UserPassword))
	// if err != nil {
	// 	return response.KnownErrJSON(c, "err.user.incorrect", errors.New("Incorrect password"))
	// }

	// // generate encoded token and send it as response.
	// t, err := permission.GenerateToken(u.ID, config.RoleUser)
	// if err != nil {
	// 	return response.KnownErrJSON(c, "err.user.token",
	// 		errors.New("Something went wrong. Please check token creating"))
	// }

	// // retreive by public user
	// publicUser := &model.PublicUser{User: u}
	// return response.SuccessInterface(c, UserForm{t, publicUser})
	return nil
}

// @Title registerUser
// @Description Register a user.
// @Accept  json
// @Produce	json
// @Param   firstname	form   string   true	"User Firstname."
// @Param   lastname   	form   string   true	"User Lastname."
// @Param   email       form   string   true	"User Email."
// @Param   password	form   string 	true	"User Password."
// @Success 200 {object} UserForm				"Returns registered user"
// @Failure 400 {object} response.BasicResponse "err.user.bind"
// @Failure 400 {object} response.BasicResponse "err.user.exist"
// @Failure 400 {object} response.BasicResponse "err.user.create"
// @Failure 400 {object} response.BasicResponse "err.user.token"
// @Resource /user/register
// @Router /user/register [post]
func registerUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return response.KnownErrJSON(c, "err.user.bind", err)
	}

	// check existed email
	if _, err := userService.ReadUserByUsername(user.Name); err == nil {
		return response.KnownErrJSON(c, "err.user.exist",
			errors.New("Same username is existed. Please input other username"))
	}

	// create user with registered info
	user, err := userService.CreateUser(user)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.create", err)
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(user.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.user.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public user
	publicUser := &model.PublicUser{User: user}
	return response.SuccessInterface(c, UserForm{t, publicUser})
}
