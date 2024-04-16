package models

import "time"

type MasterDistrict struct {
	ID           string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProvinceID   string         `json:"province_id" gorm:"type: varchar(140); unique "`
	Province     MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	CityID       string         `json:"city_id" gorm:"type: varchar(140); unique "`
	City         MasterCity     `json:"city" gorm:"foreignKey:CityID"`
	DistrictName string         `json:"district_name" gorm:"type: varchar(140); unique "`
	IsActive     bool           `json:"is_active"`
	Creation     time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time      `json:"modified" gorm:"autoCreateTime"`
}

func (MasterDistrict) TableName() string {
	return "master_district"
}
