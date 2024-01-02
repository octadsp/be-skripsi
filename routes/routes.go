package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	ReservationMasterRoutes(e)
	ReservationVehicleRoutes(e)
	ReservationItemRoutes(e)
	ReservationInsuranceRoutes(e)
	CarBrandRoutes(e)
	CarClassRoutes(e)
	CarTypeRoutes(e)
	PriceListRoutes(e)
	DemageCategoryRoutes(e)
	DemageSubCategoryRoutes(e)
	NotificationRoutes(e)
}
