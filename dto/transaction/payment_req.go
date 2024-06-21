package transactiondto

type NewOrderPaymentRequest struct {
	Image string `json:"image" form:"image"`
}

type UpdateOrderPaymentRequest struct {
	RejectionReason string `json:"rejection_reason" form:"rejection_reason"`
}
