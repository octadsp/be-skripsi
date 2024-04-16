package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/pg"
	"fmt"
)

func RunMigration() {
	err := pg.DB.AutoMigrate(
		&models.User{},
		&models.UserDetail{},
		&models.UserAddress{},
		&models.MasterProvince{},
		&models.MasterCity{},
		&models.MasterDistrict{},
		&models.Message{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductStockHistory{},
		&models.Category{},
		&models.Brand{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
