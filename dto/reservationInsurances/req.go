package reservationinsurancesdto

type ReservationInsuranceReq struct {
	EventDate      string `json:"event_date" form:"event_date" validate:"required"`
	Place          string `json:"place" form:"place" validate:"required"`
	Time           string `json:"time" form:"time" validate:"required"`
	DrivingSpeed   string `json:"driver_speed" form:"driver_speed" validate:"required"`
	DriverName     string `json:"driver_name" form:"driver_name" validate:"required"`
	DriverRelation string `json:"driver_relation" form:"driver_relation" validate:"required"`
	DriverJob      string `json:"driver_job" form:"driver_job" validate:"required"`
	DriverAge      string `json:"driver_age" form:"driver_age" validate:"required"`
	DriverLicense  string `json:"driver_license" form:"driver_license" validate:"required"`
	Status         string `json:"status" form:"status"`
}

type ReservationInsuranceReqUpdate struct {
	EventDate      string `json:"event_date" form:"event_date"`
	Place          string `json:"place" form:"place"`
	Time           string `json:"time" form:"time"`
	DrivingSpeed   string `json:"driver_speed" form:"driver_speed"`
	DriverName     string `json:"driver_name" form:"driver_name"`
	DriverRelation string `json:"driver_relation" form:"driver_relation"`
	DriverJob      string `json:"driver_job" form:"driver_job"`
	DriverAge      string `json:"driver_age" form:"driver_age"`
	DriverLicense  string `json:"driver_license" form:"driver_license"`
	Status         string `json:"status" form:"status"`
}
