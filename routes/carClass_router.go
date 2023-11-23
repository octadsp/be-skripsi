package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func CarClassRoutes(e *echo.Group) {
	carClassRepository := repositories.RepositoryCarClass(pg.DB)
	h := handlers.HandlerCarClass(carClassRepository)

	e.GET("/car-class", h.FindCarClass)
	e.GET("/car-class/:id", h.GetCarClass)
	e.POST("/car-class", h.AddCarClass)
	e.PATCH("/car-class/:id", h.UpdateCarClass)
	// e.DELETE("/car-brand/:id", h.DeleteCarBrand)
}
