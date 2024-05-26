package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/middleware"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(pg.DB)
	productImageRepository := repositories.RepositoryProductImage(pg.DB)
	productStockHistoryRepository := repositories.RepositoryProductStockHistory(pg.DB)
	brandRepository := repositories.RepositoryBrand(pg.DB)
	categoryRepository := repositories.RepositoryCategory(pg.DB)

	h := handlers.HandlerProduct(productRepository, productImageRepository, productStockHistoryRepository, brandRepository, categoryRepository)

	// Product
	e.POST("/product", h.NewProduct)
	e.GET("/products", h.GetProducts)
	e.GET("/product/:id", h.GetProduct)
	e.PUT("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)

	// * Product Image
	e.PATCH("/product-image/:product_id", middleware.UploadImage(h.UpdateProductImage))
	e.DELETE("/product-image/:product_image_id", h.DeleteProductImage)

	// * Product Stock History
	e.POST("/product-stock/:product_id", h.UpdateProductStock)

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
