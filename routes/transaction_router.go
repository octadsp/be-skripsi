package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/middleware"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(pg.DB)
	productRepository := repositories.RepositoryProduct(pg.DB)
	userRepository := repositories.RepositoryUser(pg.DB)

	h := handlers.HandlerTransaction(cartRepository, productRepository, userRepository)

	// Cart
	e.POST("/cart", middleware.Auth(h.AddToCart))
	e.PUT("/cart", middleware.Auth(h.UpdateCart))
	e.GET("/cart", middleware.Auth(h.GetCart))

	// Delivery

	// Order

	// * Payment
}
