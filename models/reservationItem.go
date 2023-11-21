package models

import "time"

type ReservationItem struct {
	ID                  uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Item                string    `json:"item" gorm:"type: varchar(100)"`
	Price               int64     `json:"price"`
	Status              string    `json:"status" gorm:"type: varchar(10)"`
	// ReservationMasterID uint32    `json:"reservation_master_id"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationItem) TableName() string {
	return "reservation_items"
}
