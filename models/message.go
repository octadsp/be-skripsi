package models

import "time"

type Message struct {
	ID           string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	Sender       string    `json:"sender" gorm:"type: varchar(140)"`
	UserSender   User      `json:"user_sender" gorm:"foreignKey:Sender"`
	Receiver     string    `json:"receiver" gorm:"type: varchar(140)"`
	UserReceiver User      `json:"user_receiver" gorm:"foreignKey:Receiver"`
	Message      string    `json:"message" gorm:"type:text"`
	IsRead       bool      `json:"is_read" default:"false"`
	Creation     time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Message) TableName() string {
	return "message"
}
