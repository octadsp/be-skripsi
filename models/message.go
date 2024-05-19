package models

import "time"

type Message struct {
	ID           string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	Sender       string    `json:"sender" gorm:"type: varchar(100); unique "`
	UserSender   User      `json:"user_sender" gorm:"foreignKey:Sender"`
	Receiver     string    `json:"receiver" gorm:"type: varchar(100); unique "`
	UserReceiver User      `json:"user_receiver" gorm:"foreignKey:Receiver"`
	Message      string    `json:"message" gorm:"type: varchar(100)"`
	IsRead       bool      `json:"is_read"`
	Creation     time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified     time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (Message) TableName() string {
	return "message"
}
