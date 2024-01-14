package reservationitemsdto

type ReservationItemReq struct {
	ReservationID       uint32 `json:"reservation_id" form:"reservation_id"`
	DemageSubCategoryID uint32 `json:"demage_sub_category_id" form:"demage_sub_category_id"`
	Image               string `json:"image" form:"image"`
	Price               int64  `json:"price" form:"price"`
	Status              bool   `json:"status" form:"status"`
	PostToUser          string `json:"post_to_user"`
}

type ReservationItemReqUpdate struct {
	ReservationID       uint32 `json:"reservation_id" form:"reservation_id"`
	DemageSubCategoryID uint32 `json:"demage_sub_category_id" form:"demage_sub_category_id"`
	Image               string `json:"image" form:"image"`
	Price               int64  `json:"price" form:"price"`
	Status              bool   `json:"status" form:"status"`
	PostToUser          string `json:"post_to_user"`
}
