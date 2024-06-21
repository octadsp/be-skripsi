package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the CartRepository interface, which defines methods
type CartRepository interface {
	CreateCartItem(cart models.CartItem, basePrice int64, installationFee int64) (models.CartItem, error)
	GetCartItems(userID string) ([]models.CartItem, error)
	GetCartItem(productID string, userID string) (models.CartItem, error)
	GetCartItemByID(id string) (models.CartItem, error)
	AddCartItemQty(
		productID string,
		userID string,
		qty int64,
		basePrice int64,
		installationFee int64,
		withInstallation bool,
	) (models.CartItem, error)
	UpdateCartItem(
		productID string,
		userID string,
		qty int64,
		basePrice int64,
		installationFee int64,
		withInstallation bool,
	) (models.CartItem, error)
	DeleteCartItemByID(id string) (models.CartItem, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateCartItem(
	cart models.CartItem,
	basePrice int64,
	installationFee int64,
) (models.CartItem, error) {
	err := r.db.
		Create(&cart).
		Update("sub_total", gorm.Expr("qty * ?", basePrice)).
		Update("sub_total", gorm.Expr("sub_total + ?", installationFee)).Error // Using Create method
	return cart, err
}

func (r *repository) GetCartItems(userID string) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.
		Preload("Product").
		Preload("Product.Brand").
		Preload("Product.Category").
		Preload("Product.ProductImage").
		Where("user_id = ?", userID).
		Find(&cartItems).Error

	return cartItems, err
}

func (r *repository) GetCartItem(productID string, userID string) (models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.
		Preload("Product").
		Preload("Product.Brand").
		Preload("Product.Category").
		Preload("Product.ProductImage").
		Where("product_id = ?", productID).
		Where("user_id = ?", userID).
		First(&cartItem).Error

	return cartItem, err
}

func (r *repository) GetCartItemByID(id string) (models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.
		Preload("Product").
		Preload("Product.Brand").
		Preload("Product.Category").
		Where("id = ?", id).
		First(&cartItem).Error

	return cartItem, err
}

func (r *repository) AddCartItemQty(
	productID string,
	userID string,
	qty int64,
	basePrice int64,
	installationFee int64,
	withInstallation bool,
) (models.CartItem, error) {
	var cartItem models.CartItem
	if !withInstallation {
		installationFee = 0
	}
	err := r.db.
		Model(&cartItem).
		Where("product_id = ?", productID).
		Where("user_id = ?", userID).
		Update("qty", gorm.Expr("qty + ?", qty)).
		Update("sub_total", gorm.Expr("qty * ?", basePrice)).
		Update("sub_total", gorm.Expr("sub_total + ?", installationFee)).Error

	return cartItem, err
}

func (r *repository) UpdateCartItem(
	productID string,
	userID string,
	qty int64,
	basePrice int64,
	installationFee int64,
	withInstallation bool,
) (models.CartItem, error) {
	var cartItem models.CartItem
	if !withInstallation {
		installationFee = 0
	}
	err := r.db.
		Model(&cartItem).
		Where("product_id = ?", productID).
		Where("user_id = ?", userID).
		Update("with_installation", withInstallation).
		Update("qty", qty).
		Update("sub_total", gorm.Expr("qty * ?", basePrice)).
		Update("sub_total", gorm.Expr("sub_total + ?", installationFee)).Error

	return cartItem, err
}

func (r *repository) DeleteCartItemByID(id string) (models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Where("id = ?", id).
		Delete(&cartItem).Error
	return cartItem, err
}
