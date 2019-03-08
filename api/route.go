// @APIVersion 1.0.0
// @Title gongan API
// @Description gongan API usually works as expected. But sometimes its not true
// @Contact tiandage719@outlook.com
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @BasePath http://127.0.0.1:8000/api/v1
// @SubApi Auth management API [/]
// @SubApi Admins management API [/admins]
// @SubApi Users management API [/users]
// @SubApi Upload management API [/upload]

package api

import (
	"../config"
	"./v1"
	"./v1/auth"

	"github.com/labstack/echo"
)

// RouteAPI contains router groups for API
func RouteAPI(parentRoute *echo.Echo) {
	route := parentRoute.Group(config.APIURL)
	{
		v1.InitUpload(route)

		auth.Init(route)
		v1.InitAdmins(route)
		v1.InitUsers(route)
		v1.InitDevices(route)
		v1.InitLocations(route)
		v1.InitSites(route)
		v1.InitProducts(route)
		v1.InitUserSettings(route)
		v1.InitCustomers(route)
		v1.InitVendors(route)
		v1.InitMakes(route)
		v1.InitMessages(route)
		v1.InitModels(route)
		v1.InitOrders(route)
		v1.InitReviews(route)
		v1.InitServices(route)
		v1.InitTransactions(route)
		v1.InitWorktimes(route)
		v1.InitVendorDevices(route)
		v1.InitVendorPricing(route)
		v1.InitVendorOrders(route)
		v1.InitSpecialtimes(route)
		v1.InitTblregions(route)
		v1.InitOrderServices(route)
		v1.InitJoblogs(route)
		v1.InitPaymentSettings(route)
		v1.InitCoinSettings(route)
		v1.InitSettings(route)

		v1.InitCompanies(route)
		v1.InitTeams(route)
		v1.InitRoles(route)
	}
}
