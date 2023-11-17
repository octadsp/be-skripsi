package models

import "time"

type ReservationMaster struct {
	ID                   uint32               `json:"id" gorm:"primary_key:auto_increment"`
	KodeOrder            string               `json:"kode_order" gorm:"type: varchar(200)"`
	Status               string               `json:"status" gorm:"type: varchar(10)"`
	OrderMasuk           time.Time            `json:"order_masuk" gorm:"autoCreateTime"`
	OrderProses          time.Time            `json:"order_proses" gorm:"autoCreateTime"`
	OrderSelesai         time.Time            `json:"order_selesai" gorm:"autoCreateTime"`
	UserID               uint32               `json:"user_id" gorm:"foreignKey:UserID"`
	ReservationVehicle   ReservationVehicle   `json:"reservation_vehicle"`
	ReservationInsurance ReservationInsurance `json:"reservation_insurance"`
	ReservationItem      ReservationItem      `json:"reservation_item"`
	CreatedAt            time.Time            `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time            `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationMaster) TableName() string {
	return "reservation_masters"
}
