package user

import (
	"latihan-assessment/backend/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository(db)
}

func (r *repository) GetAll() ([]entity.Users, error) {
	var Users []entity.Users

	err := r.db.Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}
