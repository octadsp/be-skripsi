package demagesubcategoriesdto

type DemageSubCategoryReq struct {
	DemageCategoryID uint32 `json:"demage_category_id" form:"demage_category_id" validate:"required"`
	Name             string `json:"name" form:"name" validate:"required"`
	Status           string `json:"status" form:"status"`
}
