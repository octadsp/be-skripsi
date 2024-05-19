package models

import "time"

type ProductImage struct {
	ID        string  `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProductID string  `json:"product_id" gorm:"type: varchar(140); unique "`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	ImageURL  string  `json:"image_url" gorm:"type: varchar(140); unique "`
	// Idx       int64     `json:"idx"`
	Creation time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (ProductImage) TableName() string {
	return "product_image"
}
