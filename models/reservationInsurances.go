package models

import "time"

type ReservationInsurance struct {
	ID                  uint32    `json:"id" gorm:"primary_key:auto_increment"`
	EventDate           string    `json:"event_date" gorm:"type: varchar(100)"`
	Place               string    `json:"place" gorm:"type: varchar(100)"`
	Time                string    `json:"time" gorm:"type: varchar(5)"`
	DrivingSpeed        string    `json:"driver_speed" gorm:"type: varchar(100)"`
	DriverName          string    `json:"driver_name" gorm:"type: varchar(10)"`
	DriverRelation      string    `json:"driver_relation" gorm:"type: varchar(10)"`
	DriverJob           string    `json:"driver_job" gorm:"type: varchar(10)"`
	DriverAge           string    `json:"driver_age" gorm:"type: varchar(10)"`
	DriverLicense       string    `json:"driver_license" gorm:"type: varchar(10)"`
	ReservationMasterID uint32    `json:"reservation_master_id"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationInsurance) TableName() string {
	return "reservation_insurances"
}
