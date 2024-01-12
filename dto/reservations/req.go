package reservationsdto

import "time"

type ReservationReq struct {
	//Master
	KodeOrder    string    `json:"kode_order"`
	Status       string    `json:"status"`
	OrderMasuk   time.Time `json:"order_masuk"`
	OrderProses  time.Time `json:"order_proses"`
	OrderSelesai time.Time `json:"order_selesai"`
	UserID       uint32    `json:"user_id"`

	//Vehicle
	CarBrand string `json:"car_brand"`
	CarType  string `json:"car_type"`
	CarYear  string `json:"car_year"`
	CarColor string `json:"car_color"`

	IsInsurance int `json:"is_insurance"`

	//Insurance
	InsuranceName  string `json:"insurance_name"`
	EventDate      string `json:"event_date"`
	Place          string `json:"place"`
	Time           string `json:"time"`
	DrivingSpeed   string `json:"driver_speed"`
	DriverName     string `json:"driver_name"`
	DriverRelation string `json:"driver_relation"`
	DriverJob      string `json:"driver_job"`
	DriverAge      string `json:"driver_age"`
	DriverLicense  string `json:"driver_license"`

	ReservationItemID uint32 `json:"reservation_item_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReservationReqUpdate struct {
	Status       string    `json:"status"`
	OrderProses  time.Time `json:"order_proses"`
	OrderSelesai time.Time `json:"order_selesai"`

	CarBrand string `json:"car_brand"`
	CarType  string `json:"car_type"`
	CarYear  string `json:"car_year"`
	CarColor string `json:"car_color"`

	IsInsurance int `json:"is_insurance"`

	InsuranceName  string `json:"insurance_name"`
	EventDate      string `json:"event_date"`
	Place          string `json:"place"`
	Time           string `json:"time"`
	DrivingSpeed   string `json:"driver_speed"`
	DriverName     string `json:"driver_name"`
	DriverRelation string `json:"driver_relation"`
	DriverJob      string `json:"driver_job"`
	DriverAge      string `json:"driver_age"`
	DriverLicense  string `json:"driver_license"`

	TotalItem  int `json:"total_item"`
	TotalPrice int `json:"total_price"`

	ReservationItemID uint32 `json:"reservation_item_id"`

	UpdatedAt time.Time `json:"updated_at"`
}
