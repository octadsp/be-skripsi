package reservationitemsdto

type ReservationItemResp struct {
	ID    uint32 `json:"id"`
	Item  string `json:"item"`
	Price string `json:"price"`
}
