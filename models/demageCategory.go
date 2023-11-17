package models

import "time"

type DemageCategory struct {
	ID                  uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Kode                string    `json:"kode" gorm:"type: varchar(20)"`
	Name                string    `json:"name" gorm:"type: varchar(100)"`
	Status              string    `json:"status" gorm:"type: varchar(10)"`
	DemageSubCategoryID uint32    `json:"demage_sub_category_id"`
	PriceListID         uint32    `json:"price_list_id"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (DemageCategory) TableName() string {
	return "demage_categories"
}
