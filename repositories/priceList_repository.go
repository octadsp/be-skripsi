package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the PriceListRepository interface, which defines methods
type PriceListRepository interface {
	FindPriceLists() ([]models.PriceList, error)
	GetPriceList(ID int) (models.PriceList, error)
	AddPriceList(price models.PriceList) (models.PriceList, error)
	UpdatePriceList(price models.PriceList) (models.PriceList, error)
	DeletePriceList(price models.PriceList) (models.PriceList, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryPriceList(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "pricelists" table in the database and scans the results into a slice of PriceLists models.
func (r *repository) FindPriceLists() ([]models.PriceList, error) {
	var prices []models.PriceList
	err := r.db.Order("id").Find(&prices).Error // Using Find method

	return prices, err
}

func (r *repository) GetPriceList(ID int) (models.PriceList, error) {
	var price models.PriceList
	err := r.db.First(&price, ID).Error // Using First method

	return price, err
}

func (r *repository) AddPriceList(price models.PriceList) (models.PriceList, error) {
	err := r.db.Create(&price).Error

	return price, err
}

func (r *repository) UpdatePriceList(price models.PriceList) (models.PriceList, error) {
	err := r.db.Save(&price).Error

	return price, err
}

func (r *repository) DeletePriceList(price models.PriceList) (models.PriceList, error) {
	err := r.db.Delete(&price).Error // Using Delete method

	return price, err
}
