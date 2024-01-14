package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ReservationItemRepository interface, which defines methods
type ReservationItemRepository interface {
	FindReservItems() ([]models.ReservationItem, error)
	GetReservItem(ID int) (models.ReservationItem, error)
	GetReservItemByReservId(reservId int) ([]models.ReservationItem, error)
	AddReservItem(reservItem models.ReservationItem) (models.ReservationItem, error)
	UpdateReservItem(reservItem models.ReservationItem) (models.ReservationItem, error)
	// DeleteReservItem(reservItem models.ReservationItem) (models.ReservationItem, error)

	PostToUser(reservItem models.ReservationItem) (models.ReservationItem, error)
	UpdateStatus(status models.ReservationItem) (models.ReservationItem, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryReservationItem(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "reservation_Items" table in the database and scans the results into a slice of ReservationItems models.
func (r *repository) FindReservItems() ([]models.ReservationItem, error) {
	var reservItem []models.ReservationItem
	err := r.db.Preload("Reservation").Preload("DemageSubCategory").Find(&reservItem).Error // Using Find method

	return reservItem, err
}

func (r *repository) GetReservItem(ID int) (models.ReservationItem, error) {
	var reservItem models.ReservationItem
	err := r.db.Preload("Reservation").Preload("DemageSubCategory").First(&reservItem, ID).Error // Using First method
	return reservItem, err
}

func (r *repository) GetReservItemByReservId(reservId int) ([]models.ReservationItem, error) {
	var reservItem []models.ReservationItem
	err := r.db.Preload("Reservation").Preload("DemageSubCategory").Where("reservation_id = ?", reservId).Find(&reservItem).Error
	return reservItem, err
}

func (r *repository) AddReservItem(reservItem models.ReservationItem) (models.ReservationItem, error) {
	err := r.db.Create(&reservItem).Error

	return reservItem, err
}

func (r *repository) PostToUser(reservItem models.ReservationItem) (models.ReservationItem, error) {
	err := r.db.Save(&reservItem).Error

	return reservItem, err
}

func (r *repository) UpdateReservItem(reservItem models.ReservationItem) (models.ReservationItem, error) {
	err := r.db.Save(&reservItem).Error

	return reservItem, err
}

func (r *repository) UpdateStatus(status models.ReservationItem) (models.ReservationItem, error) {
	err := r.db.Save(&status).Error

	return status, err
}

// func (r *repository) DeleteReservItem(reservItem models.ReservationItem) (models.ReservationItem, error) {
// 	err := r.db.Delete(&reservItem).Error // Using Delete method

// 	return reservItem, err
// }
