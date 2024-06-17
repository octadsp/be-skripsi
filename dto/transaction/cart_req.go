package transactiondto

type NewCartItemRequest struct {
	ProductID        string `json:"product_id" form:"product_id" validate:"required"`
	WithInstallation bool   `json:"with_installation" form:"with_installation"`
	Qty              int64  `json:"qty" form:"qty" validate:"required,numeric,min=0"`
}

type UpdateCartItemRequest struct {
	NewCartItemRequest
}
