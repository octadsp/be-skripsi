package models

import "time"

type Rating struct {
	ID            uint32      `json:"id" gorm:"primary_key:auto_increment"`
	UserID        uint32      `json:"user_id"`
	User          User        `json:"user" gorm:"foreignKey: UserID"`
	ReservationID uint32      `json:"reservation_id"`
	Reservation   Reservation `json:"reservation" gorm:"foreignKey: ReservationID"`
	Rating        int         `json:"rating"`
	RatingName    string      `json:"rating_name" gorm:"type: varchar(50)"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"autoCreateTime"`
}

func (Rating) TableName() string {
	return "ratings"
}
