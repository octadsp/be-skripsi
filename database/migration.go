package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/bcrypt"
	"be-skripsi/pkg/pg"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunMigration() {
	err := pg.DB.AutoMigrate(
		&models.User{},
		&models.UserDetail{},
		&models.UserAddress{},
		&models.MasterProvince{},
		&models.MasterRegency{},
		&models.MasterDistrict{},
		&models.Message{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductStockHistory{},
		&models.Category{},
		&models.Brand{},
		&models.Order{},
		&models.OrderItem{},
		&models.CartItem{},
		&models.OrderPayment{},
		&models.DeliveryFare{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	} else {
		if pg.DB.Migrator().HasTable(&models.User{}) {
			if err := pg.DB.Where("email = ?", "admin@superadmin.com").First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				//Insert seed data
				// OK Hashing Password
				password, _ := bcrypt.HashPassword("1234")

				user := &models.User{
					ID:       uuid.New().String()[:8],
					Email:    "admin@superadmin.com",
					Password: password,
					Role:     "ADMIN",
				}
				err = pg.DB.Create(&user).Error
				if err != nil {
					fmt.Println(err)
					panic("Failed to seed superadmin default account")
				}
			}
		}
	}

	fmt.Println("Migration Success")
}
