package handlers

import (
	demagesubcategoriesdto "be-skripsi/dto/demageSubCategories"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerDemageSubCategory struct {
	DemageSubCategory repositories.DemageSubCategoryRepository
}

func HandlerDemageSubCategory(DemageSubCategory repositories.DemageSubCategoryRepository) *handlerDemageSubCategory {
	return &handlerDemageSubCategory{DemageSubCategory}
}

func (h *handlerDemageSubCategory) FindDemageSubCategories(c echo.Context) error {
	demages, err := h.DemageSubCategory.FindDemageSubCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: demages})
}

func (h *handlerDemageSubCategory) GetDemageSubCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	demage, err := h.DemageSubCategory.GetDemageSubCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: demage})
}

func (h *handlerDemageSubCategory) AddDemageSubCategory(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(demagesubcategoriesdto.DemageSubCategoryReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	demage := models.DemageSubCategory{
		DemageCategoryID: uint32(request.DemageCategoryID),
		Name:             request.Name,
		Status:           "A",
	}

	data, err := h.DemageSubCategory.AddDemageSubCategory(demage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddDemageSubCat(data)})
}

func (h *handlerDemageSubCategory) UpdateDemageSubCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(demagesubcategoriesdto.DemageSubCategoryReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	demage, err := h.DemageSubCategory.GetDemageSubCategory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		demage.Name = request.Name
	}

	if request.DemageCategoryID != 0 {
		demage.DemageCategoryID = request.DemageCategoryID
	}

	if request.Status != "" {
		demage.Status = request.Status
	}

	demage.UpdatedAt = time.Now()

	data, err := h.DemageSubCategory.UpdateDemageSubCategory(demage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddDemageSubCat(u models.DemageSubCategory) demagesubcategoriesdto.DemageSubCategoryReq {
	return demagesubcategoriesdto.DemageSubCategoryReq{
		DemageCategoryID: u.DemageCategoryID,
		Name:             u.Name,
		Status:           u.Status,
	}
}
