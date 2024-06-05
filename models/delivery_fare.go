package models

import "time"

type DeliveryFare struct {
	ID          string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProvinceID  string         `json:"province_id" gorm:"type: varchar(2)"`
	Province    MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	RegencyID   string         `json:"regency_id" gorm:"type: varchar(4)"`
	Regency     MasterRegency  `json:"regency" gorm:"foreignKey:RegencyID"`
	DeliveryFee int64          `json:"delivery_fee"`
	Creation    time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified    time.Time      `json:"modified" gorm:"autoCreateTime"`
}

func (DeliveryFare) TableName() string {
	return "delivery_fare"
}
