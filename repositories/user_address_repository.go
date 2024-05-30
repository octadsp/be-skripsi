package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserAddressRepository interface, which defines methods
type UserAddressRepository interface {
	CreateUserAddress(user models.UserAddress) (models.UserAddress, error)
	GetProvinces() ([]models.MasterProvince, error)
	GetProvinceByID(province_id string) (models.MasterProvince, error)
	GetRegencies() ([]models.MasterRegency, error)
	GetRegenciesByProvinceID(province_id string) ([]models.MasterRegency, error)
	GetRegencyByID(regency_id string) (models.MasterRegency, error)
	GetDistricts() ([]models.MasterDistrict, error)
	GetDistrictsByRegencyID(regency_id string) ([]models.MasterDistrict, error)
	GetDistrictByID(district_id string) (models.MasterDistrict, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUserAddress(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) CreateUserAddress(UserAddress models.UserAddress) (models.UserAddress, error) {
	err := r.db.Create(&UserAddress).Error // Using Create method
	return UserAddress, err
}

// Provinces
func (r *repository) GetProvinces() ([]models.MasterProvince, error) {
	var provinces []models.MasterProvince
	err := r.db.Find(&provinces).Error
	return provinces, err
}

func (r *repository) GetProvinceByID(province_id string) (models.MasterProvince, error) {
	var province models.MasterProvince
	err := r.db.First(&province, "id = ?", province_id).Error
	return province, err
}

// Regencies
func (r *repository) GetRegencies() ([]models.MasterRegency, error) {
	var regencies []models.MasterRegency
	err := r.db.Preload("Province").Find(&regencies).Error
	return regencies, err
}

func (r *repository) GetRegencyByID(regency_id string) (models.MasterRegency, error) {
	var regency models.MasterRegency
	err := r.db.Preload("Province").First(&regency, "id = ?", regency_id).Error
	return regency, err
}

func (r *repository) GetRegenciesByProvinceID(province_id string) ([]models.MasterRegency, error) {
	var regencies []models.MasterRegency
	err := r.db.Preload("Province").Where("province_id = ?", province_id).Find(&regencies).Error
	return regencies, err
}

// Districts
func (r *repository) GetDistricts() ([]models.MasterDistrict, error) {
	var districts []models.MasterDistrict
	err := r.db.Preload("Regency").Preload("Regency.Province").Find(&districts).Error
	return districts, err
}

func (r *repository) GetDistrictByID(district_id string) (models.MasterDistrict, error) {
	var district models.MasterDistrict
	err := r.db.Preload("Regency").Preload("Regency.Province").First(&district, "id = ?", district_id).Error
	return district, err
}

func (r *repository) GetDistrictsByRegencyID(regency_id string) ([]models.MasterDistrict, error) {
	var districts []models.MasterDistrict
	err := r.db.Preload("Regency").Preload("Regency.Province").Where("regency_id = ?", regency_id).Find(&districts).Error
	return districts, err
}
