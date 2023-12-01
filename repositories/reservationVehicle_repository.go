package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ReservationVehicleRepository interface, which defines methods
type ReservationVehicleRepository interface {
	FindReservVehicles() ([]models.ReservationVehicle, error)
	GetReservVehicle(ID int) (models.ReservationVehicle, error)
	AddReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error)
	UpdateReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error)
	// DeleteReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryReservationVehicle(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "reservation_Vehicles" table in the database and scans the results into a slice of ReservationVehicles models.
func (r *repository) FindReservVehicles() ([]models.ReservationVehicle, error) {
	var reservVehicle []models.ReservationVehicle
	err := r.db.Find(&reservVehicle).Error // Using Find method

	return reservVehicle, err
}

func (r *repository) GetReservVehicle(ID int) (models.ReservationVehicle, error) {
	var reservVehicle models.ReservationVehicle
	err := r.db.First(&reservVehicle, ID).Error // Using First method

	return reservVehicle, err
}

func (r *repository) AddReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error) {
	err := r.db.Create(&reservVehicle).Error

	return reservVehicle, err
}

func (r *repository) UpdateReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error) {
	err := r.db.Save(&reservVehicle).Error

	return reservVehicle, err
}

// func (r *repository) DeleteReservVehicle(reservVehicle models.ReservationVehicle) (models.ReservationVehicle, error) {
// 	err := r.db.Delete(&reservVehicle).Error // Using Delete method

// 	return reservVehicle, err
// }
