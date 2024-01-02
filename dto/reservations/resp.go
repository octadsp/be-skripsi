package reservationsdto

import "time"

type ReservationResp struct {
	//Master
	ID           uint32    `json:"id"`
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

	//Item
	Image string `json:"image"`
	Item  string `json:"item" gorm:"type: varchar(100)"`
	Price int64  `json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
