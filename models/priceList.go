package models

import "time"

type PriceList struct {
	ID               uint32         `json:"id" gorm:"primary_key:auto_increment"`
	DemageCategoryID uint32         `json:"demage_category_id"`
	DemageCategory   DemageCategory `json:"demage_category" gorm:"foreignKey:PriceListID"`
	DemagePriceID    uint32         `json:"demage_price_id"`
	DemagePrice      DemagePrice    `json:"demage_price" gorm:"foreignKey:PriceListID"`
	Status           string         `json:"status" gorm:"type: varchar(10)"`
	CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"autoCreateTime"`
}

func (PriceList) TableName() string {
	return "price_lists"
}
