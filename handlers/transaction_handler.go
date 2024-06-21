package handlers

import (
	dto "be-skripsi/dto/results"
	transactionDto "be-skripsi/dto/transaction"
	"be-skripsi/models"
	contains "be-skripsi/pkg/contains"
	errors "be-skripsi/pkg/error"
	repository "be-skripsi/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	CartRepository                repository.CartRepository
	DeliveryFareRepository        repository.DeliveryFareRepository
	OrderRepository               repository.OrderRepository
	ProductRepository             repository.ProductRepository
	ProductStockHistoryRepository repository.ProductStockHistoryRepository
	UserAddressRepository         repository.UserAddressRepository
	UserRepository                repository.UserRepository
}

func HandlerTransaction(
	CartRepository repository.CartRepository,
	DeliveryFareRepository repository.DeliveryFareRepository,
	OrderRepository repository.OrderRepository,
	ProductRepository repository.ProductRepository,
	ProductStockHistoryRepository repository.ProductStockHistoryRepository,
	UserAddressRepository repository.UserAddressRepository,
	UserRepository repository.UserRepository,
) *handlerTransaction {
	return &handlerTransaction{
		CartRepository,
		DeliveryFareRepository,
		OrderRepository,
		ProductRepository,
		ProductStockHistoryRepository,
		UserAddressRepository,
		UserRepository,
	}
}

/*
 * 	Cart
 */
func (h *handlerTransaction) AddToCart(c echo.Context) error {
	request := new(transactionDto.NewCartItemRequest)

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

		_, err = h.CartRepository.AddCartItemQty(request.ProductID, userId, request.Qty, basePrice, installationFee, withInstallation)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		cartItemData, err := h.CartRepository.GetCartItem(request.ProductID, userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: cartItemData})
	}
}

func (h *handlerTransaction) UpdateCart(c echo.Context) error {
	request := new(transactionDto.UpdateCartItemRequest)

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

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: cartItemsData})
}

func (h *handlerTransaction) DeleteCart(c echo.Context) error {
	cartId := c.Param("id")

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	_, err = h.CartRepository.GetCartItemByID(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	_, err = h.CartRepository.DeleteCartItemByID(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: "Success delete cart item"})
}

/*
 * 	Delivery Fare
 */
func (h *handlerTransaction) AddDeliveryFare(c echo.Context) error {
	request := new(transactionDto.NewDeliveryFareRequest)

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	// Validate Province and Regency
	regencyData, err := h.UserAddressRepository.GetRegencyByID(request.RegencyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if regencyData.ProvinceID != request.ProvinceID {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid Province & Regency combination"})
	}

	// Check if exist
	_, err = h.DeliveryFareRepository.GetDeliveryFare(request.ProvinceID, request.RegencyID)
	if err == nil {
		// Existing Delivery Fare
		return c.JSON(http.StatusFound, dto.ErrorResult{Status: http.StatusFound, Message: "Delivery Fare is already exist"})
	}

	deliveryFare := &models.DeliveryFare{
		ID:          uuid.New().String()[:8],
		ProvinceID:  request.ProvinceID,
		RegencyID:   request.RegencyID,
		DeliveryFee: request.DeliveryFee,
	}

	_, err = h.DeliveryFareRepository.AddDeliveryFare(*deliveryFare)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	deliveryFareData, err := h.DeliveryFareRepository.GetDeliveryFare(request.ProvinceID, request.RegencyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: deliveryFareData})
}

func (h *handlerTransaction) GetDeliveryFare(c echo.Context) error {
	provinceId := c.QueryParam("province_id")
	if provinceId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Province ID not defined"})
	}
	regencyId := c.QueryParam("regency_id")
	if regencyId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Regency ID not defined"})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)
	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Validate Province and Regency
	regencyData, err := h.UserAddressRepository.GetRegencyByID(regencyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if regencyData.ProvinceID != provinceId {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid Province & Regency combination"})
	}

	deliveryFareData, err := h.DeliveryFareRepository.GetDeliveryFare(provinceId, regencyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: deliveryFareData})
}

func (h *handlerTransaction) GetDeliveryFares(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	deliveryFaresData, err := h.DeliveryFareRepository.GetDeliveryFares()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: deliveryFaresData})
}

