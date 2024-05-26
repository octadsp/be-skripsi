package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ProductStockHistoryRepository interface, which defines methods
type ProductStockHistoryRepository interface {
	CreateProductImage(productImage models.ProductImage) (models.ProductImage, error)
	DeleteProductImage(id string) (models.ProductImage, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryProductStockHistory(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) UpdateProductStock(productStockHistory models.ProductStockHistory) (models.ProductStockHistory, error) {
	err := r.db.Save(&productStockHistory).Error
	return productStockHistory, err
}

// func (r *repository) CreateProductImage(productImage models.ProductImage) (models.ProductImage, error) {
// 	err := r.db.Create(&productImage).Error
// 	return productImage, err
// }

// func (r *repository) DeleteProductImage(id string) (models.ProductImage, error) {
// 	var productImage models.ProductImage
// 	err := r.db.Where("id = ?", id).Delete(&productImage).Error
// 	return productImage, err
// }
