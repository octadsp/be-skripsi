package models

import "time"

type PriceList struct {
	ID                  uint32            `json:"id" gorm:"primary_key:auto_increment"`
	DemageSubCategoryID uint32            `json:"demage_sub_category_id"`
	DemageSubCategory   DemageSubCategory `json:"demage_sub_category" gorm:"foreignKey:DemageSubCategoryID"`
	CarClassID          uint32              `json:"car_class_id"`
	CarClass            CarClass          `json:"car_class" gorm:"foreignKey:CarClassID"`
	Price               int64             `json:"price"`
	Status              string            `json:"status" gorm:"type: varchar(10)"`
	CreatedAt           time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time         `json:"updated_at" gorm:"autoCreateTime"`
}

func (PriceList) TableName() string {
	return "price_lists"
}
