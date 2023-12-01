package reservationinsurancesdto

type ReservationInsuranceResp struct {
	ID             uint32 `json:"id"`
	EventDate      string `json:"event_date"`
	Place          string `json:"place"`
	Time           string `json:"time"`
	DrivingSpeed   string `json:"driver_speed"`
	DriverName     string `json:"driver_name"`
	DriverRelation string `json:"driver_relation"`
	DriverJob      string `json:"driver_job"`
	DriverAge      string `json:"driver_age"`
	DriverLicense  string `json:"driver_license"`
	Status         string `json:"status"`
}
