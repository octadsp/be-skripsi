package userdto

type UserDetailUpdateRequest struct {
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}

type NewUserAddressRequest struct {
	ProvinceID     string `json:"province_id" form:"province_id" validate:"required"`
	RegencyID      string `json:"regency_id" form:"regency_id" validate:"required"`
	DistrictID     string `json:"district_id" form:"district_id" validate:"required"`
	AddressLine    string `json:"address_line" form:"address_line" validate:"required"`
	DefaultAddress bool   `json:"default_address" form:"default_address"`
}

type UpdateUserAddressRequest struct {
	ProvinceID     string `json:"province_id" form:"province_id" validate:"required"`
	RegencyID      string `json:"regency_id" form:"regency_id" validate:"required"`
	DistrictID     string `json:"district_id" form:"district_id" validate:"required"`
	AddressLine    string `json:"address_line" form:"address_line" validate:"required"`
	DefaultAddress bool   `json:"default_address" form:"default_address"`
}
