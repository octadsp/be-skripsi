package handlers

import (
	authDto "be-skripsi/dto/auth"
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
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	UserRepository       repository.UserRepository
	UserDetailRepository repository.UserDetailRepository
}

func HandlerAuth(UserRepository repository.UserRepository, UserDetailRepository repository.UserDetailRepository) *handlerAuth {
	return &handlerAuth{UserRepository, UserDetailRepository}
}

// Register Handler
func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authDto.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	/*
	 * 	User
	 */

	// OK Hashing Password
	password, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// OK Compose payload for User
	user := &models.User{
		ID:       uuid.New().String()[:8],
		Email:    request.Email,
		Password: password,
		Role:     "CUSTOMER",
	}

	// OK Insert User
	userData, err := h.UserRepository.CreateUser(*user)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	/*
	 * 	User Detail
	 */
	//  OK Compose payload for User Detail
	userDetail := &models.UserDetail{
		ID:          uuid.New().String()[:8],
		UserID:      userData.ID,
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
	}

	// OK Insert User Detail
	userDetailData, err := h.UserDetailRepository.CreateUserDetail(*userDetail)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	// OK Generate JWT Token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // valid for 3 hours

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	registerResponse := authDto.RegisterResponse{
		ID:          userData.ID,
		FullName:    userDetailData.FullName,
		Email:       userData.Email,
		PhoneNumber: userDetailData.PhoneNumber,
		Role:        userData.Role,
		Token:       token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: registerResponse})
}

// Login Handler
func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authDto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// OK Check email
	userData, err := h.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// OK Get detail
	userDetailData, err := h.UserDetailRepository.GetUserDetail(userData.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// OK Check Password
	isValid := bcrypt.CheckPasswordHash(request.Password, userData.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid credential"})
	}

	// OK Generate JWT Token
	claims := jwt.MapClaims{}
	claims["id"] = userData.ID

	expiresIn := time.Now().Add(time.Hour * 1).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authDto.LoginResponse{
		ID:          userData.ID,
		FullName:    userDetailData.FullName,
		Email:       userData.Email,
		PhoneNumber: userDetailData.PhoneNumber,
		Role:        userData.Role,
		Token:       token,
		ExpiresIn:   expiresIn,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: loginResponse})

}

// Check-Auth Handler
func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, _ := h.UserRepository.GetUserByID(userId)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: userData})
}
