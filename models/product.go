package models

import (
	"time"
)

type Product struct {
	ID              string         `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProductName     string         `json:"product_name" gorm:"type: varchar(140); unique"`
	CategoryID      string         `json:"category_id" gorm:"type: varchar(140)"`
	Category        Category       `json:"category" gorm:"foreignKey:CategoryID"`
	BrandID         string         `json:"brand_id" gorm:"type: varchar(140)"`
	Brand           Brand          `json:"brand" gorm:"foreignKey:BrandID"`
	Description     string         `json:"description" gorm:"type: text"`
	Price           int64          `json:"price"`
	InstallationFee int64          `json:"installation_fee"`
	OpeningStock    int64          `json:"opening_stock" gorm:"default:0"`
	Stock           int64          `json:"stock" gorm:"default:0"`
	Creation        time.Time      `json:"creation" gorm:"autoCreateTime"`
	Modified        time.Time      `json:"modified" gorm:"autoCreateTime"`
	ProductImage    []ProductImage `json:"product_image" gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "product"
}
