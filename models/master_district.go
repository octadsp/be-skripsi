package models

import "time"

type MasterDistrict struct {
	ID           string        `json:"id" gorm:"primary_key;type:varchar(140);unique"`
	RegencyID    string        `json:"regency_id" gorm:"type: varchar(140)"`
	Regency      MasterRegency `json:"regency" gorm:"foreignKey:RegencyID"`
	DistrictName string        `json:"district_name" gorm:"type: varchar(140)"`
	IsActive     bool          `json:"is_active" gorm:"default:false"`
	Creation     time.Time     `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time     `json:"modified" gorm:"autoCreateTime"`
}

// is_active = 1/0

func (MasterDistrict) TableName() string {
	return "master_district"
}
