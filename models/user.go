package models

import "time"

type User struct {
	ID       string    `json:"id" gorm:"primary_key;type:varchar(140)"`
	Email    string    `json:"email" gorm:"type: varchar(140); unique "`
	Password string    `json:"password" gorm:"type: varchar(140)"`
	Role     string    `json:"role" gorm:"type: varchar(140)"`
	Creation time.Time `json:"creation" gorm:"autoCreateTime"`
	Modified time.Time `json:"modified" gorm:"autoCreateTime"`
}

// role = ADMIN/CUSTOMER

func (User) TableName() string {
	return "user"
}
