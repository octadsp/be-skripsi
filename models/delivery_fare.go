package models

import "time"

type DeliveryFare struct {
	ID          string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProvinceID  string         `json:"province_id" gorm:"type: varchar(100); unique "`
	Province    MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	CityID      string         `json:"city_id" gorm:"type: varchar(50)"`
	City        MasterCity     `json:"city" gorm:"foreignKey:CityID"`
	DeliveryFee int64          `json:"delivery_fee"`
	Creation    time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified    time.Time      `json:"modified" gorm:"autoCreateTime"`
}

func (DeliveryFare) TableName() string {
	return "delivery_fare"
}