func (h *handlerTransaction) UpdateDeliveryFare(c echo.Context) error {
	request := new(transactionDto.UpdateDeliveryFareRequest)

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	deliveryFareId := c.Param("id")

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	// Check if exist
	deliveryFareData, err := h.DeliveryFareRepository.GetDeliveryFareByID(deliveryFareId)
	if err != nil {
		// Existing Delivery Fare
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: http.StatusNotFound, Message: "Delivery Fare is not found"})
	}

	// Validate Province and Regency
	regencyData, err := h.UserAddressRepository.GetRegencyByID(deliveryFareData.RegencyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if regencyData.ProvinceID != deliveryFareData.ProvinceID {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid Province & Regency combination"})
	}

	deliveryFare := &models.DeliveryFare{
		ID:          deliveryFareId,
		DeliveryFee: request.DeliveryFee,
	}

	_, err = h.DeliveryFareRepository.UpdateDeliveryFare(*deliveryFare)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	deliveryFareReturnData, err := h.DeliveryFareRepository.GetDeliveryFare(deliveryFareData.ProvinceID, deliveryFareData.RegencyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: deliveryFareReturnData})
}

/*
 * 	Order
 */
func (h *handlerTransaction) NewOrder(c echo.Context) error {
	request := new(transactionDto.NewOrderRequest)

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

	// OK Validate userAddressId
	userAddressId := request.UserAddressID
	_, err = h.UserAddressRepository.GetUserAddressByID(userAddressId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// OK Validate deliveryFareId
	deliveryFareId := request.DeliveryFareID
	deliveryFareData, err := h.DeliveryFareRepository.GetDeliveryFareByID(deliveryFareId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	// OK Get delivery_fee
	deliveryFee := deliveryFareData.DeliveryFee
	request.DeliveryFee = deliveryFee

	// OK Validate each product_id
	orderItemsTotal := 0
	cartItems := request.CartItems
	for i := 0; i < len(cartItems); i++ {
		// OK Recalculate sub total for each product
		cartItem := cartItems[i]

		cartID := cartItem.CartID
		cartData, err := h.CartRepository.GetCartItemByID(cartID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		orderProductId := cartData.ProductID
		orderQty := cartData.Qty
		orderWithInstallation := cartData.WithInstallation

		productData, err := h.ProductRepository.GetProduct(orderProductId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		basePrice := productData.Price
		installationFee := productData.InstallationFee

		if !orderWithInstallation {
			installationFee = 0
		}

		cartItem.ProductID = orderProductId
		cartItem.Qty = orderQty
		cartItem.WithInstallation = orderWithInstallation

		cartItem.SubTotal = (orderQty * basePrice) + installationFee
		orderItemsTotal += int(cartItem.SubTotal)
	}

	// OK Calculate total (sub total products + delivery_fee)
	request.OrderTotal = int64(orderItemsTotal) + deliveryFee

	// OK Create Order
	order := &models.Order{
		ID:             uuid.New().String()[:8],
		UserID:         userId,
		UserAddressID:  userAddressId,
		DeliveryFareID: deliveryFareId,
		SubTotal:       int64(orderItemsTotal),
		DeliveryFee:    deliveryFee,
		Total:          request.OrderTotal,
		Status:         "WAITING FOR ORDER CONFIRMATION",
		EstimatedDeliveryAt: func() *time.Time {
			if request.EstimatedDeliveryAt.IsZero() {
				return nil
			}
			return &request.EstimatedDeliveryAt
		}(),
	}

	_, err = h.OrderRepository.CreateOrder(*order)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// OK Create Order Item (products = order item)
	for i := 0; i < len(cartItems); i++ {
		orderItem := cartItems[i]

		orderItemReq := &models.OrderItem{
			ID:               uuid.New().String()[:8],
			OrderID:          order.ID,
			ProductID:        orderItem.ProductID,
			WithInstallation: orderItem.WithInstallation,
			Qty:              orderItem.Qty,
			SubTotal:         orderItem.SubTotal,
		}

		_, err = h.OrderRepository.CreateOrderItem(*orderItemReq)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		// OK Remove cart item
		_, err = h.CartRepository.DeleteCartItemByID(orderItem.CartID)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
	}

	orderData, err := h.OrderRepository.GetOrderByID(order.ID)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: orderData})
}

func (h *handlerTransaction) UpdateOrder(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	orderStatus := c.QueryParam("status")
	if orderStatus != "" && !contains.Contains(contains.OrderStatuses(), orderStatus) {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid order status"})
	}

	orderId := c.Param("id")
	orderData, err := h.OrderRepository.GetOrderByID(orderId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	orderDataUpdate := &models.Order{
		Status: orderStatus,
	}

	_, err = h.OrderRepository.UpdateOrderByID(orderData.ID, *orderDataUpdate)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if orderData.Status == "WAITING FOR ORDER CONFIRMATION" && orderStatus == "WAITING FOR PAYMENT" {
		if userData.Role != "ADMIN" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
		}

		orderItemsData, err := h.OrderRepository.GetOrderItemsByOrderID(orderId)
		if err != nil {
			// Handle the error
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		for i := 0; i < len(orderItemsData); i++ {
			orderItem := orderItemsData[i]

			orderItemProductID := orderItem.ProductID
			orderItemQty := orderItem.Qty

			productData, err := h.ProductRepository.GetProduct(orderItemProductID)
			if err != nil {
				// Handle the error
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
			previousProductData := productData
			newStock := previousProductData.Stock - orderItemQty

			// Handle product stock deduction
			// OK Insert Product Stock History
			productStockHistory := &models.ProductStockHistory{
				ID:            uuid.New().String()[:8],
				ProductID:     orderItemProductID,
				PreviousStock: previousProductData.Stock,
				NewStock:      newStock,
			}
			_, err = h.ProductStockHistoryRepository.InsertProductStockHistory(*productStockHistory)
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResultJSON{Status: http.StatusBadRequest, Message: err.Error()})
			}

			// OK Update Product Stock
			_, err = h.ProductRepository.UpdateProductStock(orderItemProductID, "-", orderItemQty, models.Product{})
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
		}
	}

	orderDataResponse, err := h.OrderRepository.GetOrderByID(orderId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orderDataResponse})
}

func (h *handlerTransaction) GetOrders(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	orderStatus := c.QueryParam("status")
	if orderStatus != "" && !contains.Contains(contains.OrderStatuses(), orderStatus) {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid order status"})
	}

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	ordersData, err := h.OrderRepository.GetOrdersByUserID(userId, orderStatus)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ordersData})
}

