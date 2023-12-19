package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func CarTypeRoutes(e *echo.Group) {
	carTypeRepository := repositories.RepositoryCarType(pg.DB)
	h := handlers.HandlerCarType(carTypeRepository)

	e.GET("/car-types", h.FindCarTypes)
	e.GET("/car-all-types", h.FindAllCarTypes)
	e.GET("/car-type/:id", h.GetCarType)
	e.POST("/car-type", h.AddCarType)
	e.PATCH("/car-type/:id", h.UpdateCarType)
	e.DELETE("/car-type/:id", h.DeleteCarType)
}
