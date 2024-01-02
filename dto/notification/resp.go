package notificationsdto

import "time"

type NotifResp struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}