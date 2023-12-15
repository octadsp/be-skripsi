package carbrandsdto

type CarBrandResp struct {
	ID uint32 `json:"car_brand_id" form:"car_brand_id"`
	Name string `json:"name" form:"name"`
	Tipe string `json:"tipe" form:"tipe"`
	Status string `json:"status" form:"status"`
}