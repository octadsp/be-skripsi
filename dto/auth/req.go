package authdto

type RegisterRequest struct {
	FullName  string `json:"fullname" form:"fullname" validate:"required"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email" validate:"required"`
	Institute string `json:"institute" form:"institute"`
	Password  string `json:"password" form:"password" validate:"required"`
	Avatar    string `json:"image" form:"image"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Status    string `json:"status" form:"status"`
	Roles     string `json:"roles" form:"roles"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
