package handlers

import (
	reservInsurancedto "be-skripsi/dto/reservationInsurances"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservationInsurance struct {
	ReservationInsuranceRepository repositories.ReservationInsuranceRepository
}

func HandlerReservationInsurance(ReservationInsuranceRepository repositories.ReservationInsuranceRepository) *handlerReservationInsurance {
	return &handlerReservationInsurance{ReservationInsuranceRepository}
}

func (h *handlerReservationInsurance) FindReservInsurances(c echo.Context) error {
	insurances, err := h.ReservationInsuranceRepository.FindReservInsurances()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: insurances})
}

func (h *handlerReservationInsurance) GetReservInsurance(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	insurance, err := h.ReservationInsuranceRepository.GetReservInsurance(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: insurance})
}

func (h *handlerReservationInsurance) AddReservInsurance(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(reservInsurancedto.ReservationInsuranceReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.ReservationInsurance{
		EventDate:      request.EventDate,
		Place:          request.Place,
		Time:           request.Time,
		DrivingSpeed:   request.DrivingSpeed,
		DriverName:     request.DriverName,
		DriverRelation: request.DriverRelation,
		DriverJob:      request.DriverJob,
		DriverAge:      request.DriverAge,
		Status:         request.Status,
	}

	data, err := h.ReservationInsuranceRepository.AddReservInsurance(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddReservInsurance(data)})
}

func (h *handlerReservationInsurance) UpdateReservInsurance(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservInsurancedto.ReservationInsuranceReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationInsuranceRepository.GetReservInsurance(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
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

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationInsuranceRepository.UpdateReservInsurance(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddReservInsurance(u models.ReservationInsurance) reservInsurancedto.ReservationInsuranceReq {
	return reservInsurancedto.ReservationInsuranceReq{
		EventDate:      u.EventDate,
		Place:          u.Place,
		Time:           u.Time,
		DrivingSpeed:   u.DrivingSpeed,
		DriverName:     u.DriverName,
		DriverRelation: u.DriverRelation,
		DriverJob:      u.DriverJob,
		DriverAge:      u.DriverAge,
		DriverLicense:  u.DriverLicense,
		Status:         u.Status,
	}
}
