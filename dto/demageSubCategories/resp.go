package demagesubcategoriesdto

type DemageSubCategoryResp struct {
	DemageCategoryID uint32 `json:"demage_category_id" form:"demage_category_id"`
	Name             string `json:"name" form:"name"`
	Status           string `json:"status" form:"status"`
}
