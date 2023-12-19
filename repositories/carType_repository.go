package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CarBrandRepository interface, which defines methods
type CarTypeRepository interface {
	FindCarTypes(offer, limit int) ([]models.CarType, error)
	FindAllCarTypes() ([]models.CarType, error)
	GetCarType(ID int) (models.CarType, error)
	AddCarType(types models.CarType) (models.CarType, error)
	UpdateCarType(types models.CarType) (models.CarType, error)
	DeleteCarType(types models.CarType, ID int) (models.CarType, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryCarType(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "cartypes" table in the database and scans the results into a slice of CarTypes models.
func (r *repository) FindCarTypes(offset, limit int) ([]models.CarType, error) {
	var types []models.CarType
	err := r.db.Offset(offset).Limit(limit).Order("id").Find(&types).Error // Using Find method

	return types, err
}

func (r *repository) FindAllCarTypes() ([]models.CarType, error) {
	var types []models.CarType
	err := r.db.Order("id").Find(&types).Error // Using Find method

	return types, err
}

func (r *repository) GetCarType(ID int) (models.CarType, error) {
	var types models.CarType
	err := r.db.First(&types, ID).Error // Using First method

	return types, err
}

func (r *repository) AddCarType(types models.CarType) (models.CarType, error) {
	err := r.db.Create(&types).Error

	return types, err
}

func (r *repository) UpdateCarType(types models.CarType) (models.CarType, error) {
	err := r.db.Save(&types).Error

	return types, err
}

func (r *repository) DeleteCarType(types models.CarType, ID int) (models.CarType, error) {
	err := r.db.Delete(&types).Error // Using Delete method

	return types, err
}
