package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(pg.DB)
	userDetailRepository := repositories.RepositoryUserDetail((pg.DB))
	userAddressRepository := repositories.RepositoryUserAddress((pg.DB))
	h := handlers.HandlerUser(userRepository, userDetailRepository, userAddressRepository)

	// User Detail

	// User Address

	// * Master Province
	e.GET("/provinces", h.GetProvinces)
	e.GET("/province/:id", h.GetProvinceByID)

	// * Master Regency
	e.GET("/regencies", h.GetRegencies)
	e.GET("/regencies/:province_id", h.GetRegenciesByProvinceID)
	e.GET("/regency/:id", h.GetRegencyByID)

	// * Master District
	e.GET("/districts", h.GetDistricts)
	e.GET("/districts/:regency_id", h.GetDistrictsByRegencyID)
	e.GET("/district/:id", h.GetDistrictByID)
}
