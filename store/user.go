package store

import (
	"YTStreamGoApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (store *UserStore) GetById(uuid uuid.UUID) (*models.User, error) {
	var m models.User
	if err := store.db.First(&m, uuid).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (store *UserStore) GetByName(name string) (*models.User, error) {
	var m models.User
	if err := store.db.Where(&models.User{Name: name}).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (store *UserStore) Create(user *models.User) error {
	return store.db.Create(user).Error
}

func (store *UserStore) Update(user *models.User) error {
	return store.db.Model(user).Updates(user).Error
}
