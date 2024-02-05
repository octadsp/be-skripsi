package priceListsdto

import "be-skripsi/models"

type PriceListResp struct {
	DemageSubCategoryID uint32          `json:"demage_sub_category_id" form:"demage_sub_category_id"`
	CarClassID          uint32          `json:"car_class_id" form:"car_class_id"`
	CarClass            models.CarClass `json:"car_class"`
	Price               int64           `json:"price" form:"price"`
	Status              string          `json:"status" form:"status"`
}
