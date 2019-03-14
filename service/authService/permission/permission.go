package permission

import (
	"fmt"
	"time"

	"../../../config"
	"../../../util/timeHelper"
	// "../adminService"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GenerateToken returns token after generate with user
func GenerateToken(id uint, role string) (string, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["idx"] = id
	claims["exp"] = timeHelper.FewDaysLater(config.AuthTokenExpirationDay)
	claims["role"] = role

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AuthTokenKey))
	return t, err
}

// InfoFromToken returns idx from token
func InfoFromToken(c echo.Context) {
	fmt.Printf("InfoFromToken=%+v\n", c.Get("user"))
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	c.Set("user_idx", claims["idx"])
	c.Set("user_role", claims["role"].(string))

	// var id uint
	// var role string
	// // retrieve object id of client
	// if claims["idx"] != nil {
	// 	id = uint(claims["idx"].(float64))
	// }
	// // retrieve role: admin, user
	// if claims["role"] != nil {
	// 	role = claims["role"].(string)
	// }
	//return id, role
}

func expiredFromToken(c echo.Context) bool {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	exp, err := time.Parse(time.RFC3339, claims["exp"].(string))
	if err != nil {
		return false
	}

	return timeHelper.IsExpired(exp)
}

// AuthRequired run function when user logged in.
func AuthRequired(f func(c echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		InfoFromToken(c)
		// // checking expire date
		// if expiredFromToken(c) {
		// 	log.Error("Token is expired.")
		// 	return response.KnownErrorJSON(c, http.StatusUnauthorized, "error.token.expire", errors.New("Token is expired"))
		// }

		// // checking client validation
		// {
		// 	id, role := InfoFromToken(c)
		// 	var err error
		// 	switch role {
		// 	case config.RoleAdmin:
		// 		_, err = adminService.ReadAdmin(id)
		// 	case config.RoleUser:
		// 		_, err = userService.ReadUser(id)
		// 	}
		// 	if err != nil {
		// 		log.Error("Auth failed.")
		// 		return response.KnownErrorJSON(c, http.StatusUnauthorized, "error.auth.fail", errors.New("Auth failed"))
		// 	}
		// }
		// checking client validation
		// {
		// 	id, _ := InfoFromToken(c)
		// 	var err error
		// 	_, err = userService.ReadUser(id)

		// 	if err != nil {
		// 		log.Error("Auth failed.")
		// 		return response.KnownErrorJSON(c, http.StatusUnauthorized, "error.auth.fail", errors.New("Auth failed"))
		// 	}
		// }
		f(c)

		return nil
	}
}
