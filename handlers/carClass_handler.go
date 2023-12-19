package handlers

import (
	carClassdto "be-skripsi/dto/carClass"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCarClass struct {
	CarClassRepository repositories.CarClassRepository
}

func HandlerCarClass(CarClassRepository repositories.CarClassRepository) *handlerCarClass {
	return &handlerCarClass{CarClassRepository}
}

func (h *handlerCarClass) FindCarClass(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Jumlah item per halaman
	itemsPerPage := 10

	// Hitung offset berdasarkan halaman
	offset := (page - 1) * itemsPerPage

	class, err := h.CarClassRepository.FindCarClass(offset, itemsPerPage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: class})
}

func (h *handlerCarClass) FindAllCarClass(c echo.Context) error {
	class, err := h.CarClassRepository.FindAllCarClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: class})
}

func (h *handlerCarClass) GetCarClass(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	class, err := h.CarClassRepository.GetCarClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: class})
}

func (h *handlerCarClass) AddCarClass(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(carClassdto.CarClassReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	class := models.CarClass{
		CarBrandID: request.CarBrandID,
		CarTypeID:  request.CarTypeID,
		Golongan:   request.Golongan,
		Status:     "A",
	}

	data, err := h.CarClassRepository.AddCarClass(class)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddClass(data)})
}

func (h *handlerCarClass) UpdateCarClass(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(carClassdto.CarClassReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	class, err := h.CarClassRepository.GetCarClass(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.CarBrandID != 0 {
		class.CarBrandID = request.CarBrandID
	}

	if request.CarTypeID != 0 {
		class.CarTypeID = request.CarTypeID
	}

	if request.Status != "" {
		class.Status = request.Status
	}

	if request.Golongan != "" {
		class.Golongan = request.Golongan
	}

	class.UpdatedAt = time.Now()

	data, err := h.CarClassRepository.UpdateCarClass(class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerCarClass) DeleteCarClass(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.CarClassRepository.GetCarClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CarClassRepository.DeleteCarClass(brand, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddClass(u models.CarClass) carClassdto.CarClassResp {
	return carClassdto.CarClassResp{
		CarBrandID: u.CarBrandID,
		CarTypeID:  u.CarTypeID,
		Golongan:   u.Golongan,
		Status:     u.Status,
	}
}
