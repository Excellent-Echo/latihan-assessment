package user

import (
	"backend/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindbyID(ID int) (entity.User, error)
	UpdateByID(ID int, dataUpdate map[string]interface{}) (entity.User, error)
	DeleteByID(ID int) (string, error)
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

func (r *repository) FindbyID(ID int) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ? ", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateByID(ID int, dataUpdate map[string]interface{}) (entity.User, error) {

}
