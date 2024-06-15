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
	deliveryFareRepository := repositories.RepositoryDeliveryFare(pg.DB)
	productRepository := repositories.RepositoryProduct(pg.DB)
	userAddressRepository := repositories.RepositoryUserAddress(pg.DB)
	userRepository := repositories.RepositoryUser(pg.DB)

	h := handlers.HandlerTransaction(cartRepository, deliveryFareRepository, productRepository, userAddressRepository, userRepository)

	// Cart
	e.POST("/cart", middleware.Auth(h.AddToCart))
	e.PUT("/cart", middleware.Auth(h.UpdateCart))
	e.GET("/cart", middleware.Auth(h.GetCart))

	// Delivery Fare
	e.GET("/delivery/fare", middleware.Auth(h.GetDeliveryFare))
	// * Admin Operations only
	e.POST("/delivery/fare", middleware.Auth(h.AddDeliveryFare))
	e.GET("/delivery/fares", middleware.Auth(h.GetDeliveryFares))
	e.PUT("/delivery/fare/:id", middleware.Auth(h.UpdateDeliveryFare))

	// Order

	// * Payment
}
