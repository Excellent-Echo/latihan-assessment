package user

import (
	"books-project/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	ReadAllUser() ([]entity.User, error)
	ReadUserByID(id string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) ReadAllUser() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) ReadUserByID(id string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
