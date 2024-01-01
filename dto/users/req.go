package usersdto

type UserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	LastName string `json:"lastname" form:"lastname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Avatar   string `json:"image" form:"image"`
	Status   string `json:"status" form:"status"`
	Roles    string `json:"roles" form:"roles"`
}

type UserUpdateRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	LastName string `json:"lastname" form:"lastname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Avatar   string `json:"image" form:"image"`
}
