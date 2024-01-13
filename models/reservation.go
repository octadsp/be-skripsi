package models

import "time"

type Reservation struct {
	//Master
	ID           uint32    `json:"id" gorm:"primary_key:auto_increment"`
	KodeOrder    string    `json:"kode_order" gorm:"type: varchar(100)"`
	Status       string    `json:"status" gorm:"type: varchar(10)"`
	SubStatus    string    `json:"sub_status" gorm:"type: varchar(10)"`
	OrderMasuk   time.Time `json:"order_masuk"`
	OrderProses  time.Time `json:"order_proses"`
	OrderSelesai time.Time `json:"order_selesai"`
	UserID       uint32    `json:"user_id"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`

	//Vehicle
	CarBrand string `json:"car_brand" gorm:"type: varchar(50)"`
	CarType  string `json:"car_type" gorm:"type: varchar(50)"`
	CarYear  string `json:"car_year" gorm:"type: varchar(5)"`
	CarColor string `json:"car_color" gorm:"type: varchar(50)"`

	IsInsurance int `json:"is_insurance"`

	//Insurance
	InsuranceName  string `json:"insurance_name" gorm:"type: varchar(100)"`
	EventDate      string `json:"event_date" gorm:"type: varchar(50)"`
	Place          string `json:"place" gorm:"type: varchar(50)"`
	Time           string `json:"time" gorm:"type: varchar(10)"`
	DrivingSpeed   string `json:"driver_speed" gorm:"type: varchar(50)"`
	DriverName     string `json:"driver_name" gorm:"type: varchar(50)"`
	DriverRelation string `json:"driver_relation" gorm:"type: varchar(10)"`
	DriverJob      string `json:"driver_job" gorm:"type: varchar(20)"`
	DriverAge      string `json:"driver_age" gorm:"type: varchar(5)"`
	DriverLicense  string `json:"driver_license" gorm:"type: varchar(10)"`

	TotalItem  int `json:"total_item"`
	TotalPrice int `json:"total_price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
