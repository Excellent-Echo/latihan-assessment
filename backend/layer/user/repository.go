package user

import (
	"book-list/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	FindUserByEmail(email string) (entity.Users, error)
	FindUserByID(ID string) (entity.Users, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Users, error)
	DeleteByID(ID string) (string, error)
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

func (r *repository) FindUserByID(ID string) (entity.Users, error) {
	var user entity.Users

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Users, error) {
	var user entity.Users

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Users{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
