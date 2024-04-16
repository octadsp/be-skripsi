package models

import "time"

type Product struct {
	ID              string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProductName     string    `json:"product_name" gorm:"type: varchar(140); unique "`
	CategoryID      string    `json:"category_id" gorm:"type: varchar(140); unique "`
	Category        Category  `json:"category" gorm:"foreignKey:CategoryID"`
	BrandID         string    `json:"brand_id" gorm:"type: varchar(140); unique "`
	Brand           Brand     `json:"brand" gorm:"foreignKey:BrandID"`
	Price           int64     `json:"price"`
	InstallationFee int64     `json:"installation_fee"`
	Stock           int64     `json:"stock"`
	Creation        time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified        time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Product) TableName() string {
	return "product"
}
