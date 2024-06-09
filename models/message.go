package models

import "time"

type Message struct {
	ID           string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	Admin        string    `json:"admin" gorm:"type: varchar(140)"`
	UserAdmin    User      `json:"user_admin" gorm:"foreignKey:Admin"`
	Customer     string    `json:"customer" gorm:"type: varchar(140)"`
	UserCustomer User      `json:"user_customer" gorm:"foreignKey:Customer"`
	Sender       string    `json:"sender" gorm:"type: varchar(140)"`
	Message      string    `json:"message" gorm:"type:text"`
	IsRead       bool      `json:"is_read" default:"false"`
	Creation     time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Message) TableName() string {
	return "message"
}
