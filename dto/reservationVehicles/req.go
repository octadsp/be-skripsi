package reservationvehiclesdto

type ReservationVehicleReq struct {
	CarBrand string `json:"car_brand" form:"car_brand" validate:"required"`
	CarType  string `json:"car_type" form:"car_type" validate:"required"`
	CarYear  string `json:"car_year" form:"car_year" validate:"required"`
	CarColor string `json:"car_color" form:"car_color" validate:"required"`
	Status   string `json:"status" form:"status"`
}

type ReservationVehicleReqUpdate struct {
	CarBrand string `json:"car_brand" form:"car_brand"`
	CarType  string `json:"car_type" form:"car_type"`
	CarYear  string `json:"car_year" form:"car_year"`
	CarColor string `json:"car_color" form:"car_color"`
	Status   string `json:"status" form:"status"`
}
