package models

import "time"

type UserAddress struct {
	ID             string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	UserID         string         `json:"user_id" gorm:"type: varchar(100); unique "`
	User           User           `json:"user" gorm:"foreignKey:UserID"`
	ProvinceID     string         `json:"province_id" gorm:"type: varchar(100)"`
	Province       MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	CityID         string         `json:"city_id" gorm:"type: varchar(50)"`
	City           MasterCity     `json:"city" gorm:"foreignKey:CityID"`
	DistrictID     string         `json:"district_id" gorm:"type: varchar(50)"`
	District       MasterDistrict `json:"district" gorm:"foreignKey:DistrictID"`
	AddressLine    string         `json:"address_line" gorm:"type: varchar(50)"`
	DefaultAddress string         `json:"default_address" gorm:"type: varchar(50)"`
	Creation       time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified       time.Time      `json:"modified" gorm:"autoCreateTime"`
}

func (UserAddress) TableName() string {
	return "user_address"
}
