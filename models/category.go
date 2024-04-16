package models

import "time"

type Category struct {
	ID           string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	CategoryName string    `json:"category_name" gorm:"type: varchar(140); unique "`
	Creation     time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Category) TableName() string {
	return "category"
}
