package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CategoryRepository interface, which defines methods
type CategoryRepository interface {
	CreateCategory(user models.Category) (models.Category, error)
	GetCategories() ([]models.Category, error)
	GetCategory(id string) (models.Category, error)
	UpdateCategory(id string, category models.Category) (models.Category, error)
	DeleteCategory(id string) (models.Category, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateCategory(user models.Category) (models.Category, error) {
	err := r.db.Create(&user).Error // Using Create method
	return user, err
}

func (r *repository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *repository) GetCategory(id string) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, "id = ?", id).Error
	return category, err
}

func (r *repository) UpdateCategory(id string, category models.Category) (models.Category, error) {
	err := r.db.Model(&category).Where("id = ?", id).Updates(&category).Error
	return category, err
}

func (r *repository) DeleteCategory(id string) (models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).Delete(&category).Error
	return category, err
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