package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/middleware"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationItemRoutes(e *echo.Group) {
	reservationItemRepository := repositories.RepositoryReservationItem(pg.DB)
	h := handlers.HandlerReservationItem(reservationItemRepository)

	e.GET("/reservation-items", h.FindReservItems)
	e.GET("/reservation-item/:id", h.GetReservItem)
	e.GET("/reservation-item-byreservation/:reservId", h.GetReservItemByReservId)
	e.POST("/reservation-item", middleware.UploadImage(h.AddReservItem))
	e.PATCH("/reservation-item/:id", middleware.UploadImage(h.UpdateReservItem))
	// e.DELETE("/reservation-Item/:id", h.DeleteReservItem)

	e.PATCH("/reservation-item-status/:id", h.UpdateStatus)
}
