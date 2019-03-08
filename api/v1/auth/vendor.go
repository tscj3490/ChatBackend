package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../../config"
	"../../../model"
	"../../../service/authService/permission"
	"../../../service/authService/vendorOrdersService"
	"../../../service/authService/vendorsService"
	"../../response"

	"github.com/cydev/zero"
	"github.com/labstack/echo"
)

// @Title loginVendor
// @Description Login a vendor.
// @Accept  json
// @Produce	json
// @Param   email       form   string   true	"Vendor Email."
// @Param   password	form   string 	true	"Vendor Password."
// @Success 200 {object} VendorForm 				"Returns login Vendor"
// @Failure 400 {object} response.BasicResponse "err.vendor.bind"
// @Failure 400 {object} response.BasicResponse "err.vendor.incorrect"
// @Failure 400 {object} response.BasicResponse "err.vendor.token"
// @Resource /vendor/login
// @Router /vendor/login [post]
func loginVendor(c echo.Context) error {
	vendor := &model.Vendor{}
	if err := c.Bind(&vendor); err != nil {
		return response.KnownErrJSON(c, "err.vendor.bind", err)
	}
	//	fmt.Printf("%+v", vendor)
	// check user crediential
	u, err := vendorsService.ReadVendorByUsername(vendor.UserName)
	//	fmt.Println("user:", u, "err:", err)
	//if err != nil {
	//	return response.KnownErrJSON(c, "err.vendor.incorrect", errors.New("Incorrect username"))
	//}
	fmt.Println(u)
	if err != nil {
		return response.KnownErrJSON(c, "err.customer.incorrect", errors.New("Incorrect username"))
	}
	fmt.Println(vendor.UserName)
	fmt.Println(vendor.Email)
	fmt.Println(vendor.Password)
	fmt.Println(u.Password)

	if u.Password != vendor.Password {
		return response.KnownErrJSON(c, "err.customer.incorrect", errors.New("Incorrect password"))
	}
	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(u.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public vendor
	publicVendor := &model.PublicVendor{Vendor: u}
	return response.SuccessInterface(c, UserForm{t, publicVendor})
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
func registerVendor(c echo.Context) error {
	vendor := &model.Vendor{}
	if err := c.Bind(vendor); err != nil {
		return response.KnownErrJSON(c, "err.vendor.bind", err)
	}

	// check existed email
	if _, err := vendorsService.ReadVendorByUsername(vendor.UserName); err == nil {
		return response.KnownErrJSON(c, "err.vendor.exist",
			errors.New("Same username is existed. Please input other username"))
	}

	// payment_type register

	vo, err := vendorOrdersService.SetVendorOrder(vendor)
	if err != nil {
		fmt.Println(vo)
		return response.KnownErrJSON(c, "err.vendororder.create", err)
	}

	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + vendor.PostCode
	url = url + "&sensor=false&key=AIzaSyBl74CajeHYL2nszJXq-rQVhYrN6-9mG7A"

	req, _ := http.NewRequest("GET", url, nil)
	res_api, _ := http.DefaultClient.Do(req)

	defer res_api.Body.Close()

	body, _ := ioutil.ReadAll(res_api.Body)

	regioninfo := model.RegionInfo{}

	if err := json.Unmarshal(body, &regioninfo); err == nil {
		if !zero.IsZero(regioninfo) {
			lat := regioninfo.Results[0].Geometry.Location.Lat
			lng := regioninfo.Results[0].Geometry.Location.Lng

			vendor.Lat = lat
			vendor.Lng = lng

			fmt.Println(vendor.Lat, vendor.Lng)
		}
	}
	vendor.CoinAmount = 50.0
	// create user with registered info
	vendor, err = vendorsService.CreateVendorWithEmail(vendor)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.create", err)
	}

	// generate encoded token and send it as response.
	t, err := permission.GenerateToken(vendor.ID, config.RoleUser)
	if err != nil {
		return response.KnownErrJSON(c, "err.vendor.token",
			errors.New("Something went wrong. Please check token creating"))
	}

	// retreive by public user
	publicVendor := &model.PublicVendor{Vendor: vendor}
	return response.SuccessInterface(c, UserForm{t, publicVendor})
}
