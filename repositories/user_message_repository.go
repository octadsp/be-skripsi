package repositories

import (
	userdto "be-skripsi/dto/user"
	"be-skripsi/models"

	"gorm.io/gorm"
)

// declaration of the UserMessageRepository interface, which defines methods
type UserMessageRepository interface {
	GetChats(userRole string, userId string) ([]userdto.UserInboxResponse, error)
	GetChatLogs(userId string, otherUserId string) ([]userdto.UserChatLogResponse, error)
	SendMessage(userRole string, userId string, otherUserId string, message models.Message) error
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUserMessage(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

func (r *repository) GetChats(userRole string, userId string) ([]userdto.UserInboxResponse, error) {
	// OK query variables
	table := "message"
	condition := ""
	group := "admin, customer, message, creation"
	order := ""

	// OK conditional field based on userRole
	collocutorRole := ""
	switch userRole {
	case "ADMIN":
		collocutorRole = "customer"
		condition = "admin = ?"
		order = "customer, creation desc"
	case "CUSTOMER":
		collocutorRole = "admin"
		condition = "customer = ?"
		order = "admin, creation desc"
	}

	var message []userdto.UserInboxResponse
	// Define the subquery SQL string
	subQuery := "(SELECT COUNT(*) FROM " + table + " WHERE admin = mainQuery.admin AND customer = mainQuery.customer AND sender != ? AND is_read = false) AS total_unread"

	// Build the main query
	err := r.db.
		Table(table+" AS mainQuery").
		Select("DISTINCT ON "+"("+collocutorRole+") "+collocutorRole+", "+"message as last_message, "+subQuery, userRole).
		Where(condition, userId).
		Group(group).
		Order(order).
		Find(&message).
		Error

	return message, err
}

func (r *repository) SendMessage(userRole string, userId string, otherUserId string, message models.Message) error {
	switch userRole {
	case "ADMIN":
		message.Admin = userId
		message.Customer = otherUserId
	case "CUSTOMER":
		message.Customer = userId
		message.Admin = otherUserId
	}

	message.Sender = userRole
	err := r.db.Create(&message).Error // Using Create method
	return err
}

func (r *repository) GetChatLogs(userId string, otherUserId string) ([]userdto.UserChatLogResponse, error) {
	// OK query variables
	table := "message"
	order := "creation desc"

	var message []userdto.UserChatLogResponse
	err := r.db.
		Table(table).
		Select("*").
		Where(
			r.db.Where("admin = ?", userId).Where("customer = ?", otherUserId),
		).
		Or(
			r.db.Where("admin = ?", otherUserId).Where("customer = ?", userId),
		).
		Order(order).
		Find(&message).
		Error

	return message, err
}

// func (r *repository) CreateUserDetail(userDetail models.UserDetail) (models.UserDetail, error) {
// 	err := r.db.Create(&userDetail).Error // Using Create method
// 	return userDetail, err
// }

// func (r *repository) GetUserDetail(userID string) (models.UserDetail, error) {
// 	var userDetail models.UserDetail
// 	err := r.db.Where("user_id = ?", userID).First(&userDetail).Error // Using First method

// 	return userDetail, err
// }

// func (r *repository) UpdateUserDetail(userID string, userDetail models.UserDetail) (models.UserDetail, error) {
// 	err := r.db.Model(&userDetail).Where("user_id = ?", userID).Updates(&userDetail).Error
// 	return userDetail, err
// }
