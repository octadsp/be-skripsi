package reservationvehiclesdto

type ReservationVehicleResp struct {
	ID       uint32 `json:"id"`
	CarBrand string `json:"car_brand"`
	CarType  string `json:"car_type"`
	CarYear  string `json:"car_year"`
	CarColor string `json:"car_color"`
	Status   string `json:"status"`
}
