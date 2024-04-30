package productdto

type NewCategoryRequest struct {
	CategoryName string `json:"category_name" form:"category_name" validate:"required"`
}

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name" form:"category_name" validate:"required"`
}
