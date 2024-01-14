package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func RatingRoutes(e *echo.Group) {
	ratingRepository := repositories.RepositoryRating(pg.DB)
	h := handlers.HandlerRating(ratingRepository)

	e.GET("/rating/:userid", h.FindRatingByUser)
	e.POST("/rating", h.AddRating)
}
