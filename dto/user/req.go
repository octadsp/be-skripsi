package userdto

// User Detail
type UserDetailUpdateRequest struct {
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}

// User Address
type NewUserAddressRequest struct {
	ProvinceID     string `json:"province_id" form:"province_id" validate:"required"`
	RegencyID      string `json:"regency_id" form:"regency_id" validate:"required"`
	DistrictID     string `json:"district_id" form:"district_id" validate:"required"`
	AddressLine    string `json:"address_line" form:"address_line" validate:"required"`
	DefaultAddress bool   `json:"default_address" form:"default_address"`
}

type UpdateUserAddressRequest struct {
	ProvinceID     string `json:"province_id" form:"province_id"`
	RegencyID      string `json:"regency_id" form:"regency_id"`
	DistrictID     string `json:"district_id" form:"district_id"`
	AddressLine    string `json:"address_line" form:"address_line"`
	DefaultAddress bool   `json:"default_address" form:"default_address"`
}

// User Message
type NewUserMessageRequest struct {
	Message string `json:"message" form:"message" validate:"required"`
}
