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
	orderRepository := repositories.RepositoryOrder(pg.DB)
	productRepository := repositories.RepositoryProduct(pg.DB)
	productStockHistoryRepository := repositories.RepositoryProductStockHistory(pg.DB)
	userAddressRepository := repositories.RepositoryUserAddress(pg.DB)
	userRepository := repositories.RepositoryUser(pg.DB)

	h := handlers.HandlerTransaction(
		cartRepository,
		deliveryFareRepository,
		orderRepository,
		productRepository,
		productStockHistoryRepository,
		userAddressRepository,
		userRepository,
	)

	// Cart
	e.POST("/cart", middleware.Auth(h.AddToCart))
	e.PUT("/cart", middleware.Auth(h.UpdateCart))
	e.GET("/cart", middleware.Auth(h.GetCart))
	e.DELETE("/cart/:id", middleware.Auth(h.DeleteCart))

	// Delivery Fare
	e.GET("/delivery/fare", middleware.Auth(h.GetDeliveryFare))
	// * Admin Operations only
	e.POST("/delivery/fare", middleware.Auth(h.AddDeliveryFare))
	e.GET("/delivery/fares", middleware.Auth(h.GetDeliveryFares))
	e.PUT("/delivery/fare/:id", middleware.Auth(h.UpdateDeliveryFare))

	// Order
	e.POST("/order", middleware.Auth(h.NewOrder))             // OK Customer create new order
	e.GET("/orders", middleware.Auth(h.GetOrders))            // OK Customer get all orders
	e.GET("/order/:id", middleware.Auth(h.GetOrder))          // OK Customer get specific order
	e.PUT("/order/:id", middleware.Auth(h.UpdateOrder))       // OK Customer update specific order status
	e.GET("/orders/admin", middleware.Auth(h.AdminGetOrders)) // OK Admin get all orders (can use filter)

	// * Payment
	e.POST("/order/:id/payment", middleware.Auth(middleware.UploadImage(h.SubmitNewPayment))) // OK Customer submit new payment
	e.GET("/order/payments", middleware.Auth(h.GetAllPayment))                                // OK Admin get all payment
	e.GET("/order/:id/payment", middleware.Auth(h.GetPaymentByOrderID))                       // OK Customer get payment from specific order
	e.GET("/order/payment/:id", middleware.Auth(h.GetPaymentByPaymentID))                     // OK Admin get payment detail
	e.PUT("/order/payment/:id", middleware.Auth(h.UpdatePaymentByPaymentID))                  // Admin update payment
	// * Admin Operations only
	// e.GET("/order/payment/admin", middleware.Auth(h.AdminGetAllPayment))
}
