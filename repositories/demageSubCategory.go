package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the DemageCategoryRepository interface, which defines methods
type DemageSubCategoryRepository interface {
	FindDemageSubCategories(offset, limit int) ([]models.DemageSubCategory, error)
	FindAllDemageSubCategories() ([]models.DemageSubCategory, error)
	GetDemageSubCategory(ID int) (models.DemageSubCategory, error)
	AddDemageSubCategory(demage models.DemageSubCategory) (models.DemageSubCategory, error)
	UpdateDemageSubCategory(demage models.DemageSubCategory) (models.DemageSubCategory, error)
	DeleteDemageSubCategory(demage models.DemageSubCategory, ID int) (models.DemageSubCategory, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryDemageSubCategory(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "DemageCategorys" table in the database and scans the results into a slice of DemageCategorys models.
func (r *repository) FindDemageSubCategories(offset, limit int) ([]models.DemageSubCategory, error) {
	var demages []models.DemageSubCategory
	err := r.db.Offset(offset).Limit(limit).Preload("DemageCategory").Order("demage_category_id, name asc").Find(&demages).Error // Using Find method

	return demages, err
}

func (r *repository) FindAllDemageSubCategories() ([]models.DemageSubCategory, error) {
	var demages []models.DemageSubCategory
	err := r.db.Preload("DemageCategory").Order("demage_category_id, name asc").Find(&demages).Error // Using Find method

	return demages, err
}

func (r *repository) GetDemageSubCategory(ID int) (models.DemageSubCategory, error) {
	var demage models.DemageSubCategory
	err := r.db.First(&demage, ID).Error // Using First method

	return demage, err
}

func (r *repository) AddDemageSubCategory(demage models.DemageSubCategory) (models.DemageSubCategory, error) {
	err := r.db.Create(&demage).Error

	return demage, err
}

func (r *repository) UpdateDemageSubCategory(demage models.DemageSubCategory) (models.DemageSubCategory, error) {
	err := r.db.Save(&demage).Error

	return demage, err
}

func (r *repository) DeleteDemageSubCategory(demage models.DemageSubCategory, ID int) (models.DemageSubCategory, error) {
	err := r.db.Delete(&demage).Error // Using Delete method

	return demage, err
}
