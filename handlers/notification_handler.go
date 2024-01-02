package handlers

import (
	notificationsdto "be-skripsi/dto/notification"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerNotification struct {
	NotificationRepository repositories.NotificationRepository
}

func HandlerNotification(NotificationRepository repositories.NotificationRepository) *handlerNotification {
	return &handlerNotification{NotificationRepository}
}

func (h *handlerNotification) GetNotificationsByUserID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	}

	// Konversi userID ke uint
	userIDUint := uint(userID)

	// Panggil fungsi repository untuk mendapatkan notifikasi berdasarkan ID pengguna
	notifications, err := h.NotificationRepository.GetNotificationsByUserID(userIDUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: notifications})
}

func (h *handlerNotification) CreateNotification(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := new(notificationsdto.NotifReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	notif := &models.Notification{
		UserID:  userId,
		Message: request.Message,
		IsRead:  false,
		CreatedAt: time.Now(),
	}

	data, err := h.PostRepository.CreateNotification(notif)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerNotification) UpdateNotificationStatus(c echo.Context) error {
	notifID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid notification ID"})
	}

	isReadParam := c.QueryParam("isRead")
	isRead, err := strconv.ParseBool(isReadParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid isRead parameter"})
	}

	notification, err := h.NotificationRepository.GetNotif(uint(notifID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Update read status
	notification.IsRead = isRead

	err = h.NotificationRepository.UpdateNotificationStatus(uint(notifID), notification.IsRead)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: notification})
}

