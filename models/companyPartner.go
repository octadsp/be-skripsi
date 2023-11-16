package models

import "time"

type CompanyPartner struct {
	ID        uint32    `json:"id" gorm:"primary_key:auto_increment"`
	Kode      string    `json:"kode" gorm:"type: varchar(20)"`
	Tipe      string    `json:"tipe" gorm:"type: varchar(100)"`
	Name      string    `json:"name" gorm:"type: varchar(100)"`
	Image     string    `json:"image"`
	Status    string    `json:"status" gorm:"type: varchar(10)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (CompanyPartner) TableName() string {
	return "company_partners"
}
