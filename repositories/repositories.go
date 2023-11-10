package repositories

import "gorm.io/gorm"

// repository struct, which implements the interface containing the methods.
type repository struct {
	db *gorm.DB // pointer to a GORM database connection
}
