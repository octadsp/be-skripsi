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

	product := &models.Product{
		ID:              uuid.New().String()[:8],
		ProductName:     request.ProductName,
		BrandID:         request.BrandID,
		CategoryID:      request.CategoryID,
		Price:           request.Price,
		InstallationFee: request.InstallationFee,
		OpeningStock:    request.OpeningStock,
	}

	productData, err := h.ProductRepository.CreateProduct(*product)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: productData})
}

func (h *handlerProduct) GetProducts(c echo.Context) error {
	productsData, err := h.ProductRepository.GetProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: productsData})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id := c.Param("id")

	productData, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: productData})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	request := new(productDto.UpdateProductRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	_, err = h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	product := &models.Product{
		ProductName:     request.ProductName,
		BrandID:         request.BrandID,
		CategoryID:      request.CategoryID,
		Price:           request.Price,
		InstallationFee: request.InstallationFee,
	}

	_, err = h.ProductRepository.UpdateProduct(id, *product)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Product updated successfully!"})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	_, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	_, err = h.ProductRepository.DeleteProduct(id)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Product deleted successfully!"})
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

func (h *handlerProduct) GetBrands(c echo.Context) error {
	brandsData, err := h.BrandRepository.GetBrands()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: brandsData})
}

func (h *handlerProduct) GetBrand(c echo.Context) error {
	id := c.Param("id")

	brandData, err := h.BrandRepository.GetBrand(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: brandData})
}

func (h *handlerProduct) UpdateBrand(c echo.Context) error {
	id := c.Param("id")
	request := new(productDto.UpdateBrandRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	_, err = h.BrandRepository.GetBrand(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	brand := &models.Brand{
		BrandName: request.BrandName,
	}

	_, err = h.BrandRepository.UpdateBrand(id, *brand)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Brand updated successfully!"})
}

func (h *handlerProduct) DeleteBrand(c echo.Context) error {
	id := c.Param("id")

	_, err := h.BrandRepository.GetBrand(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	_, err = h.BrandRepository.DeleteBrand(id)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Brand deleted successfully!"})
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

func (h *handlerProduct) GetCategories(c echo.Context) error {
	categoriesData, err := h.CategoryRepository.GetCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: categoriesData})
}

func (h *handlerProduct) GetCategory(c echo.Context) error {
	id := c.Param("id")

	categoryData, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: categoryData})
}

func (h *handlerProduct) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	request := new(productDto.UpdateCategoryRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	_, err = h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	category := &models.Category{
		CategoryName: request.CategoryName,
	}

	_, err = h.CategoryRepository.UpdateCategory(id, *category)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Category updated successfully!"})
}

func (h *handlerProduct) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	_, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	_, err = h.CategoryRepository.DeleteCategory(id)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Category deleted successfully!"})
}
