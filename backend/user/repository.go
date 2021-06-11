package user

import (
	"backend/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {

	if err := r.db.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
