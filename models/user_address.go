package models

import "time"

type UserAddress struct {
	ID             string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	UserID         string         `json:"user_id" gorm:"type: varchar(100); unique "`
	User           User           `json:"user" gorm:"foreignKey:UserID"`
	ProvinceID     string         `json:"province_id" gorm:"type: varchar(100)"`
	Province       MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	RegencyID      string         `json:"regency_id" gorm:"type: varchar(50)"`
	Regency        MasterRegency  `json:"regency" gorm:"foreignKey:RegencyID"`
	DistrictID     string         `json:"district_id" gorm:"type: varchar(50)"`
	District       MasterDistrict `json:"district" gorm:"foreignKey:DistrictID"`
	AddressLine    string         `json:"address_line" gorm:"type: varchar(50)"`
	DefaultAddress bool           `json:"default_address" default:"false"`
	Creation       time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified       time.Time      `json:"modified" gorm:"autoCreateTime"`
}

// default_address = 1/0

func (UserAddress) TableName() string {
	return "user_address"
}
