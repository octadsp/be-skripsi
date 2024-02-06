package notificationsdto

type NotifReq struct {
	UserID   uint32 `json:"user_id"`
	AuthorBy string `json:"author_by"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

type NotifReqStatus struct {
	IsRead bool `json:"is_read"`
}
