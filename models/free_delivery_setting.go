package models

import "time"

type FreeDeliverySetting struct {
	ID                 string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	MinimumTransaction int64     `json:"minimum_transaction"`
	Creation           time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified           time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (FreeDeliverySetting) TableName() string {
	return "free_delivery_setting"
}
