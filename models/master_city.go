package models

import "time"

type MasterCity struct {
	ID         string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProvinceID string         `json:"province_id" gorm:"type: varchar(140); unique "`
	Province   MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	CityName   string         `json:"city_name" gorm:"type: varchar(140); unique "`
	IsActive   bool           `json:"is_active"`
	Creation   time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified   time.Time      `json:"modified" gorm:"autoCreateTime"`
}

// is_active = 1/0

func (MasterCity) TableName() string {
	return "master_city"
}
