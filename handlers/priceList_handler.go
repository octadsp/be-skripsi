package handlers

import (
	priceListsdto "be-skripsi/dto/priceLists"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerPriceList struct {
	PriceListRepository repositories.PriceListRepository
}

func HandlerPriceList(PriceListRepository repositories.PriceListRepository) *handlerPriceList {
	return &handlerPriceList{PriceListRepository}
}

func (h *handlerPriceList) FindPriceLists(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Jumlah item per halaman
	itemsPerPage := 10

	// Hitung offset berdasarkan halaman
	offset := (page - 1) * itemsPerPage

	price, err := h.PriceListRepository.FindPriceLists(offset, itemsPerPage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: price})
}

func (h *handlerPriceList) FindAllPriceLists(c echo.Context) error {
	price, err := h.PriceListRepository.FindAllPriceLists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: price})
}

func (h *handlerPriceList) GetPriceList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	price, err := h.PriceListRepository.GetPriceList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: price})
}

func (h *handlerPriceList) AddPriceList(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId, ok := userLogin.(jwt.MapClaims)["id"].(float64)
	// if !ok {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid user ID"})
	// }

	// fmt.Println("user_id :", int(userId))

	request := new(priceListsdto.PriceListReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	price := models.PriceList{
		DemageSubCategoryID: request.DemageSubCategoryID,
		CarClassID:          request.CarClassID,
		Price:               request.Price,
		Status:              "A",
	}

	data, err := h.PriceListRepository.AddPriceList(price)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddprice(data)})
}

func (h *handlerPriceList) UpdatePriceList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(priceListsdto.PriceListReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	price, err := h.PriceListRepository.GetPriceList(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.DemageSubCategoryID != 0 {
		price.DemageSubCategoryID = request.DemageSubCategoryID
	}

	if request.CarClassID != 0 {
		price.CarClassID = request.CarClassID
	}

	if request.Price != 0 {
		price.Price = request.Price
	}

	if request.Status != "" {
		price.Status = request.Status
	}

	price.UpdatedAt = time.Now()

	data, err := h.PriceListRepository.UpdatePriceList(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerPriceList) DeletePriceList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	price, err := h.PriceListRepository.GetPriceList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.PriceListRepository.DeletePriceList(price, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func respAddprice(u models.PriceList) priceListsdto.PriceListResp {
	return priceListsdto.PriceListResp{
		DemageSubCategoryID: u.DemageSubCategoryID,
		CarClassID:          u.CarClassID,
		Price:               u.Price,
		Status:              u.Status,
	}
}

// func respAllPrice(u []models.PriceList) []priceListsdto.PriceListResp {
// 	var response []priceListsdto.PriceListResp

// 	for _, item := range u {
// 		resp := priceListsdto.PriceListResp{
// 			DemageSubCategoryID: item.DemageSubCategoryID,
// 			CarClassID:          item.CarClassID,
// 			CarClass:            models.CarClass(item.CarClass),
// 			Price:               item.Price,
// 			Status:              item.Status,
// 		}

// 		response = append(response, resp)
// 	}

// 	return response
// }