func (h *handlerTransaction) GetOrder(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	orderId := c.Param("id")
	orderData, err := h.OrderRepository.GetOrderByID(orderId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orderData})
}

func (h *handlerTransaction) AdminGetOrders(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	orderStatus := c.QueryParam("status")
	if orderStatus != "" && !contains.Contains(contains.OrderStatuses(), orderStatus) {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid order status"})
	}

	ordersData, err := h.OrderRepository.GetOrdersAdmin(orderStatus)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: ordersData})
}

/*
 * 	Payment
 */
func (h *handlerTransaction) SubmitNewPayment(c echo.Context) error {
	orderId := c.Param("id")

	imageFile := c.Get("image").(string)
	request := transactionDto.NewOrderPaymentRequest{
		Image: imageFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, dto.ErrorResultJSON{Status: http.StatusNotAcceptable, Message: errors.ValidationErrors(err)})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err = h.UserRepository.GetUserByID(userId)
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
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "skripsi-anstem-payment"})

	if err != nil {
		fmt.Println(err.Error())
	}

	orderPayment := &models.OrderPayment{
		ID:       uuid.New().String()[:8],
		OrderID:  orderId,
		ImageURL: resp.SecureURL,
		Status:   "WAITING FOR PAYMENT CONFIRMATION",
	}
	_, err = h.OrderRepository.CreateOrderPayment(*orderPayment)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	now := time.Now()
	orderDataUpdate := &models.Order{
		Status: orderPayment.Status,
		PaidAt: &now,
	}
	_, err = h.OrderRepository.UpdateOrderByID(orderPayment.OrderID, *orderDataUpdate)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	orderPaymentDataResponse, err := h.OrderRepository.GetOrderPaymentByOrderID(orderPayment.OrderID)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: http.StatusCreated, Data: orderPaymentDataResponse})
}

func (h *handlerTransaction) GetAllPayment(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	orderPaymentStatus := c.QueryParam("status")
	if orderPaymentStatus != "" && !contains.Contains(contains.OrderPaymentStatuses(), orderPaymentStatus) {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid order payment status"})
	}

	orderPaymentDatas, err := h.OrderRepository.GetAllOrderPayments(orderPaymentStatus)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orderPaymentDatas})
}

func (h *handlerTransaction) GetPaymentByOrderID(c echo.Context) error {
	orderId := c.Param("id")

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	_, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	orderPaymentData, err := h.OrderRepository.GetOrderPaymentByOrderID(orderId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orderPaymentData})
}

func (h *handlerTransaction) GetPaymentByPaymentID(c echo.Context) error {
	paymentId := c.Param("id")

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(string)

	userData, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	if userData.Role != "ADMIN" {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Unauthorized user action"})
	}

	orderPaymentData, err := h.OrderRepository.GetOrderPaymentByID(paymentId)
	if err != nil {
		// Handle the error
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orderPaymentData})
}

func (h *handlerTransaction) UpdatePaymentByPaymentID(c echo.Context) error {
	return nil
}
