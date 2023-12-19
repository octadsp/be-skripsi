package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the DemageCategoryRepository interface, which defines methods
type DemageCategoryRepository interface {
	FindDemageCategories() ([]models.DemageCategory, error)
	GetDemageCategory(ID int) (models.DemageCategory, error)
	AddDemageCategory(demage models.DemageCategory) (models.DemageCategory, error)
	UpdateDemageCategory(demage models.DemageCategory) (models.DemageCategory, error)
	DeleteDemageCategory(demage models.DemageCategory, ID int) (models.DemageCategory, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryDemageCategory(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "DemageCategorys" table in the database and scans the results into a slice of DemageCategorys models.
func (r *repository) FindDemageCategories() ([]models.DemageCategory, error) {
	var demages []models.DemageCategory
	err := r.db.Order("id").Find(&demages).Error // Using Find method

	return demages, err
}

func (r *repository) GetDemageCategory(ID int) (models.DemageCategory, error) {
	var demage models.DemageCategory
	err := r.db.First(&demage, ID).Error // Using First method

	return demage, err
}

func (r *repository) AddDemageCategory(demage models.DemageCategory) (models.DemageCategory, error) {
	err := r.db.Create(&demage).Error

	return demage, err
}

func (r *repository) UpdateDemageCategory(demage models.DemageCategory) (models.DemageCategory, error) {
	err := r.db.Save(&demage).Error

	return demage, err
}

func (r *repository) DeleteDemageCategory(demage models.DemageCategory, ID int) (models.DemageCategory, error) {
	err := r.db.Delete(&demage).Error // Using Delete method

	return demage, err
}
