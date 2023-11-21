package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ReservationMasterRepository interface, which defines methods
type ReservationMasterRepository interface {
	FindReservMasters() ([]models.ReservationMaster, error)
	GetReservMaster(ID int) (models.ReservationMaster, error)
	AddReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error)
	UpdateReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error)
	// DeleteReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryReservation(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "reservation_masters" table in the database and scans the results into a slice of ReservationMasters models.
func (r *repository) FindReservMasters() ([]models.ReservationMaster, error) {
	var reservMaster []models.ReservationMaster
	err := r.db.Preload("ReservationVehicle").Preload("ReservationItem").Preload("ReservationInsurance").Preload("User").Find(&reservMaster).Error // Using Find method

	return reservMaster, err
}

func (r *repository) GetReservMaster(ID int) (models.ReservationMaster, error) {
	var reservMaster models.ReservationMaster
	err := r.db.First(&reservMaster, ID).Error // Using First method

	return reservMaster, err
}

func (r *repository) AddReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error) {
	err := r.db.Create(&reservMaster).Error

	return reservMaster, err
}

func (r *repository) UpdateReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error) {
	err := r.db.Save(&reservMaster).Error

	return reservMaster, err
}

// func (r *repository) DeleteReservMaster(reservMaster models.ReservationMaster) (models.ReservationMaster, error) {
// 	err := r.db.Delete(&reservMaster).Error // Using Delete method

// 	return reservMaster, err
// }
