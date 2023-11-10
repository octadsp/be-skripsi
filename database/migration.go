package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/pg"
	"fmt"
)

func RunMigration() {
	err := pg.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
