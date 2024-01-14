package models

import "time"

type ReservationItem struct {
	ID                  uint32            `json:"id" gorm:"primary_key:auto_increment"`
	ReservationID       uint32            `json:"reservation_id"`
	Reservation         Reservation       `json:"reservatin" gorm:"foreignKey: ReservationID"`
	Image               string            `json:"image"`
	DemageSubCategoryID uint32            `json:"demage_sub_category_id"`
	DemageSubCategory   DemageSubCategory `json:"demage_sub_category" gorm:"foreignKey: DemageSubCategoryID"`
	Price               int64             `json:"price"`
	Status              bool              `json:"status"`
	PostToUser          string            `json:"post_to_user"`
	CreatedAt           time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time         `json:"updated_at" gorm:"autoCreateTime"`
}

func (ReservationItem) TableName() string {
	return "reservation_items"
}
