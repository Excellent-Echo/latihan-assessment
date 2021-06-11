package user

import (
	"gorm.io/gorm"
	"latihan-assessment/entity"
)

type Repository interface {
	GetAll()([]entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository  {
	return &repository{db}
}

func (r *repository)GetAll()([]entity.User, error)  {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, nil
	}

	return users, nil
}