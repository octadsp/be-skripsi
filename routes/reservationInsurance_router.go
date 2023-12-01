package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func ReservationInsuranceRoutes(e *echo.Group) {
	reservationInsuranceRepository := repositories.RepositoryReservationInsurance(pg.DB)
	h := handlers.HandlerReservationInsurance(reservationInsuranceRepository)

	e.GET("/reservation-insurances", h.FindReservInsurances)
	e.GET("/reservation-insurance/:id", h.GetReservInsurance)
	e.POST("/reservation-insurance", h.AddReservInsurance)
	e.PATCH("/reservation-insurance/:id", h.UpdateReservInsurance)
	// e.DELETE("/reservation-Insurance/:id", h.DeleteReservInsurance)
}
