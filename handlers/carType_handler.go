package handlers

import (
	carTypedto "be-skripsi/dto/carTypes"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCarType struct {
	CarTypeRepository repositories.CarTypeRepository
}

func HandlerCarType(CarTypeRepository repositories.CarTypeRepository) *handlerCarType {
	return &handlerCarType{CarTypeRepository}
}

func (h *handlerCarType) FindCarTypes(c echo.Context) error {
	types, err := h.CarTypeRepository.FindCarTypes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: types})
}

func (h *handlerCarType) GetCarType(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	types, err := h.CarTypeRepository.GetCarType(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: types})
}

func (h *handlerCarType) AddCarType(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(carTypedto.CarTypeReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	types := models.CarType{
		Name:   request.Name,
		Tipe:   request.Tipe,
		Status: "A",
	}

	data, err := h.CarTypeRepository.AddCarType(types)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddType(data)})
}

func (h *handlerCarType) UpdateCarType(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(carTypedto.CarTypeReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	types, err := h.CarTypeRepository.GetCarType(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		types.Name = request.Name
	}

	if request.Tipe != "" {
		types.Tipe = request.Tipe
	}

	if request.Status != "" {
		types.Status = request.Status
	}

	types.UpdatedAt = time.Now()

	data, err := h.CarTypeRepository.UpdateCarType(types)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerCarType) DeleteCarType(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.CarTypeRepository.GetCarType(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CarTypeRepository.DeleteCarType(brand, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddType(u models.CarType) carTypedto.CarTypeReq {
	return carTypedto.CarTypeReq{
		Name:   u.Name,
		Tipe:   u.Tipe,
		Status: u.Status,
	}
}
