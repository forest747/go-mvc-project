package dblayer

import (
	"errors"
	"github.com/forest747/go-mvc-project/backend/src/models"
)

type DBLayer interface {
	GetAllAnimals() ([]models.Animal, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
