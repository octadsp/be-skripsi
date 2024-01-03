package handlers

import (
	reservItemdto "be-skripsi/dto/reservationItems"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
	// imageFile := c.Get("image").(string)
	price, _ := strconv.Atoi(c.FormValue("price"))

	request := reservItemdto.ReservationItemReqUpdate{
		Image: c.FormValue("image"),
		Item:  c.FormValue("item"),
		Price: int64(price),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	// resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "waysgallery"})

	if err != nil {
		fmt.Println(err.Error())
	}

	reserv := models.ReservationItem{
		Item:   request.Item,
		Image:  request.Image,
		Price:  int64(request.Price),
		Status: "A",
	}

	data, err := h.ReservationItemRepository.AddReservItem(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respAddReservItem(data)})
}

func (h *handlerReservationItem) UpdateReservItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	imageFile := c.Get("image").(string)
	price, _ := strconv.Atoi(c.FormValue("price"))

	request := reservItemdto.ReservationItemReqUpdate{
		Image: imageFile,
		Item:  c.FormValue("item"),
		Price: int64(price),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "waysgallery"})

	if err != nil {
		fmt.Println(err.Error())
	}

	reserv, err := h.ReservationItemRepository.GetReservItem(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Item != "" {
		reserv.Item = request.Item
	}

	if request.Image != "" {
		reserv.Image = resp.SecureURL
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
