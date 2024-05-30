package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserDetailRepository interface, which defines methods
type UserDetailRepository interface {
	CreateUserDetail(user models.UserDetail) (models.UserDetail, error)
	GetUserDetail(ID string) (models.UserDetail, error)
	UpdateUserDetail(userID string, userDetail models.UserDetail) (models.UserDetail, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUserDetail(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateUserDetail(userDetail models.UserDetail) (models.UserDetail, error) {
	err := r.db.Create(&userDetail).Error // Using Create method
	return userDetail, err
}

func (r *repository) GetUserDetail(userID string) (models.UserDetail, error) {
	var userDetail models.UserDetail
	err := r.db.Where("user_id = ?", userID).First(&userDetail).Error // Using First method

	return userDetail, err
}

func (r *repository) UpdateUserDetail(userID string, userDetail models.UserDetail) (models.UserDetail, error) {
	err := r.db.Model(&userDetail).Where("user_id = ?", userID).Updates(&userDetail).Error
	return userDetail, err
}
