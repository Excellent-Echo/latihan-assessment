package user

import (
	"github.com/marwanjuna/latihan-assessment/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllUsers() ([]entity.User, error) {
	var Users []entity.User

	err := r.db.Find(&Users).Error

	if err != nil {
		return Users, err
	}

	return Users, nil
}
