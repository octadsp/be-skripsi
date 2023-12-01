package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationVehicleRoutes(e *echo.Group) {
	reservationVehicleRepository := repositories.RepositoryReservationVehicle(pg.DB)
	h := handlers.HandlerReservationVehicle(reservationVehicleRepository)

	e.GET("/reservation-vehicles", h.FindReservVehicles)
	e.GET("/reservation-vehicle/:id", h.GetReservVehicle)
	e.POST("/reservation-vehicle", h.AddReservVehicle)
	e.PATCH("/reservation-vehicle/:id", h.UpdateReservVehicle)
	// e.DELETE("/reservation-Vehicle/:id", h.DeleteReservVehicle)
}
