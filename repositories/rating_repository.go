package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CarBrandRepository interface, which defines methods
type RatingRepository interface {
	FindRatingByUser(userID int) ([]models.Rating, error)
	AddRating(rating models.Rating) (models.Rating, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryRating(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "carbrands" table in the database and scans the results into a slice of CarBrands models.
func (r *repository) FindRatingByUser(userID int) ([]models.Rating, error) {
	var ratings []models.Rating
	err := r.db.Preload("Reservation").Preload("User").Where("user_id = ?", userID).Find(&ratings).Error // Using Find method

	return ratings, err
}

func (r *repository) AddRating(rating models.Rating) (models.Rating, error) {
	err := r.db.Create(&rating).Error

	return rating, err
}
