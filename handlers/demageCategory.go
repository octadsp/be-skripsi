package handlers

import (
	demagecategoriesdto "be-skripsi/dto/demageCategories"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerDemageCategory struct {
	DemageCategory repositories.DemageCategoryRepository
}

func HandlerDemageCategory(DemageCategory repositories.DemageCategoryRepository) *handlerDemageCategory {
	return &handlerDemageCategory{DemageCategory}
}

func (h *handlerDemageCategory) FindDemageCategories(c echo.Context) error {
	demages, err := h.DemageCategory.FindDemageCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: demages})
}

func (h *handlerDemageCategory) GetDemageCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	demage, err := h.DemageCategory.GetDemageCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: demage})
}

func (h *handlerDemageCategory) AddDemageCategory(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(demagecategoriesdto.DemageCategoryReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	demage := models.DemageCategory{
		Name:   request.Name,
		Kode:   request.Kode,
		Status: "A",
	}

	data, err := h.DemageCategory.AddDemageCategory(demage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddDemageCat(data)})
}

func (h *handlerDemageCategory) UpdateDemageCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(demagecategoriesdto.DemageCategoryReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	demage, err := h.DemageCategory.GetDemageCategory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		demage.Name = request.Name
	}

	if request.Kode != "" {
		demage.Kode = request.Kode
	}

	if request.Status != "" {
		demage.Status = request.Status
	}

	demage.UpdatedAt = time.Now()

	data, err := h.DemageCategory.UpdateDemageCategory(demage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerDemageCategory) DeleteDemageCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.DemageCategory.GetDemageCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.DemageCategory.DeleteDemageCategory(brand, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddDemageCat(u models.DemageCategory) demagecategoriesdto.DemageCategoryReq {
	return demagecategoriesdto.DemageCategoryReq{
		Name:   u.Name,
		Kode:   u.Kode,
		Status: u.Status,
	}
}
