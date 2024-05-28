package models

import "time"

type ProductStockHistory struct {
	ID            string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	ProductID     string    `json:"product_id" gorm:"type: varchar(140)"`
	Product       Product   `json:"product" gorm:"foreignKey:ProductID"`
	PreviousStock int64     `json:"previous_stock"`
	NewStock      int64     `json:"new_stock"`
	Creation      time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified      time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (ProductStockHistory) TableName() string {
	return "product_stock_history"
}
