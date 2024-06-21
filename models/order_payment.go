package models

import "time"

type OrderPayment struct {
	ID       string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	OrderID  string    `json:"order_id" gorm:"type: varchar(140)"`
	Order    Order     `json:"order" gorm:"foreignKey:OrderID"`
	ImageURL string    `json:"image_url" gorm:"type: varchar(140)"`
	Status   string    `json:"status"`
	Creation time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified time.Time `json:"modified" gorm:"autoCreateTime"`
}

// status = WAITING FOR PAYMENT CONFIRMATION / ACCEPTED / REJECTED

func (OrderPayment) TableName() string {
	return "order_payment"
}
