package handlers

import (
	reservVehicledto "be-skripsi/dto/reservationVehicles"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservationVehicle struct {
	ReservationVehicleRepository repositories.ReservationVehicleRepository
}

func HandlerReservationVehicle(ReservationVehicleRepository repositories.ReservationVehicleRepository) *handlerReservationVehicle {
	return &handlerReservationVehicle{ReservationVehicleRepository}
}

func (h *handlerReservationVehicle) FindReservVehicles(c echo.Context) error {
	vehicles, err := h.ReservationVehicleRepository.FindReservVehicles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: vehicles})
}

func (h *handlerReservationVehicle) GetReservVehicle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	vehicle, err := h.ReservationVehicleRepository.GetReservVehicle(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: vehicle})
}

func (h *handlerReservationVehicle) AddReservVehicle(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(reservVehicledto.ReservationVehicleReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.ReservationVehicle{
		CarBrand: request.CarBrand,
		CarType:  request.CarType,
		CarYear:  request.CarYear,
		CarColor: request.CarColor,
		Status:   request.Status,
	}

	data, err := h.ReservationVehicleRepository.AddReservVehicle(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddReservVehicle(data)})
}

func (h *handlerReservationVehicle) UpdateReservVehicle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservVehicledto.ReservationVehicleReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationVehicleRepository.GetReservVehicle(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

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

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationVehicleRepository.UpdateReservVehicle(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddReservVehicle(u models.ReservationVehicle) reservVehicledto.ReservationVehicleReq {
	return reservVehicledto.ReservationVehicleReq{
		CarBrand: u.CarBrand,
		CarType:  u.CarType,
		CarYear:  u.CarYear,
		CarColor: u.CarColor,
		Status:   u.Status,
	}
}
