package cartypesdto

type CarTypeReq struct {
	Name string `json:"name" form:"name"`
	Tipe string `json:"tipe" form:"tipe"`
	Status string `json:"status" form:"status"`
}
