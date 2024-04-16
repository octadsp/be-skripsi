package models

import "time"

type ReservationVehicle struct {
	ID       uint32 `json:"id" gorm:"primary_key:auto_increment"`
	CarBrand string `json:"car_brand" gorm:"type: varchar(100)"`
	CarType  string `json:"car_type" gorm:"type: varchar(100)"`
	CarYear  string `json:"car_year" gorm:"type: varchar(5)"`
	CarColor string `json:"car_color" gorm:"type: varchar(100)"`
	Status   string `json:"status" gorm:"type: varchar(10)"`
	// ReservationMasterID uint32    `json:"reservation_master_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationVehicle) TableName() string {
	return "reservation_vehicles"
}
