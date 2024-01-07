package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/middleware"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(pg.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user-image/:id", middleware.UploadImage(h.UpdateUser))
	e.PATCH("/user-info/:id", h.UpdateInfoUser)
	e.DELETE("/user/:id", h.DeleteUser)
}
