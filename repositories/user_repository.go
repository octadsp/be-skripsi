package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserRepository interface, which defines methods
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	UpdateInfoUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "users" table in the database and scans the results into a slice of Users models.
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error // Using Find method

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error // Using First method

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) UpdateInfoUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error // Using Delete method

	return user, err
}
