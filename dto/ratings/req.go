package ratingsdto

type RatingReq struct {
	UserID        uint32 `json:"user_id"`
	ReservationID uint32 `json:"reservation_id"`
	Rating        int    `json:"rating"`
	RatingName    string `json:"rating_name"`
}
