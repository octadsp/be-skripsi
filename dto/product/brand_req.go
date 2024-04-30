package productdto

type NewBrandRequest struct {
	BrandName string `json:"brand_name" form:"brand_name" validate:"required"`
}

type UpdateBrandRequest struct {
	BrandName string `json:"brand_name" form:"brand_name" validate:"required"`
}
