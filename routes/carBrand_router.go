package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func CarBrandRoutes(e *echo.Group) {
	carBrandRepository := repositories.RepositoryCarBrand(pg.DB)
	h := handlers.HandlerCarBrand(carBrandRepository)

	e.GET("/car-brands", h.FindCarBrands)
	e.GET("/car-brand/:id", h.GetCarBrand)
	e.POST("/car-brand", h.AddCarBrand)
	e.PATCH("/car-brand/:id", h.UpdateCarBrand)
	// e.DELETE("/car-brand/:id", h.DeleteCarBrand)
}
