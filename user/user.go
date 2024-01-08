package user

import (
	"YTStreamGoApi/models"
	"github.com/google/uuid"
)

type Store interface {
	GetById(uuid uuid.UUID) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}
