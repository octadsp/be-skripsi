package authdto

type RegisterResponse struct {
	ID       uint32 `json:"id"`
	FullName string `json:"fullname"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Avatar   string `json:"image"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
	Token    string `json:"token"`
}

type LoginResponse struct {
	ID       uint32 `json:"id"`
	FullName string `json:"fullname"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Avatar   string `json:"image"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
	Token    string `json:"token"`
	ExpiresIn int64 `json:"expiresIn"`
}
