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
	GetOrdersByUserID(userID string, orderStatus string) ([]models.Order, error)
	GetOrdersAdmin(orderStatus string) ([]models.Order, error)
	UpdateOrderByID(id string, order models.Order) (models.Order, error)
	GetOrderItemsByOrderID(orderID string) ([]models.OrderItem, error)
	CreateOrderPayment(orderPayment models.OrderPayment) (models.OrderPayment, error)
	GetOrderPaymentByOrderID(orderID string) (models.OrderPayment, error)
	GetOrderPaymentByID(orderPaymentID string) (models.OrderPayment, error)
	GetAllOrderPayments(orderPaymentStatus string) ([]models.OrderPayment, error)
	UpdateOrderPaymentByID(orderPaymentID string, orderPayment models.OrderPayment) (models.OrderPayment, error)
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

func (r *repository) GetOrdersByUserID(userID string, orderStatus string) ([]models.Order, error) {
	orderStatus = "%" + orderStatus + "%"

	var order []models.Order
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
		Where("status like ?", orderStatus).
		Find(&order, "user_id = ?", userID).Error
	return order, err
}

func (r *repository) GetOrdersAdmin(orderStatus string) ([]models.Order, error) {
	orderStatus = "%" + orderStatus + "%"

	var order []models.Order
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
		Where("status like ?", orderStatus).
		Find(&order).Error
	return order, err
}

func (r *repository) UpdateOrderByID(id string, order models.Order) (models.Order, error) {
	err := r.db.Model(&order).Where("id = ?", id).Updates(&order).Error
	return order, err
}

func (r *repository) GetOrderItemsByOrderID(orderID string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Brand").
		Find(&orderItems, "order_id = ?", orderID).Error
	return orderItems, err
}

func (r *repository) CreateOrderPayment(orderPayment models.OrderPayment) (models.OrderPayment, error) {
	err := r.db.Create(&orderPayment).Error
	return orderPayment, err
}

func (r *repository) GetOrderPaymentByOrderID(orderID string) (models.OrderPayment, error) {
	var orderPayment models.OrderPayment
	err := r.db.
		Preload("Order").
		Preload("Order.UserAddress").
		Preload("Order.UserAddress.Province").
		Preload("Order.UserAddress.Regency").
		Preload("Order.UserAddress.District").
		Preload("Order.DeliveryFare").
		Preload("Order.DeliveryFare.Province").
		Preload("Order.DeliveryFare.Regency").
		Preload("Order.OrderItem").
		Preload("Order.OrderItem.Product").
		Preload("Order.OrderItem.Product.Category").
		Preload("Order.OrderItem.Product.Brand").
		First(&orderPayment, "order_id = ?", orderID).Error
	return orderPayment, err
}

func (r *repository) GetAllOrderPayments(orderPaymentStatus string) ([]models.OrderPayment, error) {
	orderPaymentStatus = "%" + orderPaymentStatus + "%"

	var orderPayments []models.OrderPayment
	err := r.db.
		Preload("Order").
		Preload("Order.UserAddress").
		Preload("Order.UserAddress.Province").
		Preload("Order.UserAddress.Regency").
		Preload("Order.UserAddress.District").
		Preload("Order.DeliveryFare").
		Preload("Order.DeliveryFare.Province").
		Preload("Order.DeliveryFare.Regency").
		Preload("Order.OrderItem").
		Preload("Order.OrderItem.Product").
		Preload("Order.OrderItem.Product.Category").
		Preload("Order.OrderItem.Product.Brand").
		Where("status like ?", orderPaymentStatus).
		Find(&orderPayments).Error
	return orderPayments, err
}

func (r *repository) GetOrderPaymentByID(orderPaymentID string) (models.OrderPayment, error) {
	var orderPayment models.OrderPayment
	err := r.db.
		Preload("Order").
		Preload("Order.UserAddress").
		Preload("Order.UserAddress.Province").
		Preload("Order.UserAddress.Regency").
		Preload("Order.UserAddress.District").
		Preload("Order.DeliveryFare").
		Preload("Order.DeliveryFare.Province").
		Preload("Order.DeliveryFare.Regency").
		Preload("Order.OrderItem").
		Preload("Order.OrderItem.Product").
		Preload("Order.OrderItem.Product.Category").
		Preload("Order.OrderItem.Product.Brand").
		First(&orderPayment, "id = ?", orderPaymentID).Error
	return orderPayment, err
}

func (r *repository) UpdateOrderPaymentByID(orderPaymentID string, orderPayment models.OrderPayment) (models.OrderPayment, error) {
	err := r.db.Model(&orderPayment).Where("id = ?", orderPaymentID).Updates(&orderPayment).Error
	return orderPayment, err
}
