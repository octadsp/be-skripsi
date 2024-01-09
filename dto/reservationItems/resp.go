package reservationitemsdto

type ReservationItemResp struct {
	ID                  uint32 `json:"id"`
	ReservationID       uint32 `json:"reservation_id"`
	Image               string `json:"image"`
	DemageSubCategoryID uint32 `json:"demage_sub_category_id"`
	Price               int64  `json:"price"`
	Status              bool   `json:"status"`
}
