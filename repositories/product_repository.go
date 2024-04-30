package repositories

import (
	"gorm.io/gorm"
)

// declaration of the ProductRepository interface, which defines methods
type ProductRepository interface {
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// func (r *repository) CreateUser(user models.User) (models.User, error) {
// 	err := r.db.Create(&user).Error // Using Create method
// 	return user, err
// }

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