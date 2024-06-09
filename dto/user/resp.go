package userdto

import "time"

// User Message
type UserInboxResponse struct {
	Customer    string `json:"customer"`
	Admin       string `json:"admin"`
	LastMessage string `json:"last_message"`
	TotalUnread int    `json:"total_unread"`
}

type UserChatLogResponse struct {
	ID       string    `json:"id"`
	Admin    string    `json:"admin"`
	Customer string    `json:"customer"`
	Sender   string    `json:"sender"`
	Message  string    `json:"message"`
	IsRead   bool      `json:"is_read"`
	Creation time.Time `json:"creation"`
	Modified time.Time `json:"modified"`
}
