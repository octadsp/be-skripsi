package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the OrderRepository interface, which defines methods
type OrderRepository interface {
	CreateOrder(order models.Order) (models.Order, error)
	CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error)
	GetOrderByID(id string) (models.Order, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *repository) CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	err := r.db.Create(&orderItem).Error
	return orderItem, err
}

func (r *repository) GetOrderByID(id string) (models.Order, error) {
	var order models.Order
	err := r.db.
		Preload("UserAddress").
		Preload("UserAddress.Province").
		Preload("UserAddress.Regency").
		Preload("UserAddress.District").
		Preload("DeliveryFare").
		Preload("DeliveryFare.Province").
		Preload("DeliveryFare.Regency").
		Preload("OrderItem").
		Preload("OrderItem.Product").
		Preload("OrderItem.Product.Category").
		Preload("OrderItem.Product.Brand").
		First(&order, "id = ?", id).Error
	return order, err
}
