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
	userDetailRepository := repositories.RepositoryUserDetail(pg.DB)
	userAddressRepository := repositories.RepositoryUserAddress(pg.DB)
	userMessageRepository := repositories.RepositoryUserMessage(pg.DB)
	h := handlers.HandlerUser(userRepository, userDetailRepository, userAddressRepository, userMessageRepository)

	// User Detail
	e.PUT("/user-detail", middleware.Auth(h.UpdateUserDetail))

	// User Address
	e.POST("/user-address", middleware.Auth(h.NewUserAddress))
	e.GET("/user-addresses", middleware.Auth(h.GetUserAddresses))
	e.GET("/user-address/:id", middleware.Auth(h.GetUserAddressByID))
	e.PUT("/user-address/:id", middleware.Auth(h.UpdateUserAddressByID))
	e.DELETE("/user-address/:id", middleware.Auth(h.DeleteUserAddressByID))

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

	// Chatbox
	e.GET("/chats", middleware.Auth(h.GetChats)) //OK
	e.GET("/chat/unread", middleware.Auth(h.GetUnreadChats))
	e.GET("/chat/:other_user_id", middleware.Auth(h.GetChatLogs)) //OK
	e.GET("/chat/read/:other_user_id", middleware.Auth(h.ReadChat))
	e.POST("/chat/:other_user_id", middleware.Auth(h.SendMessage)) //OK
}
