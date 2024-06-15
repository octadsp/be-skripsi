package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the DeliveryFareRepository interface, which defines methods
type DeliveryFareRepository interface {
	AddDeliveryFare(deliveryFare models.DeliveryFare) (models.DeliveryFare, error)
	GetDeliveryFare(provinceID string, regencyID string) (models.DeliveryFare, error)
	GetDeliveryFareByID(id string) (models.DeliveryFare, error)
	GetDeliveryFares() ([]models.DeliveryFare, error)
	UpdateDeliveryFare(deliveryFare models.DeliveryFare) (models.DeliveryFare, error)
}

func RepositoryDeliveryFare(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func (r *repository) AddDeliveryFare(deliveryFare models.DeliveryFare) (models.DeliveryFare, error) {
	err := r.db.Create(&deliveryFare).Error
	return deliveryFare, err
}

func (r *repository) GetDeliveryFare(provinceID string, regencyID string) (models.DeliveryFare, error) {
	var deliveryFare models.DeliveryFare

	err := r.db.Model(&deliveryFare).Preload("Province").Preload("Regency").Preload("Regency.Province").Where("province_id = ?", provinceID).Where("regency_id = ?", regencyID).First(&deliveryFare).Error
	return deliveryFare, err
}

func (r *repository) GetDeliveryFareByID(id string) (models.DeliveryFare, error) {
	var deliveryFare models.DeliveryFare
	err := r.db.Preload("Province").Preload("Regency").Preload("Regency.Province").First(&deliveryFare, "id = ?", id).Error
	return deliveryFare, err
}

func (r *repository) GetDeliveryFares() ([]models.DeliveryFare, error) {
	var deliveryFares []models.DeliveryFare
	err := r.db.Preload("Province").Preload("Regency").Preload("Regency.Province").Find(&deliveryFares).Error
	return deliveryFares, err
}

func (r *repository) UpdateDeliveryFare(deliveryFare models.DeliveryFare) (models.DeliveryFare, error) {
	err := r.db.Model(&deliveryFare).Where("id = ?", deliveryFare.ID).Updates(&deliveryFare).Error
	return deliveryFare, err
}
