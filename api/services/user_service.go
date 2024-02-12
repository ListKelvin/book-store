package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

// UserService : represent the user's service contract
type UserService interface {
	Prepare()
	CreateUser(db *gorm.DB) (models.User, error)
	GetAllUsers(db *gorm.DB) ([]models.User, error)
	GetUserById(Id int64, db *gorm.DB ) (models.User, error)
	DeleteUser(ID int64, db *gorm.DB ) (models.User, error)
}

type userService struct {
	userRepository models.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(r models.UserRepository) userService {
	return userService{
		userRepository: r,
	}
}


func (u userService) Create(db *gorm.DB) ( *models.User, error) {

	return u.userRepository.CreateUser(db)
}

func (u userService) GetAll(db *gorm.DB) ([]models.User, error) {

	return u.userRepository.GetAllUsers(db)
}

func (u userService) GetById(Id int64, db *gorm.DB) ( *models.User, error) {

	return u.userRepository.GetUserById(Id, db)
}

func (u userService) Delete(ID int64, db *gorm.DB)  (*models.User, error) {

	return u.userRepository.DeleteUser(ID, db)
}