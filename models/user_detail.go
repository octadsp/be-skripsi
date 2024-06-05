package models

import "time"

type UserDetail struct {
	ID          string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	UserID      string    `json:"user_id" gorm:"type: varchar(140); unique"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	FullName    string    `json:"full_name" gorm:"type: varchar(140)"`
	PhoneNumber string    `json:"phone_number" gorm:"type: varchar(140); unique"`
	Creation    time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified    time.Time `json:"modified" gorm:"autoCreateTime"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}
