package models

import "time"

type CompanyService struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(100)"`
	Desc      string    `json:"desc" gorm:"type: varchar(200)"`
	Image     string    `json:"image"`
	Status    string    `json:"status" gorm:"type: varchar(10)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (CompanyService) TableName() string {
	return "company_services"
}
