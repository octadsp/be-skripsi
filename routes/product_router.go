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
	e.POST("/product/", h.NewProduct)
	// e.GET("/products", h.GetProducts)
	// e.GET("/product/:id", h.GetProduct)
	// e.PUT("/product/:id", h.UpdateProduct)
	// e.PATCH("/product/:id", middleware.UploadImage(h.UpdateProductImage))
	// e.DELETE("/product/:id", h.DeleteProduct)

	// Brand
	e.POST("/brand", h.NewBrand)
	e.GET("/brands", h.GetBrands)
	e.GET("/brand/:id", h.GetBrand)
	e.PUT("/brand/:id", h.UpdateBrand)
	e.DELETE("/brand/:id", h.DeleteBrand)

	// Category
	e.POST("/category", h.NewCategory)
	e.GET("/categories", h.GetCategories)
	e.GET("/category/:id", h.GetCategory)
	e.PUT("/category/:id", h.UpdateCategory)
	e.DELETE("/category/:id", h.DeleteCategory)
}
