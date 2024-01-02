package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CarClassRepository interface, which defines methods
type CarClassRepository interface {
	FindCarClass(offset, limit int) ([]models.CarClass, error)
	FindAllCarClass() ([]models.CarClass, error)
	GetCarClass(ID int) (models.CarClass, error)
	AddCarClass(class models.CarClass) (models.CarClass, error)
	UpdateCarClass(class models.CarClass) (models.CarClass, error)
	DeleteCarClass(class models.CarClass, ID int) (models.CarClass, error)

	GetCarClasbyBrand(brandID int) ([]models.CarClass, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryCarClass(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "carclass" table in the database and scans the results into a slice of CarClass models.
func (r *repository) FindCarClass(offset, limit int) ([]models.CarClass, error) {
	var class []models.CarClass
	err := r.db.Offset(offset).Limit(limit).Order("id").Preload("CarBrand").Preload("CarType").Find(&class).Error // Using Find method

	return class, err
}

func (r *repository) FindAllCarClass() ([]models.CarClass, error) {
	var class []models.CarClass
	err := r.db.Order("id").Preload("CarBrand").Preload("CarType").Find(&class).Error // Using Find method

	return class, err
}

func (r *repository) GetCarClass(ID int) (models.CarClass, error) {
	var class models.CarClass
	err := r.db.Preload("CarBrand").Preload("CarType").First(&class, ID).Error // Using First method

	return class, err
}

func (r *repository) GetCarClasbyBrand(brandID int) ([]models.CarClass, error) {
	var class []models.CarClass
	err := r.db.Preload("CarBrand").Preload("CarType").Find(&class, "car_brand_id=?", brandID).Error
	
	return class, err
}

func (r *repository) AddCarClass(class models.CarClass) (models.CarClass, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) UpdateCarClass(class models.CarClass) (models.CarClass, error) {
	err := r.db.Save(&class).Error

	return class, err
}

func (r *repository) DeleteCarClass(class models.CarClass, ID int) (models.CarClass, error) {
	err := r.db.Delete(&class).Error // Using Delete method

	return class, err
}
