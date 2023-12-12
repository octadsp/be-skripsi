package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationItemRoutes(e *echo.Group) {
	reservationItemRepository := repositories.RepositoryReservationItem(pg.DB)
	h := handlers.HandlerReservationItem(reservationItemRepository)

	e.GET("/reservation-items", h.FindReservItems)
	e.GET("/reservation-item/:id", h.GetReservItem)
	e.POST("/reservation-item", h.AddReservItem)
	e.PATCH("/reservation-item/:id", h.UpdateReservItem)
	// e.DELETE("/reservation-Item/:id", h.DeleteReservItem)
}