package handlers

import (
	carBranddto "be-skripsi/dto/carBrands"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCarBrand struct {
	CarBrandRepository repositories.CarBrandRepository
}

func HandlerCarBrand(CarBrandRepository repositories.CarBrandRepository) *handlerCarBrand {
	return &handlerCarBrand{CarBrandRepository}
}

func (h *handlerCarBrand) FindCarBrands(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Jumlah item per halaman
	itemsPerPage := 10

	// Hitung offset berdasarkan halaman
	offset := (page - 1) * itemsPerPage

	brands, err := h.CarBrandRepository.FindCarBrands(offset, itemsPerPage)
	if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: brands})
}

func (h *handlerCarBrand) FindAllBrands(c echo.Context) error {
	brands, err := h.CarBrandRepository.FindAllBrands()
	if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: brands})
}

func (h *handlerCarBrand) GetCarBrand(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	brand, err := h.CarBrandRepository.GetCarBrand(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: brand})
}

func (h *handlerCarBrand) AddCarBrand(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(carBranddto.CarBrandReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	brand := models.CarBrand{
		Name:   request.Name,
		Tipe:   request.Tipe,
		Status: "A",
	}

	data, err := h.CarBrandRepository.AddCarBrand(brand)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddBrand(data)})
}

func (h *handlerCarBrand) UpdateCarBrand(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(carBranddto.CarBrandReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	brand, err := h.CarBrandRepository.GetCarBrand(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		brand.Name = request.Name
	}

	if request.Tipe != "" {
		brand.Tipe = request.Tipe
	}

	if request.Status != "" {
		brand.Status = request.Status
	}

	brand.UpdatedAt = time.Now()

	data, err := h.CarBrandRepository.UpdateCarBrand(brand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerCarBrand) DeleteCarBrand(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.CarBrandRepository.GetCarBrand(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CarBrandRepository.DeleteCarBrand(brand, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddBrand(u models.CarBrand) carBranddto.CarBrandReq {
	return carBranddto.CarBrandReq{
		Name:   u.Name,
		Tipe:   u.Tipe,
		Status: u.Status,
	}
}
