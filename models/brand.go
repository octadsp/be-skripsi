package models

import "time"

type Brand struct {
	ID        string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	BrandName string    `json:"brand_name" gorm:"type: varchar(140); unique "`
	Creation  time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified  time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Brand) TableName() string {
	return "brand"
}
