package models

import "time"

type MasterProvince struct {
	ID           string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProvinceName string    `json:"province_name" gorm:"type: varchar(140); unique "`
	IsActive     bool      `json:"is_active"`
	Creation     time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
