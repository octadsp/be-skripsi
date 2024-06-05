package repositories

import (
	"be-skripsi/models"
	"fmt"

	"gorm.io/gorm"
)

// declaration of the ProductRepository interface, which defines methods
type ProductRepository interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProduct(id string) (models.Product, error)
	UpdateProduct(id string, product models.Product) (models.Product, error)
	DeleteProduct(id string) (models.Product, error)
	UpdateProductStock(id string, operator string, quantity int64, product models.Product) (models.Product, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error // Using Create method
	return product, err
}

func (r *repository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Brand").Preload("Category").Preload("ProductImage").Find(&products).Error
	return products, err
}

func (r *repository) GetProduct(id string) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Brand").Preload("Category").Preload("ProductImage").First(&product, "id = ?", id).Error
	return product, err
}

func (r *repository) UpdateProduct(id string, product models.Product) (models.Product, error) {
	err := r.db.Model(&product).Where("id = ?", id).Updates(&product).Error
	return product, err
}

func (r *repository) UpdateProductStock(id string, operator string, quantity int64, product models.Product) (models.Product, error) {
	err := r.db.Model(&product).Where("id = ?", id).UpdateColumn("stock", gorm.Expr(fmt.Sprintf("stock %s ?", operator), quantity)).Error
	return product, err
}

func (r *repository) DeleteProduct(id string) (models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	return product, err
}
