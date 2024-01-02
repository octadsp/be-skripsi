package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func NotificationRoutes(e *echo.Group) {
	notificationRepository := repositories.RepositoryNotification(pg.DB)
	h := handlers.HandlerNotification(notificationRepository)

	e.GET("/notifications/:userID", h.GetNotificationsByUserIDHandler)
	e.GET("/notification/:id", h.GetNotif)
	e.POST("/notification", h.CreateNotification)
	e.PATCH("/notification/:notifID", h.UpdateNotificationStatus)

}
