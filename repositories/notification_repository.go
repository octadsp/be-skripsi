package repositories

import (
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserRepository interface, which defines methods
type NotificationRepository interface {
	// FindNotifications() ([]models.Notification, error)
	GetNotificationsByUserID(userID uint) ([]models.Notification, error)
	GetNotif(ID int) (models.Notification, error)
	CreateNotification(notif models.Notification) (models.Notification, error)
	UpdateNotificationStatus(isRead models.Notification) (models.Notification, error)
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryNotification(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) GetNotificationsByUserID(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Preload("User").Where("user_id = ? AND is_read = ?", userID, false).Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func (r *repository) GetNotif(ID int) (models.Notification, error) {
	var notif models.Notification
	err := r.db.Preload("User").First(&notif, ID).Error

	return notif, err
}

func (r *repository) CreateNotification(notif models.Notification) (models.Notification, error) {
	err := r.db.Create(&notif).Error
	return notif, err
}

func (r *repository) UpdateNotificationStatus(isRead models.Notification) (models.Notification, error) {
	err := r.db.Save(&isRead).Error

	return isRead, err
}
