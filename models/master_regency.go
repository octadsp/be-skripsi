package models

import "time"

type MasterRegency struct {
	ID          string         `json:"id" gorm:"primary_key;type:varchar(140);unique"`
	ProvinceID  string         `json:"province_id" gorm:"type: varchar(140)"`
	Province    MasterProvince `json:"province" gorm:"foreignKey:ProvinceID"`
	RegencyName string         `json:"regency_name" gorm:"type:varchar(140)"`
	IsActive    bool           `json:"is_active" gorm:"default:false"`
	Creation    time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified    time.Time      `json:"modified" gorm:"autoCreateTime"`
}

// is_active = 1/0

func (MasterRegency) TableName() string {
	return "master_regency"
}
