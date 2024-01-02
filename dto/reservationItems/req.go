package reservationitemsdto

type ReservationItemReq struct {
	Item   string `json:"item" form:"item"`
	Image  string `json:"image" form:"image"`
	Price  int64  `json:"price" form:"price"`
	Status string `json:"status" form:"status"`
}

type ReservationItemReqUpdate struct {
	Item   string `json:"item" form:"item"`
	Image  string `json:"image" form:"image"`
	Price  int64  `json:"price" form:"price"`
	Status string `json:"status" form:"status"`
}
