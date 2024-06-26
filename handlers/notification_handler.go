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

	// Konversi userID ke uint
	userIDUint := uint(id)

	// Panggil fungsi repository untuk mendapatkan notifikasi berdasarkan ID pengguna
	notifications, err := h.NotificationRepository.GetNotificationsByUserID(userIDUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: notifications})
}

func (h *handlerNotification) GetNotif(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.NotificationRepository.GetNotif(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: post})
}

func (h *handlerNotification) CreateNotification(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(uint32)

	request := new(notificationsdto.NotifReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	notif := models.Notification{
		UserID:    request.UserID,
		AuthorBy:  request.AuthorBy,
		Title:     request.Title,
		Message:   request.Message,
		IsRead:    false,
		CreatedAt: time.Now(),
	}

	data, err := h.NotificationRepository.CreateNotification(notif)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerNotification) UpdateNotificationStatus(c echo.Context) error {
	notifID, err := strconv.Atoi(c.Param("notifID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid notification ID"})
	}

	request := new(notificationsdto.NotifReqStatus)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Get the existing notification
	notification, err := h.NotificationRepository.GetNotif(notifID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Update read status
	if request.IsRead {
		notification.IsRead = request.IsRead
	}

	notification.UpdatedAt = time.Now()

	// Update the notification status and get the updated notification
	updatedNotification, err := h.NotificationRepository.UpdateNotificationStatus(notification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: updatedNotification})
}
