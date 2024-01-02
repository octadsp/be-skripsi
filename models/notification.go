package models

import "time"

type Notification struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	UserID    uint32    `json:"user_id"`
	AuthorBy  string    `json:"author_by" gorm:"type: varchar(100)"`
	Title     string    `json:"title" gorm:"type: varchar(100)"`
	Message   string    `json:"message" gorm:"type: varchar(150)"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (Notification) TableName() string {
	return "notifications"
}
