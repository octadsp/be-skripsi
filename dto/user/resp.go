package userdto

// User Message
type UserMessagesResponse struct {
	Customer    string `json:"customer"`
	Admin       string `json:"admin"`
	LastMessage string `json:"last_message"`
	TotalUnread int    `json:"total_unread"`
}
