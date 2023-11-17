package models

import "time"

type User struct {
	ID                uint32            `json:"id" gorm:"primary_key:auto_increment"`
	FullName          string            `json:"fullname" gorm:"type: varchar(100)"`
	LastName          string            `json:"lastname" gorm:"type: varchar(100)"`
	Email             string            `json:"email" gorm:"type: varchar(100); unique "`
	Password          string            `json:"password" gorm:"type: varchar(100)"`
	Phone             string            `json:"phone" gorm:"type: varchar(50)"`
	Address           string            `json:"address" gorm:"type: varchar(100)"`
	Avatar            string            `json:"image"`
	Status            string            `json:"status" gorm:"type: varchar(50)"`
	Roles             string            `json:"roles" gorm:"type: varchar(50)"`
	ReservationMaster ReservationMaster `json:"reservation_master"`
	CreatedAt         time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time         `json:"updated_at" gorm:"autoCreateTime"`
}

// type UserResponse struct {
// 	ID       int    `json:"id"`
// 	FullName string `json:"fullname"`
// 	LastName string `json:"lastname"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	Phone    string `json:"phone"`
// 	Address  string `json:"address"`
// 	Avatar   string `json:"image"`
// 	Status   string `json:"status"`
// 	Roles    string `json:"roles"`
// }

func (User) TableName() string {
	return "users"
}
