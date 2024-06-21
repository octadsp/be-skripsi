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

	h := handlers.HandlerTransaction(
		cartRepository,
		deliveryFareRepository,
		productRepository,
		userAddressRepository,
		userRepository,
	)

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

	// TODO Order
	e.POST("/order", middleware.Auth(h.NewOrder))       // Customer create new order
	e.GET("/orders", middleware.Auth(h.GetOrders))      // Customer get all orders
	e.GET("/order/:id", middleware.Auth(h.GetOrder))    // Customer get specific order
	e.PUT("/order/:id", middleware.Auth(h.UpdateOrder)) // Customer update specific order (status)
	// * Admin Operations only
	// e.GET("/orders/admin", middleware.Auth(h.AdminGetOrders))      // Admin get all orders
	// e.PUT("/order/:id/admin", middleware.Auth(h.AdminUpdateOrder)) // Admin update specific order (status)

	// // * Payment
	// e.POST("/order/:id/payment", middleware.Auth(h.SubmitNewPayment))
	// e.GET("/order/payments", middleware.Auth(h.GetAllPayment))
	// e.GET("/order/:id/payment", middleware.Auth(h.GetPaymentByTransactionID))
	// e.GET("/order/payment/:id", middleware.Auth(h.GetPaymentByPaymentID))
	// e.PUT("/order/payment/:id", middleware.Auth(h.UpdatePaymentByPaymentID))
	// // * Admin Operations only
	// e.GET("/order/payment/admin", middleware.Auth(h.AdminGetAllPayment))
}
