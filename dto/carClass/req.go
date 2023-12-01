package carclassdto

type CarClassReq struct {
	CarBrandID  uint32 `json:"car_brand_id" form:"car_brand_id" validate:"required"`
	CarTypeID   uint32 `json:"car_type_id" form:"car_type_id" validate:"required"`
	Golongan    string `json:"golongan" form:"golongan" validate:"required"`
	Status      string `json:"status" form:"status"`
	PriceListID uint32 `json:"price_list_id" form:"price_list_id" validate:"required"`
}
