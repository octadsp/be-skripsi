package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationRoutes(e *echo.Group) {
	reservationRepository := repositories.RepositoryReservation(pg.DB)
	h := handlers.HandlerReservation(reservationRepository)

	e.GET("/reservations", h.FindReservations)
	e.GET("/reservation/:id", h.GetReservation)
	e.POST("/reservation", h.AddReservation)
	e.PATCH("/reservation/:id", h.UpdateReservation)
	e.DELETE("/reservation/:id", h.DeleteReservation)
	e.PATCH("/reservation-status/:id", h.UpdateStatusReserv)

	e.GET("/reservation-substatus/:substatus", h.GetReservSubStatus)
}