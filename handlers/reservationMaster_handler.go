package handlers

import (
	reservMasterdto "be-skripsi/dto/reservationMasters"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservationMaster struct {
	ReservationMasterRepository repositories.ReservationMasterRepository
}

func HandlerReservationMaster(ReservationMasterRepository repositories.ReservationMasterRepository) *handlerReservationMaster {
	return &handlerReservationMaster{ReservationMasterRepository}
}

func (h *handlerReservationMaster) FindReservMasters(c echo.Context) error {
	vehicles, err := h.ReservationMasterRepository.FindReservMasters()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: vehicles})
}

func (h *handlerReservationMaster) GetReservMaster(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	vehicle, err := h.ReservationMasterRepository.GetReservMaster(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: vehicle})
}

func (h *handlerReservationMaster) AddReservMaster(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(reservMasterdto.ReservationMasterReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.ReservationMaster{
		KodeOrder: request.KodeOrder,
		Status:    request.Status,
		// OrderMasuk: request.OrderMasuk,
		// UserID:                 uint32(userId),
		UserID:                 uint32(request.UserID),
		ReservationVehicleID:   uint32(request.ReservationVehicleID),
		ReservationInsuranceID: uint32(request.ReservationInsuranceID),
		ReservationItemID:      uint32(request.ReservationItemID),
	}

	data, err := h.ReservationMasterRepository.AddReservMaster(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddReserv(data)})
}

func (h *handlerReservationMaster) UpdateReservMaster(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservMasterdto.ReservationMasterReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationMasterRepository.GetReservMaster(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Simpan status sebelum pembaruan
	oldStatus := reserv.Status

	// if request.KodeOrder != "" {
	// 	reserv.KodeOrder = request.KodeOrder
	// }

	if request.Status != "" {
		// Jika status berubah, update tanggal sesuai kondisi
		if oldStatus != request.Status {
			switch request.Status {
			case "M":
				reserv.OrderMasuk = time.Now()
			case "P":
				reserv.OrderProses = time.Now()
			case "S":
				reserv.OrderSelesai = time.Now()
			}
		}
		reserv.Status = request.Status
	}

	// if request.UserID != 0 {
	// 	reserv.UserID = request.UserID
	// }

	// if request.ReservationVehicleID != 0 {
	// 	reserv.ReservationVehicleID = request.ReservationVehicleID
	// }

	// if request.ReservationInsuranceID != 0 {
	// 	reserv.ReservationInsuranceID = request.ReservationInsuranceID
	// }

	// if request.ReservationItemID != 0 {
	// 	reserv.ReservationItemID = request.ReservationItemID
	// }

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationMasterRepository.UpdateReservMaster(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddReserv(u models.ReservationMaster) reservMasterdto.ReservationMasterReq {
	return reservMasterdto.ReservationMasterReq{
		KodeOrder:              u.KodeOrder,
		Status:                 u.Status,
		OrderMasuk:             u.OrderMasuk,
		UserID:                 u.UserID,
		ReservationInsuranceID: u.ReservationInsuranceID,
		ReservationVehicleID:   u.ReservationVehicleID,
		ReservationItemID:      u.ReservationItemID,
	}
}
