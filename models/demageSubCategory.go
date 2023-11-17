package models

import "time"

type DemageSubCategory struct {
	ID               uint32         `json:"id" gorm:"primary_key:auto_increment"`
	DemageCategoryID uint32         `json:"demage_category_id"`
	DemageCategory   DemageCategory `json:"demage_category" gorm:"foreignKey:DemageSubCategoryID"`
	Name             string         `json:"name" gorm:"type: varchar(100)"`
	Status           string         `json:"status" gorm:"type: varchar(10)"`
	PriceListID      uint32         `json:"price_list_id"`
	CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"autoCreateTime"`
}

func (DemageSubCategory) TableName() string {
	return "demage_prices"
}
