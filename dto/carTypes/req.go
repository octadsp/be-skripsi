package cartypesdto

type CarTypeReq struct {
	Name string `json:"name" form:"name" validate:"required"`
	Tipe string `json:"tipe" form:"tipe" validate:"required"`
	Status string `json:"status" form:"status" `
}
