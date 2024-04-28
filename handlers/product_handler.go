package handlers

import (
	productDto "be-skripsi/dto/product"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	errors "be-skripsi/pkg/error"
	repository "be-skripsi/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository  repository.ProductRepository
	BrandRepository    repository.BrandRepository
	CategoryRepository repository.CategoryRepository
}

func HandlerProduct(ProductRepository repository.ProductRepository, BrandRepository repository.BrandRepository, CategoryRepository repository.CategoryRepository) *handlerProduct {
	return &handlerProduct{ProductRepository, BrandRepository, CategoryRepository}
}

/*
 * 	Product
 */
func (h *handlerProduct) NewProduct(c echo.Context) error {
	request := new(productDto.NewProductRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: "Ehehe"})
}

/*
 * 	Brand
 */
func (h *handlerProduct) NewBrand(c echo.Context) error {
	request := new(productDto.NewBrandRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	brand := &models.Brand{
		ID:        uuid.New().String()[:8],
		BrandName: request.BrandName,
	}

	brandData, err := h.BrandRepository.CreateBrand(*brand)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: brandData})
}

/*
 *	Category
 */
func (h *handlerProduct) NewCategory(c echo.Context) error {
	request := new(productDto.NewCategoryRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	category := &models.Category{
		ID:           uuid.New().String()[:8],
		CategoryName: request.CategoryName,
	}

	categoryData, err := h.CategoryRepository.CreateCategory(*category)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: categoryData})
}
