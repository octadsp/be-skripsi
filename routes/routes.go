package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	ReservationMasterRoutes(e)
	CarBrandRoutes(e)
	CarClassRoutes(e)
	CarTypeRoutes(e)
	PriceListRoutes(e)
	DemageCategoryRoutes(e)
	DemageSubCategoryRoutes(e)
}
