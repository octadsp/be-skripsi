package models

import "time"

type User struct {
	ID        uint32 `json:"id" gorm:"primary_key:auto_increment"`
	FullName  string `json:"fullname" gorm:"type: varchar(100)"`
	LastName  string `json:"lastname" gorm:"type: varchar(100)"`
	Institute string `json:"institute" gorm:"type: varchar(50)"`
	Email     string `json:"email" gorm:"type: varchar(100); unique "`
	Password  string `json:"password" gorm:"type: varchar(100)"`
	Phone     string `json:"phone" gorm:"type: varchar(50)"`
	Address   string `json:"address" gorm:"type: varchar(100)"`
	Avatar    string `json:"image"`
	Status    string `json:"status" gorm:"type: varchar(50)"`
	Roles     string `json:"roles" gorm:"type: varchar(50)"`
	// ReservationMasterID uint32            `json:"reservation_master_id"`
	// ReservationMaster   ReservationMaster `json:"reservation_master" gorm:"foreignKey:ReservationMasterID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
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
