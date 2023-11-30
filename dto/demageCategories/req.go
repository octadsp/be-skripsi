package demagecategoriesdto

type DemageCategoryReq struct {
	Kode   string `json:"kode" form:"kode"`
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
}
