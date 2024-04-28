package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(pg.DB)
	brandRepository := repositories.RepositoryBrand(pg.DB)
	categoryRepository := repositories.RepositoryCategory(pg.DB)

	h := handlers.HandlerProduct(productRepository, brandRepository, categoryRepository)

	// Product
	e.POST("/product/new-product", h.NewProduct)
	// e.POST("/product/edit-product", h.EditProduct)

	// Brand
	e.POST("/product/new-brand", h.NewBrand)
	// e.POST("/product/edit-brand", h.EditBrand)

	// Category
	e.POST("/product/new-category", h.NewCategory)
	// e.POST("/product/edit-category", h.EditCategory)
}
