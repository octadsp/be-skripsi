package handlers

import (
	dto "be-skripsi/dto/results"
	usersdto "be-skripsi/dto/users"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: users})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: user})
}

func (h *handlerUser) UpdateUser(c echo.Context) error {
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)
	id, _ := strconv.Atoi(c.Param("id"))

	imageFile := c.Get("image").(string)

	request := usersdto.UserUpdateRequest{
		FullName: c.FormValue("fullname"),
		LastName: c.FormValue("lastname"),
		Email:    c.FormValue("email"),
		Phone:    c.FormValue("phone"),
		Address:  c.FormValue("address"),
		Avatar:   imageFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "skripsi"})

	if err != nil {
		fmt.Println(err.Error())
	}

	user, err := h.UserRepository.GetUser(id)
	// user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}
	if request.LastName != "" {
		user.LastName = request.LastName
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Address != "" {
		user.Address = request.Address
	}
	if request.Avatar != "" {
		user.Avatar = resp.SecureURL
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerUser) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: convertResponse(data)})

}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		LastName: u.LastName,
		Email:    u.Email,
		Status:   u.Status,
		Roles:    u.Roles,
	}
}
