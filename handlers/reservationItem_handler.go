package handlers

import (
	reservItemdto "be-skripsi/dto/reservationItems"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservationItem struct {
	ReservationItemRepository repositories.ReservationItemRepository
}

func HandlerReservationItem(ReservationItemRepository repositories.ReservationItemRepository) *handlerReservationItem {
	return &handlerReservationItem{ReservationItemRepository}
}

func (h *handlerReservationItem) FindReservItems(c echo.Context) error {
	items, err := h.ReservationItemRepository.FindReservItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: items})
}

func (h *handlerReservationItem) GetReservItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	item, err := h.ReservationItemRepository.GetReservItem(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: item})
}

func (h *handlerReservationItem) AddReservItem(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(reservItemdto.ReservationItemReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.ReservationItem{
		Item:   request.Item,
		Price:  int64(request.Price),
		Status: request.Status,
	}

	data, err := h.ReservationItemRepository.AddReservItem(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddReservItem(data)})
}

func (h *handlerReservationItem) UpdateReservItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservItemdto.ReservationItemReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationItemRepository.GetReservItem(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Item != "" {
		reserv.Item = request.Item
	}

	if request.Price != 0 {
		reserv.Price = request.Price
	}

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationItemRepository.UpdateReservItem(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddReservItem(u models.ReservationItem) reservItemdto.ReservationItemReq {
	return reservItemdto.ReservationItemReq{
		Item:   u.Item,
		Price:  u.Price,
		Status: u.Status,
	}
}
