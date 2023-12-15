package carclassdto

import "be-skripsi/models"

type CarClassResp struct {
	ID         uint32 `json:"car_class_id" form:"car_class_id"`
	CarBrandID uint32 `json:"car_brand_id" form:"car_brand_id"`
	CarBrand   models.CarBrand `json:"car_brand" form:"car_brand"`
	CarTypeID  uint32 `json:"car_type_id" form:"car_type_id"`
	CarType models.CarType `json:"car_type" form:"car_type"`
	Golongan   string `json:"golongan" form:"golongan"`
	Status     string `json:"status" form:"status"`
}
