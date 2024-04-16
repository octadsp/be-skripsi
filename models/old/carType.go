package models

import "time"

type CarType struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(100)"`
	Tipe      string    `json:"tipe" gorm:"type: varchar(100)"`
	Status    string    `json:"status" gorm:"type: varchar(10)"`
	// CarClassID uint32    `json:"car_class_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

type CarTypeResponse struct {
	Name   string `json:"name"`
	Tipe   string `json:"tipe"`
	Status string `json:"status"`
}

func (CarType) TableName() string {
	return "car_types"
}
