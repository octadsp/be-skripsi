package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func DemageSubCategoryRoutes(e *echo.Group) {
	demageSubCategoryRepository := repositories.RepositoryDemageSubCategory(pg.DB)
	h := handlers.HandlerDemageSubCategory(demageSubCategoryRepository)

	e.GET("/demage-subcategories", h.FindDemageSubCategories)
	e.GET("/demage-subcategory/:id", h.GetDemageSubCategory)
	e.POST("/demage-subcategory", h.AddDemageSubCategory)
	e.PATCH("/demage-subcategory/:id", h.UpdateDemageSubCategory)
	// e.DELETE("/car-brand/:id", h.DeleteCarBrand)
}
