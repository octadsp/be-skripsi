package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/pg"
	"fmt"
)

func RunMigration() {
	err := pg.DB.AutoMigrate(
		&models.User{},
		&models.Roles{},
		&models.CarBrand{},
		&models.CarType{},
		&models.FamilyRelation{},
		&models.SimClass{},
		&models.CompanyPartner{},
		&models.CompanyService{},
		&models.DemageCategory{},
		&models.DemageSubCategory{},
		&models.StatusReservation{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
