package handlers

import (
	// authDto "be-skripsi/dto/auth"
	dto "be-skripsi/dto/results"

	// bcrypt "be-skripsi/pkg/bcrypt"
	"be-skripsi/repositories"

	// "fmt"
	"net/http"
	// "os"
	// "strconv"

	// "context"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository        repositories.UserRepository
	UserDetailRepository  repositories.UserDetailRepository
	UserAddressRepository repositories.UserAddressRepository
}

func HandlerUser(UserRepository repositories.UserRepository, UserDetailRepository repositories.UserDetailRepository, UserAddressRepository repositories.UserAddressRepository) *handlerUser {
	return &handlerUser{UserRepository, UserDetailRepository, UserAddressRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	// users, _ := h.UserRepository.FindUsers()
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "test"})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))
	// // userLogin := c.Get("userLogin")
	// // userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// user, err := h.UserRepository.GetUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "user"})
}

func (h *handlerUser) UpdateUser(c echo.Context) error {
	// // userLogin := c.Get("userLogin")
	// // userId := userLogin.(jwt.MapClaims)["id"].(float64)
	// id, _ := strconv.Atoi(c.Param("id"))

	// imageFile := c.Get("image").(string)

	// request := usersdto.UserUpdateRequest{
	// 	Avatar: imageFile,
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// // Add your Cloudinary credentials ...
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// // Upload file to Cloudinary ...
	// resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "skripsi"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// user, err := h.UserRepository.GetUser(id)
	// // user, err := h.UserRepository.GetUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// if request.Avatar != "" {
	// 	user.Avatar = resp.SecureURL
	// }

	// data, err := h.UserRepository.UpdateUser(user)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "data"})
}

func (h *handlerUser) UpdateInfoUser(c echo.Context) error {
	// // userLogin := c.Get("userLogin")
	// // userId := userLogin.(jwt.MapClaims)["id"].(float64)
	// id, _ := strconv.Atoi(c.Param("id"))

	// request := new(usersdto.UserUpdateRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// user, err := h.UserRepository.GetUser(id)
	// // user, err := h.UserRepository.GetUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// if request.FullName != "" {
	// 	user.FullName = request.FullName
	// }
	// if request.LastName != "" {
	// 	user.LastName = request.LastName
	// }
	// if request.Email != "" {
	// 	user.Email = request.Email
	// }
	// if request.Institute != "" {
	// 	user.Institute = request.Institute
	// }
	// if request.Phone != "" {
	// 	user.Phone = request.Phone
	// }
	// if request.Address != "" {
	// 	user.Address = request.Address
	// }

	// data, err := h.UserRepository.UpdateInfoUser(user)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "data"})
}

func (h *handlerUser) DeleteUser(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))

	// user, err := h.UserRepository.GetUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	// }

	// data, err := h.UserRepository.DeleteUser(user)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "success delete user"})

}

// func IdentityResponse(u models.User, ud models.UserDetail) userDto.UserIdentityResponse {
// 	return userDto.UserIdentityResponse{
// 		ID:       u.ID,
// 		FullName: ud.FullName,
// 		Email:    u.Email,
// 		Phone:    ud.PhoneNumber,
// 	}
// }

// func LoginResponse(u models.User, ud models.UserDetail) userDto.UserLoginResponse {
// 	return userDto.UserLoginResponse{
// 		ID:       ud.ID,
// 		FullName: ud.FullName,
// 		Email:    u.Email,
// 		Phone:    ud.PhoneNumber,
// 	}
// }

// func DefaultUserReponse(u models.User, ud models.UserDetail) userDto.UserDefaultResponse {
// 	return userDto.UserDefaultResponse{
// 		ID:       u.ID,
// 		FullName: ud.FullName,
// 		Email:    u.Email,
// 		Phone:    ud.PhoneNumber,
// 	}
// }
