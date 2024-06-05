package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	ProductRoutes(e)
	UserRoutes(e)
}
