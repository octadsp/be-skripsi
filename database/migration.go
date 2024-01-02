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
		&models.ReservationMaster{},
		&models.ReservationVehicle{},
		&models.ReservationInsurance{},
		&models.ReservationItem{},
		&models.Roles{},
		&models.SimClass{},
		&models.User{},
		&models.Notification{},
		&models.Reservation{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
