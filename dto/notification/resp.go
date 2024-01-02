package notificationsdto

import "time"

type NotifResp struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	AuthorBy  string    `json:"author_by"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
