package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/pg"
	"fmt"
)

func RunMigration() {
	err := pg.DB.AutoMigrate(
		&models.CarBrand{},
		&models.CarClass{},
		&models.CarType{},
		&models.CompanyPartner{},
		&models.CompanyService{},
		&models.DemageCategory{},
		&models.DemageSubCategory{},
		&models.FamilyRelation{},
		&models.PriceList{},
		&models.Roles{},
		&models.SimClass{},
		&models.StatusReservation{},
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
