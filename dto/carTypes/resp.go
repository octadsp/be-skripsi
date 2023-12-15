package cartypesdto

type CarTypeResp struct {
	ID     uint32 `json:"car_type_id" form:"car_type_id"`
	Name   string `json:"name" form:"name"`
	Tipe   string `json:"tipe" form:"tipe"`
	Status string `json:"status" form:"status"`
}