package priceListsdto

type PriceListReq struct {
	DemageSubCategoryID uint32 `json:"demage_sub_category_id" form:"demage_sub_category_id" validate:"required"`
	CarClassID          uint32 `json:"car_class_id" form:"car_class_id" validate:"required"`
	Price               int64  `json:"price" form:"price" validate:"required"`
	Status              string `json:"status" form:"status"`
}
