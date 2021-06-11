package user

import (
	"book-list/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	FindUserByEmail(email string) (entity.Users, error)
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

func (r *repository) CreateUser(user entity.Users) (entity.Users, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindUserByEmail(email string) (entity.Users, error) {
	var users entity.Users

	if err := r.db.Where("email = ?", email).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
