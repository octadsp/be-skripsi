package handlers

import (
	dto "be-skripsi/dto/results"
	cartDto "be-skripsi/dto/transaction"
	"be-skripsi/models"
	errors "be-skripsi/pkg/error"
	repository "be-skripsi/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	CartRepository    repository.CartRepository
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
}

func HandlerTransaction(CartRepository repository.CartRepository, ProductRepository repository.ProductRepository, UserRepository repository.UserRepository) *handlerTransaction {
	return &handlerTransaction{CartRepository, ProductRepository, UserRepository}
}

func (h *handlerTransaction) AddToCart(c echo.Context) error {
	request := new(cartDto.NewCartItemRequest)

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	productData, err := h.ProductRepository.GetProduct(request.ProductID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	basePrice := productData.Price
	installationFee := productData.InstallationFee

	// Validation if ProductID is already exist in UserID cart
	cartItemData, err := h.CartRepository.GetCartItem(request.ProductID, userId)
	if err != nil {
		// New Cart Item

		// Check if product is available
		if productData.Stock < request.Qty {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Insufficient Stock"})
		}

		withInstallation := request.WithInstallation

		cartItem := &models.CartItem{
			ID:               uuid.New().String()[:8],
			UserID:           userId,
			ProductID:        request.ProductID,
			WithInstallation: withInstallation,
			Qty:              request.Qty,
		}

		_, err = h.CartRepository.CreateCartItem(*cartItem, basePrice, installationFee)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		cartItemsData, err := h.CartRepository.GetCartItems(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: cartItemsData})
	} else {
		// Existing Cart Item

		// Check if product is available
		if productData.Stock < request.Qty {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Insufficient Stock"})
		}

		withInstallation := cartItemData.WithInstallation

		_, err = h.CartRepository.AddCartItemQty(request.ProductID, userId, 1, basePrice, installationFee, withInstallation)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		cartItemData, err := h.CartRepository.GetCartItem(request.ProductID, userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: cartItemData})
	}
}

func (h *handlerTransaction) UpdateCart(c echo.Context) error {
	request := new(cartDto.UpdateCartItemRequest)

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	productData, err := h.ProductRepository.GetProduct(request.ProductID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	basePrice := productData.Price
	installationFee := productData.InstallationFee

	// Check if product is available
	if productData.Stock < request.Qty {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Insufficient Stock"})
	}

	cartItemData, err := h.CartRepository.GetCartItem(request.ProductID, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	withInstallation := cartItemData.WithInstallation

	_, err = h.CartRepository.UpdateCartItem(request.ProductID, userId, request.Qty, basePrice, installationFee, withInstallation)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	cartItemDataResponse, err := h.CartRepository.GetCartItem(request.ProductID, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: cartItemDataResponse})
}

func (h *handlerTransaction) GetCart(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	cartItemsData, err := h.CartRepository.GetCartItems(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: cartItemsData})
}
