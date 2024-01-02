package handlers

import (
	reservationsdto "be-skripsi/dto/reservations"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservation struct {
	ReservationRepository repositories.ReservationRepository
}

func HandlerReservation(ReservationRepository repositories.ReservationRepository) *handlerReservation {
	return &handlerReservation{ReservationRepository}
}

func (h *handlerReservation) FindReservations(c echo.Context) error {
	s, err := h.ReservationRepository.FindReservations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: s})
}

func (h *handlerReservation) GetReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	reserv, err := h.ReservationRepository.GetReservation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reserv})
}

func (h *handlerReservation) AddReservation(c echo.Context) error {
	request := new(reservationsdto.ReservationReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.Reservation{
		KodeOrder:  request.KodeOrder,
		Status:     request.Status,
		OrderMasuk: request.OrderMasuk,
		UserID:     request.UserID,

		CarBrand: request.CarBrand,
		CarType:  request.CarType,
		CarYear:  request.CarYear,
		CarColor: request.CarColor,

		IsInsurance: request.IsInsurance,

		InsuranceName:  request.InsuranceName,
		EventDate:      request.EventDate,
		Place:          request.Place,
		Time:           request.Time,
		DrivingSpeed:   request.DrivingSpeed,
		DriverName:     request.DriverName,
		DriverRelation: request.DriverRelation,
		DriverJob:      request.DriverJob,
		DriverAge:      request.DriverAge,
		DriverLicense:  request.DriverLicense,

		ReservationItemID: request.ReservationItemID,

		CreatedAt: time.Now(),
	}

	data, err := h.ReservationRepository.AddReservation(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerReservation) UpdateReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservationsdto.ReservationReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationRepository.GetReservation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Gunakan time.Time.Nil untuk menandakan bahwa field tersebut tidak diubah
	var orderProses, orderSelesai time.Time

	// Mengecek apakah request.OrderProses dan request.OrderSelesai berisi nilai yang tidak nol
	if !request.OrderProses.IsZero() {
		orderProses = request.OrderProses
	}

	if !request.OrderSelesai.IsZero() {
		orderSelesai = request.OrderSelesai
	}

	// Update nilai-nilai yang tidak nol
	reserv.OrderProses = orderProses
	reserv.OrderSelesai = orderSelesai

	if request.CarBrand != "" {
		reserv.CarBrand = request.CarBrand
	}

	if request.CarType != "" {
		reserv.CarType = request.CarType
	}

	if request.CarYear != "" {
		reserv.CarYear = request.CarYear
	}

	if request.CarColor != "" {
		reserv.CarColor = request.CarColor
	}

	if request.IsInsurance != 0 {
		reserv.IsInsurance = request.IsInsurance
	}

	if request.InsuranceName != "" {
		reserv.InsuranceName = request.InsuranceName
	}

	if request.EventDate != "" {
		reserv.EventDate = request.EventDate
	}
	if request.Place != "" {
		reserv.Place = request.Place
	}
	if request.Time != "" {
		reserv.Time = request.Time
	}
	if request.DrivingSpeed != "" {
		reserv.DrivingSpeed = request.DrivingSpeed
	}
	if request.DriverName != "" {
		reserv.DriverName = request.DriverName
	}
	if request.DriverRelation != "" {
		reserv.DriverRelation = request.DriverRelation
	}
	if request.DriverJob != "" {
		reserv.DriverJob = request.DriverJob
	}
	if request.DriverAge != "" {
		reserv.DriverAge = request.DriverAge
	}
	if request.DriverLicense != "" {
		reserv.DriverLicense = request.DriverLicense
	}
	if request.ReservationItemID != 0 {
		reserv.ReservationItemID = request.ReservationItemID
	}

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationRepository.UpdateReservation(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerReservation) DeleteReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.ReservationRepository.GetReservation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ReservationRepository.DeleteReservation(brand)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
