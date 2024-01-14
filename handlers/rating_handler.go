package handlers

import (
	ratingdto "be-skripsi/dto/ratings"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerRating struct {
	ratingRepository repositories.RatingRepository
}

func HandlerRating(ratingRepository repositories.RatingRepository) *handlerRating {
	return &handlerRating{ratingRepository}
}

func (h *handlerRating) FindRatingByUser(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("userid"))

	ratings, err := h.ratingRepository.FindRatingByUser(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ratings})
}

func (h *handlerRating) AddRating(c echo.Context) error {

	request := new(ratingdto.RatingReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	rating := models.Rating{
		UserID:     request.UserID,
		Rating:     request.Rating,
		RatingName: request.RatingName,
	}

	data, err := h.ratingRepository.AddRating(rating)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
