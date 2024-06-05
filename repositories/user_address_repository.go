package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserAddressRepository interface, which defines methods
type UserAddressRepository interface {
	CreateUserAddress(user models.UserAddress) (models.UserAddress, error)
	GetUserAddressByID(user_address_id string) (models.UserAddress, error)
	GetUserAddresses(user_id string) ([]models.UserAddress, error)
	UpdateUserAddressByID(user_address_id string, user_address models.UserAddress) (models.UserAddress, error)
	UpdateUserDefaultAddress(user_address_id string, user_id string) error
	DeleteUserAddressByID(user_address_id string) (models.UserAddress, error)
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

// User Address
func (r *repository) CreateUserAddress(user_address models.UserAddress) (models.UserAddress, error) {
	err := r.db.Create(&user_address).Error // Using Create method
	return user_address, err
}

func (r *repository) GetUserAddressByID(user_address_id string) (models.UserAddress, error) {
	var userAddress models.UserAddress
	err := r.db.Preload("User").Preload("Province").Preload("Regency").Preload("District").First(&userAddress, "id = ?", user_address_id).Error
	return userAddress, err
}

func (r *repository) GetUserAddresses(user_id string) ([]models.UserAddress, error) {
	var userAddresses []models.UserAddress
	err := r.db.Preload("User").Preload("Province").Preload("Regency").Preload("District").Where("user_id = ?", user_id).Find(&userAddresses).Error
	return userAddresses, err
}

func (r *repository) UpdateUserAddressByID(user_address_id string, user_address models.UserAddress) (models.UserAddress, error) {
	err := r.db.Model(&user_address).Where("id = ?", user_address_id).Updates(&user_address).Error
	return user_address, err
}

func (r *repository) UpdateUserDefaultAddress(user_address_id string, user_id string) error {
	err := r.db.Model(&models.UserAddress{}).Where("user_id = ?", user_id).Where("id != ?", user_address_id).Updates(map[string]interface{}{"default_address": false}).Error
	return err
}

func (r *repository) DeleteUserAddressByID(user_address_id string) (models.UserAddress, error) {
	var userAddress models.UserAddress
	err := r.db.Where("id = ?", user_address_id).Delete(&userAddress).Error
	return userAddress, err
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
