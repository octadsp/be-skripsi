package models

import "time"

type StatusReservation struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Kode      string    `json:"kode" gorm:"type: varchar(20)"`
	Name      string    `json:"name" gorm:"type: varchar(100)"`
	Status    string    `json:"status" gorm:"type: varchar(10)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (StatusReservation) TableName() string {
	return "status_reservations"
}
