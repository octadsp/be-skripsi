package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ReservationInsuranceRepository interface, which defines methods
type ReservationRepository interface {
	FindReservations() ([]models.Reservation, error)
	GetReservation(ID int) (models.Reservation, error)
	AddReservation(reserv models.Reservation) (models.Reservation, error)
	UpdateReservation(reserv models.Reservation) (models.Reservation, error)
	DeleteReservation(reserv models.Reservation) (models.Reservation, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryReservation(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "reservation_s" table in the database and scans the results into a slice of Reservations models.
func (r *repository) FindReservations() ([]models.Reservation, error) {
	var reserv []models.Reservation
	err := r.db.Preload("ReservationItem").Preload("User").Order("order_masuk").Find(&reserv).Error // Using Find method

	return reserv, err
}

func (r *repository) GetReservation(ID int) (models.Reservation, error) {
	var reserv models.Reservation
	err := r.db.Preload("ReservationItem").Preload("User").First(&reserv, ID).Error // Using First method

	return reserv, err
}

func (r *repository) AddReservation(reserv models.Reservation) (models.Reservation, error) {
	err := r.db.Create(&reserv).Error

	return reserv, err
}

func (r *repository) UpdateReservation(reserv models.Reservation) (models.Reservation, error) {
	err := r.db.Save(&reserv).Error

	return reserv, err
}

func (r *repository) DeleteReservation(reserv models.Reservation) (models.Reservation, error) {
	err := r.db.Delete(&reserv).Error // Using Delete method

	return reserv, err
}
