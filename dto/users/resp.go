package usersdto

type UserResponse struct {
	ID       uint32 `json:"id"`
	FullName string `json:"fullname"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Avatar   string `json:"image"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
}
