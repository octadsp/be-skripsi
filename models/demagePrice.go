package models

import "time"

type DemagePrice struct {
	ID          uint32 `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" gorm:"type: varchar(100)"`
	Price       int32  `json:"price"`
	Status      string `json:"status" gorm:"type: varchar(10)"`
	PriceListID uint32 `json:"price_list_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (DemagePrice) TableName() string {
	return "demage_prices"
}
