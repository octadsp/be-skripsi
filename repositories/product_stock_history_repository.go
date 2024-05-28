package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ProductStockHistoryRepository interface, which defines methods
type ProductStockHistoryRepository interface {
	InsertProductStockHistory(productStockHistory models.ProductStockHistory) (models.ProductStockHistory, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryProductStockHistory(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) InsertProductStockHistory(productStockHistory models.ProductStockHistory) (models.ProductStockHistory, error) {
	err := r.db.Save(&productStockHistory).Error
	return productStockHistory, err
}
