package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the ReservationInsuranceRepository interface, which defines methods
type ReservationInsuranceRepository interface {
	FindReservInsurances() ([]models.ReservationInsurance, error)
	GetReservInsurance(ID int) (models.ReservationInsurance, error)
	AddReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error)
	UpdateReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error)
	// DeleteReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryReservationInsurance(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "reservation_Insurances" table in the database and scans the results into a slice of ReservationInsurances models.
func (r *repository) FindReservInsurances() ([]models.ReservationInsurance, error) {
	var reservInsurance []models.ReservationInsurance
	err := r.db.Find(&reservInsurance).Error // Using Find method

	return reservInsurance, err
}

func (r *repository) GetReservInsurance(ID int) (models.ReservationInsurance, error) {
	var reservInsurance models.ReservationInsurance
	err := r.db.First(&reservInsurance, ID).Error // Using First method

	return reservInsurance, err
}

func (r *repository) AddReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error) {
	err := r.db.Create(&reservInsurance).Error

	return reservInsurance, err
}

func (r *repository) UpdateReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error) {
	err := r.db.Save(&reservInsurance).Error

	return reservInsurance, err
}

// func (r *repository) DeleteReservInsurance(reservInsurance models.ReservationInsurance) (models.ReservationInsurance, error) {
// 	err := r.db.Delete(&reservInsurance).Error // Using Delete method

// 	return reservInsurance, err
// }
