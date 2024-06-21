package models

import "time"

type Order struct {
	ID                  string       `json:"id" gorm:"primary_key;type:varchar(140)"`
	UserID              string       `json:"user_id" gorm:"type: varchar(140)"`
	User                User         `json:"user" gorm:"foreignKey:UserID"`
	UserAddressID       string       `json:"user_address_id" gorm:"type: varchar(140)"`
	UserAddress         UserAddress  `json:"user_address" gorm:"foreignKey:UserAddressID"`
	DeliveryFareID      string       `json:"delivery_fare_id" gorm:"type: varchar(140)"`
	DeliveryFare        DeliveryFare `json:"delivery_fare" gorm:"foreignKey:DeliveryFareID"`
	SubTotal            int64        `json:"sub_total"`
	DeliveryFee         int64        `json:"delivery_fee"`
	Total               int64        `json:"total"`
	Status              string       `json:"status"`
	EstimatedDeliveryAt *time.Time   `json:"estimated_delivery_at"`
	PaidAt              *time.Time   `json:"paid_at"`
	AcceptedAt          *time.Time   `json:"accepted_at"`
	RejectedAt          *time.Time   `json:"rejected_at"`
	RejectionReason     string       `json:"rejection_reason"`
	DeliveryAt          *time.Time   `json:"delivery_at"`
	DoneAt              *time.Time   `json:"done_at"`
	Creation            time.Time    `json:"creation" gorm:"autoCreateTime"`
	Modified            time.Time    `json:"modified" gorm:"autoCreateTime"`
	OrderItem           []OrderItem  `json:"order_item" gorm:"foreignKey:OrderID"`
}

// status =
// WAITING FOR ORDER CONFIRMATION /
// WAITING FOR PAYMENT /
// WAITING FOR PAYMENT CONFIRMATION /
// ACCEPTED / REJECTED /
// ON DELIVERY /
// WAITING FOR CUSTOMER CONFIRMATION /
// DONE

func (Order) TableName() string {
	return "order"
}
