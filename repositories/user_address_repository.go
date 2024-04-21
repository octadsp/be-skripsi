package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserAddressRepository interface, which defines methods
type UserAddressRepository interface {
	CreateUserAddress(user models.UserAddress) (models.UserAddress, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUserAddress(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateUserAddress(UserAddress models.UserAddress) (models.UserAddress, error) {
	err := r.db.Create(&UserAddress).Error // Using Create method
	return UserAddress, err
}
