package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/middleware"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(pg.DB)
	userDetailRepository := repositories.RepositoryUserDetail(pg.DB)

	h := handlers.HandlerAuth(userRepository, userDetailRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
	e.POST("/update-password", middleware.Auth(h.UpdatePassword))
	e.POST("/make-admin/:id", middleware.Auth(h.MakeAdmin))
}
