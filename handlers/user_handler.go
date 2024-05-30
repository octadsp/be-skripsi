package handlers

import (
	dto "be-skripsi/dto/results"

	"be-skripsi/repositories"

	"net/http"

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
