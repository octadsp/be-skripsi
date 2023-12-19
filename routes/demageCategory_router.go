package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func DemageCategoryRoutes(e *echo.Group) {
	demageCategoryRepository := repositories.RepositoryDemageCategory(pg.DB)
	h := handlers.HandlerDemageCategory(demageCategoryRepository)

	e.GET("/demage-categories", h.FindDemageCategories)
	e.GET("/demage-category/:id", h.GetDemageCategory)
	e.POST("/demage-category", h.AddDemageCategory)
	e.PATCH("/demage-category/:id", h.UpdateDemageCategory)
	e.DELETE("/demage-category/:id", h.DeleteDemageCategory)
}
