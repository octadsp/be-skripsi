package models

import "time"

type CarYear struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(100)"`
	Tipe      string    `json:"tipe" gorm:"type: varchar(100)"`
	Status    string    `json:"status" gorm:"type: varchar(10)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (CarYear) TableName() string {
	return "car_years"
}
