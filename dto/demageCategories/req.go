package demagecategoriesdto

type DemageCategoryReq struct {
	Kode   string `json:"kode" form:"kode" validate:"required"`
	Name   string `json:"name" form:"name" validate:"required"`
	Status string `json:"status" form:"status"`
}
