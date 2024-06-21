package database

import (
	"be-skripsi/models"
	"be-skripsi/pkg/bcrypt"
	"be-skripsi/pkg/pg"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
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
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		var DEFAULT_SUPERADMIN_EMAIL = os.Getenv("DEFAULT_SUPERADMIN_EMAIL")
		var DEFAULT_SUPERADMIN_PASSWORD = os.Getenv("DEFAULT_SUPERADMIN_PASSWORD")
		var DEFAULT_SUPERADMIN_FULLNAME = os.Getenv("DEFAULT_SUPERADMIN_FULLNAME")

		if pg.DB.Migrator().HasTable(&models.User{}) {
			if err := pg.DB.Where("email = ?", DEFAULT_SUPERADMIN_EMAIL).First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				// Insert seed data
				// OK Hashing Password
				password, _ := bcrypt.HashPassword(DEFAULT_SUPERADMIN_PASSWORD)

				user := &models.User{
					ID:       uuid.New().String()[:8],
					Email:    DEFAULT_SUPERADMIN_EMAIL,
					Password: password,
					Role:     "ADMIN",
				}
				err = pg.DB.Create(&user).Error
				if err != nil {
					fmt.Println(err)
					panic("Failed to seed superadmin default account")
				}

				// OK Compose payload for User Detail
				userDetail := &models.UserDetail{
					ID:       uuid.New().String()[:8],
					UserID:   user.ID,
					FullName: DEFAULT_SUPERADMIN_FULLNAME,
				}
				err = pg.DB.Create(&userDetail).Error
				if err != nil {
					fmt.Println(err)
					panic("Failed to seed superadmin default account details")
				}
			}
		}
	}

	fmt.Println("Migration Success")
}
