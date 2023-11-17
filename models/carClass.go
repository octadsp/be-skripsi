package models

import "time"

type CarClass struct {
	ID          uint32    `json:"id" gorm:"primary_key:auto_increment"`
	CarBrandID  uint32    `json:"car_brand_id"`
	CarBrand    CarBrand  `json:"car_brand" gorm:"foreignKey:CarClassID"`
	CarTypeID   uint      `json:"car_type_id"`
	CarType     CarType   `json:"car_type" gorm:"foreignKey:CarClassID"`
	Golongan    string    `json:"golongan" gorm:"type: varchar(100)"`
	Status      string    `json:"status" gorm:"type: varchar(10)"`
	PriceListID uint32    `json:"price_list_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (CarClass) TableName() string {
	return "car_class"
}
