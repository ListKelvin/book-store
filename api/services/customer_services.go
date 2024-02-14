package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type CustomerService interface {
	Prepare()
	CreateCustomer(db *gorm.DB) (models.Customer, error)
	GetAllCustomers(db *gorm.DB) ([]models.Customer, error)
	GetCustomerById(Id int64, db *gorm.DB) (models.Customer, error)
	DeleteCustomer(ID int64, db *gorm.DB) (models.Customer, error)
}

type customerService struct {
	CustomerRepository models.CustomerRepository
}

// NewCustomerService -> returns new Customer service
func NewCustomerService(r models.CustomerRepository) customerService {
	return customerService{
		CustomerRepository: r,
	}
}

func (u customerService) Create(db *gorm.DB) (*models.Customer, error) {

	return u.CustomerRepository.CreateCustomer(db)
}

func (u customerService) GetAll(db *gorm.DB) ([]models.Customer, error) {

	return u.CustomerRepository.GetAllCustomers(db)
}

func (u customerService) GetById(Id int64, db *gorm.DB) (*models.Customer, error) {

	return u.CustomerRepository.GetCustomerById(Id, db)
}

func (u customerService) Delete(ID int64, db *gorm.DB) (*models.Customer, error) {

	return u.CustomerRepository.DeleteCustomer(ID, db)
}