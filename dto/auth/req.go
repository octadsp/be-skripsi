package authdto

type RegisterRequest struct {
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
