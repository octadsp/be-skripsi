package carclassdto

type CarClassResp struct {
	CarBrandID uint32 `json:"car_brand_id" form:"car_brand_id"`
	CarTypeID  uint32 `json:"car_type_id" form:"car_type_id"`
	Golongan   string `json:"golongan" form:"golongan"`
	Status     string `json:"status" form:"status"`
}
