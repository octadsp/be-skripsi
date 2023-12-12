package handlers

import (
	authdto "be-skripsi/dto/auth"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/pkg/bcrypt"
	jwtToken "be-skripsi/pkg/jwt"
	repository "be-skripsi/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

// Register Handler
func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		FullName: request.FullName,
		LastName: request.LastName,
		Address:  request.Address,
		Email:    request.Email,
		Password: password,
		Avatar:   "https://res.cloudinary.com/dpxazv6a6/image/upload/v1685689207/skripsi/defaultAvatar_gsslte.png",
		Phone:    request.Phone,
		Status:   "A",
		Roles:    "User",
	}

	data, err := h.AuthRepository.Register(user)
	data.CreatedAt = time.Now()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	// Mengenerate token jwt nya
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // 3 jam token expired nya

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	registerResponse := authdto.RegisterResponse{
		ID:       data.ID,
		FullName: data.FullName,
		LastName: data.LastName,
		Address:  data.Address,
		Phone:    data.Phone,
		Email:    data.Email,
		Avatar:   data.Avatar,
		Status:   data.Status,
		Roles:    data.Roles,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: registerResponse})
}

// Login Handler
func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Check Password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "wrong email or password"})
	}

	// Generate token jwt nya
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // 3 Jam expired tokennya

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		ID:       user.ID,
		FullName: user.FullName,
		LastName: user.LastName,
		Email:    user.Email,
		Phone:    user.Phone,
		Address:  user.Address,
		Avatar:   user.Avatar,
		Status:   user.Status,
		Roles:    user.Roles,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: loginResponse})

}

// Check-Auth Handler
func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: user})
}
