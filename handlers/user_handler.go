package handlers

import (
	dto "be-skripsi/dto/results"
	userDto "be-skripsi/dto/user"
	"be-skripsi/models"
	errors "be-skripsi/pkg/error"

	"be-skripsi/repositories"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

// User Detail
func (h *handlerUser) UpdateUserDetail(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	request := new(userDto.UserDetailUpdateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	_, err = h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	userDetail := &models.UserDetail{
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
	}

	_, err = h.UserDetailRepository.UpdateUserDetail(userId, *userDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "User Detail updated successfully!"})
}

// User Address
func (h *handlerUser) NewUserAddress(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	request := new(userDto.NewUserAddressRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	userAddress := &models.UserAddress{
		ID:             uuid.New().String()[:8],
		UserID:         userId,
		ProvinceID:     request.ProvinceID,
		RegencyID:      request.RegencyID,
		DistrictID:     request.DistrictID,
		AddressLine:    request.AddressLine,
		DefaultAddress: request.DefaultAddress,
	}

	_, err = h.UserAddressRepository.CreateUserAddress(*userAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	userAddressData, err := h.UserAddressRepository.GetUserAddressByID(userAddress.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: userAddressData})
}

func (h *handlerUser) GetUserAddressByID(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userAddressId := c.Param("id")

	// OK Validate if userId is owner of addressId
	userAddressData, err := h.UserAddressRepository.GetUserAddressByID(userAddressId)
	if userAddressData.UserID != userId {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "User Address ID not found!"})
	} else if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: userAddressData})
}

func (h *handlerUser) GetUserAddresses(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	// OK Get addresses by userId
	userAddressesData, err := h.UserAddressRepository.GetUserAddresses(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: userAddressesData})
}

func (h *handlerUser) UpdateUserAddressByID(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userAddressId := c.Param("id")
	// OK Validate if userId is owner of addressId
	userAddressData, err := h.UserAddressRepository.GetUserAddressByID(userAddressId)
	if userAddressData.UserID != userId {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "User Address ID not found!"})
	} else if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	request := new(userDto.UpdateUserAddressRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	userAddress := &models.UserAddress{
		ProvinceID:  request.ProvinceID,
		RegencyID:   request.RegencyID,
		DistrictID:  request.DistrictID,
		AddressLine: request.AddressLine,
	}

	_, err = h.UserAddressRepository.UpdateUserAddressByID(userAddressId, *userAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "User Address updated successfully!"})
}

func (h *handlerUser) DeleteUserAddressByID(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userAddressId := c.Param("id")
	// OK Validate if userId is owner of addressId
	userAddressData, err := h.UserAddressRepository.GetUserAddressByID(userAddressId)
	if userAddressData.UserID != userId {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "User Address ID not found!"})
	} else if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: err.Error()})
	}

	_, err = h.UserAddressRepository.DeleteUserAddressByID(userAddressId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "User Address deleted successfully!"})
}

// Provinces
func (h *handlerUser) GetProvinces(c echo.Context) error {
	provincesData, err := h.UserAddressRepository.GetProvinces()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: provincesData})
}

func (h *handlerUser) GetProvinceByID(c echo.Context) error {
	id := c.Param("id")
	provinceData, err := h.UserAddressRepository.GetProvinceByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: provinceData})
}

// Regencies
func (h *handlerUser) GetRegencies(c echo.Context) error {
	regenciesData, err := h.UserAddressRepository.GetRegencies()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: regenciesData})
}

func (h *handlerUser) GetRegenciesByProvinceID(c echo.Context) error {
	province_id := c.Param("province_id")
	regenciesData, err := h.UserAddressRepository.GetRegenciesByProvinceID(province_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: regenciesData})
}

func (h *handlerUser) GetRegencyByID(c echo.Context) error {
	id := c.Param("id")
	regencyData, err := h.UserAddressRepository.GetRegencyByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: regencyData})
}

// Districts
func (h *handlerUser) GetDistricts(c echo.Context) error {
	districtsData, err := h.UserAddressRepository.GetDistricts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: districtsData})
}

func (h *handlerUser) GetDistrictsByRegencyID(c echo.Context) error {
	regency_id := c.Param("regency_id")
	districtsData, err := h.UserAddressRepository.GetDistrictsByRegencyID(regency_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: districtsData})
}

func (h *handlerUser) GetDistrictByID(c echo.Context) error {
	id := c.Param("id")
	districtData, err := h.UserAddressRepository.GetDistrictByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: districtData})
}
