package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationMasterRoutes(e *echo.Group) {
	reservationMasterRepository := repositories.RepositoryReservation(pg.DB)
	h := handlers.HandlerReservationMaster(reservationMasterRepository)

	e.GET("/reservation-masters", h.FindReservMasters)
	e.GET("/reservation-master/:id", h.GetReservMaster)
	e.POST("/reservation-master", h.AddReservMaster)
	e.PATCH("/reservation-master/:id", h.UpdateReservMaster)
	// e.DELETE("/reservation-master/:id", h.DeleteReservMaster)
}
