package notificationsdto

type NotifReq struct {
	UserID  uint32 `json:"user_id"`
	Message string `json:"message"`
}