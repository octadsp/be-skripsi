package transactiondto

import "time"

type NewOrderRequest struct {
	UserAddressID       string             `json:"user_address_id" validate:"required"`
	DeliveryFareID      string             `json:"delivery_fare_id" validate:"required"`
	CartItems           []*CartItemRequest `json:"cart_items" validate:"required,min=1,dive,required"`
	DeliveryFee         int64              `json:"delivery_fee"`
	OrderTotal          int64              `json:"total"`
	EstimatedDeliveryAt time.Time          `json:"estimated_delivery_at"`
}

type CartItemRequest struct {
	CartID           string `json:"id" validate:"required"`
	ProductID        string `json:"product_id"`
	WithInstallation bool   `json:"with_installation"`
	Qty              int64  `json:"qty"`
	SubTotal         int64  `json:"sub_total"`
}
