package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the BrandRepository interface, which defines methods
type BrandRepository interface {
	CreateBrand(brand models.Brand) (models.Brand, error)
	GetBrands() ([]models.Brand, error)
	GetBrand(id string) (models.Brand, error)
	UpdateBrand(id string, brand models.Brand) (models.Brand, error)
	DeleteBrand(id string) (models.Brand, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryBrand(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateBrand(brand models.Brand) (models.Brand, error) {
	err := r.db.Create(&brand).Error // Using Create method
	return brand, err
}

func (r *repository) GetBrands() ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.Find(&brands).Error
	return brands, err
}

func (r *repository) GetBrand(id string) (models.Brand, error) {
	var brand models.Brand
	err := r.db.First(&brand, "id = ?", id).Error
	return brand, err
}

func (r *repository) UpdateBrand(id string, brand models.Brand) (models.Brand, error) {
	err := r.db.Model(&brand).Where("id = ?", id).Updates(&brand).Error
	return brand, err
}

func (r *repository) DeleteBrand(id string) (models.Brand, error) {
	var brand models.Brand
	err := r.db.Where("id = ?", id).Delete(&brand).Error
	return brand, err
}

// func (r *repository) GetUserByEmail(email string) (models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, "email = ?", email).Error
// 	return user, err
// }

// func (r *repository) GetUserByID(ID string) (models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, "id = ?", ID).Error

// 	return user, err
// }

// func (r *repository) GetUsers() ([]models.User, error) {
// 	var users []models.User
// 	err := r.db.Find(&users).Error
// 	return users, err
// }

// func (r *repository) UpdateUserByEmail(email string, user models.User) (models.User, error) {
// 	err := r.db.Model(&user).Where("email = ?", email).Updates(&user).Error
// 	return user, err
// }

// func (r *repository) DeleteUserByEmail(email string) (models.User, error) {
// 	var user models.User
// 	err := r.db.Where("email = ?", email).Delete(&user).Error
// 	return user, err
// }
