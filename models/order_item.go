package models

import "time"

type OrderItem struct {
	ID               string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	OrderID          string    `json:"order_id" gorm:"type: varchar(100); unique "`
	Order            Order     `json:"order" gorm:"foreignKey:OrderID"`
	ProductID        string    `json:"product_id" gorm:"type: varchar(100); unique "`
	Product          Product   `json:"product" gorm:"foreignKey:ProductID"`
	WithInstallation bool      `json:"with_installation"`
	Qty              int64     `json:"qty"`
	SubTotal         int64     `json:"sub_total"`
	Creation         time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified         time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
