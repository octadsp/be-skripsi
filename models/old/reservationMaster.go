package models

import "time"

type ReservationMaster struct {
	ID                     uint32               `json:"id" gorm:"primary_key:auto_increment"`
	KodeOrder              string               `json:"kode_order" gorm:"type: varchar(200)"`
	Status                 string               `json:"status" gorm:"type: varchar(10)"`
	OrderMasuk             time.Time            `json:"order_masuk" gorm:"autoCreateTime"`
	OrderProses            time.Time            `json:"order_proses" gorm:"autoCreateTime"`
	OrderSelesai           time.Time            `json:"order_selesai" gorm:"autoCreateTime"`
	UserID                 uint32               `json:"user_id"`
	User                   User                 `json:"user" gorm:"foreignKey:UserID"`
	ReservationVehicleID   uint32               `json:"reservation_vehicle_id"`
	ReservationVehicle     ReservationVehicle   `json:"reservation_vehicle" gorm:"foreignKey:ReservationVehicleID"`
	ReservationInsuranceID uint32               `json:"reservation_insurance_id"`
	ReservationInsurance   ReservationInsurance `json:"reservation_insurance" gorm:"foreignKey:ReservationInsuranceID"`
	ReservationItemID      uint32               `json:"reservation_item_id"`
	ReservationItem        ReservationItem      `json:"reservation_item" gorm:"foreignKey:ReservationItemID"`
	CreatedAt              time.Time            `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              time.Time            `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationMaster) TableName() string {
	return "reservation_masters"
}
