package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
}

// pointer repository yang digunakan untuk mengakses database dan melakukan operasi CRUD
func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db} // melakukan pointer ke struct repository baru yang diinisialisasi dengan koneksi database yang dibuat
}

// sebuah method register yang dimiliki oleh sebuah struct repository
// method ini menerima sebuah object models.User dan mengembalikan kembali models.User
func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// sebuah method login  yang dimiliki oleh sebuah struct repository
// method yang menerima sebuah string email dan akan mengembalikan sebuah object models.User dan error sebagai hasil operasi
func (r *repository) Login(email string) (models.User, error) {

	// Mendeklarasikan user sebagai object models.User yang nantinya akan berisi data pengguna dari database
	var user models.User
	// menjalankan query First untuk mengambil data pertama yang ditemukan di database, berarti kriteria email yang sama dengan email
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
