package ratingsdto

type RatingResp struct {
	UserID     uint32 `json:"user_id"`
	Rating     int    `json:"rating"`
	RatingName string `json:"rating_name"`
}
