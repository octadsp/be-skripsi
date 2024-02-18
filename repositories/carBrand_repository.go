package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CarBrandRepository interface, which defines methods
type CarBrandRepository interface {
	FindCarBrands(offset, limit int) ([]models.CarBrand, error)
	FindAllBrands() ([]models.CarBrand, error)
	GetCarBrand(ID int) (models.CarBrand, error)
	AddCarBrand(brand models.CarBrand) (models.CarBrand, error)
	UpdateCarBrand(brand models.CarBrand) (models.CarBrand, error)
	DeleteCarBrand(brand models.CarBrand, ID int) (models.CarBrand, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryCarBrand(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "carbrands" table in the database and scans the results into a slice of CarBrands models.
func (r *repository) FindCarBrands(offset, limit int) ([]models.CarBrand, error) {
	var brands []models.CarBrand
	err := r.db.Offset(offset).Limit(limit).Order("name").Find(&brands).Error // Using Find method

	return brands, err
}

func (r *repository) FindAllBrands() ([]models.CarBrand, error) {
	var brands []models.CarBrand
	err := r.db.Order("name").Find(&brands).Error // Using Find method

	return brands, err
}

func (r *repository) GetCarBrand(ID int) (models.CarBrand, error) {
	var brand models.CarBrand
	err := r.db.First(&brand, ID).Error // Using First method

	return brand, err
}

func (r *repository) AddCarBrand(brand models.CarBrand) (models.CarBrand, error) {
	err := r.db.Create(&brand).Error

	return brand, err
}

func (r *repository) UpdateCarBrand(brand models.CarBrand) (models.CarBrand, error) {
	err := r.db.Save(&brand).Error

	return brand, err
}

func (r *repository) DeleteCarBrand(brand models.CarBrand, ID int) (models.CarBrand, error) {
	err := r.db.Delete(&brand).Error // Using Delete method

	return brand, err
}
