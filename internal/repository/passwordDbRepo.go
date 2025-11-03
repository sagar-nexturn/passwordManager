package repository

import (
	"github.com/sagar-nexturn/passwordManager/internal/models"
)

type PasswordDbRepo interface {
	AddPassword(password *models.Password) error
	GetPasswordByID(id string) (*models.Password, error)
	GetAllPasswords() ([]models.Password, error)
	GetPasswordByName(name string) (*models.Password, error)
	UpdatePassword(password *models.Password) error
	DeletePassword(id string) error
}
