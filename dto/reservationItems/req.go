package reservationitemsdto

type ReservationItemReq struct {
	Item   string `json:"item" form:"item" validate:"required"`
	Price  int64  `json:"price" form:"price" validate:"required"`
	Status string `json:"status" form:"status"`
}

type ReservationItemReqUpdate struct {
	Item   string `json:"item" form:"item"`
	Price  int64  `json:"price" form:"price"`
	Status string `json:"status" form:"status"`
}
