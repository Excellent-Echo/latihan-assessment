package user

import (
	"book-list/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Users, error) {
	var users []entity.Users

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
