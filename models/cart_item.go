package models

import "time"

type CartItem struct {
	ID               string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	UserID           string    `json:"user_id" gorm:"type: varchar(140)"`
	User             User      `json:"user" gorm:"foreignKey:UserID"`
	ProductID        string    `json:"product_id" gorm:"type: varchar(140)"`
	Product          Product   `json:"product" gorm:"foreignKey:ProductID"`
	WithInstallation bool      `json:"with_installation"`
	Qty              int64     `json:"qty"`
	SubTotal         int64     `json:"sub_total"`
	Creation         time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified         time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (CartItem) TableName() string {
	return "cart_item"
}
