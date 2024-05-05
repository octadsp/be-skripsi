package productdto

type NewProductRequest struct {
	ProductName     string `json:"product_name" form:"product_name" validate:"required"`
	BrandID         string `json:"brand_id" form:"brand_id" validate:"required"`
	CategoryID      string `json:"category_id" form:"category_id" validate:"required"`
	Price           int64  `json:"price" form:"price" validate:"required"`
	InstallationFee int64  `json:"installation_fee" form:"installation_fee" validate:"required"`
	OpeningStock    int64  `json:"opening_stock" form:"opening_stock"`
}

type UpdateProductRequest struct {
	ProductName     string `json:"product_name" form:"product_name"`
	BrandID         string `json:"brand_id" form:"brand_id"`
	CategoryID      string `json:"category_id" form:"category_id"`
	Price           int64  `json:"price" form:"price"`
	InstallationFee int64  `json:"installation_fee" form:"installation_fee"`
}
