package reservationitemsdto

type ReservationItemResp struct {
	ID     uint32 `json:"id"`
	Item   string `json:"item"`
	Price  int64  `json:"price"`
	Status string `json:"status"`
}
