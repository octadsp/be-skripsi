package transactiondto

type GetDeliveryFareRequest struct {
	ProvinceID string `json:"province_id" form:"province_id" validate:"required"`
	RegencyID  string `json:"regency_id" form:"regency_id" validate:"required"`
}

type UpdateDeliveryFareRequest struct {
	DeliveryFee int64 `json:"delivery_fee" form:"delivery_fee"`
}
type NewDeliveryFareRequest struct {
	GetDeliveryFareRequest
	UpdateDeliveryFareRequest
}
