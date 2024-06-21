package transactiondto

type NewOrderRequest struct {
	UserAddressID  string              `json:"user_address_id" validate:"required"`
	DeliveryFareID string              `json:"delivery_fare_id" validate:"required"`
	OrderItems     []*OrderItemRequest `json:"order_items" validate:"required,min=1,dive,required"`
	DeliveryFee    int64               `json:"delivery_fee"`
	OrderTotal     int64               `json:"total"`
}

type OrderItemRequest struct {
	ProductID        string `json:"product_id" validate:"required"`
	WithInstallation bool   `json:"with_installation"`
	Qty              int64  `json:"qty" validate:"required"`
	SubTotal         int64  `json:"sub_total"`
}
