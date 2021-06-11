package book

import (
	"latihanassesment/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Book, error)
	Create(user entity.Book) (entity.Book, error)
	FindBYID(ID string) (entity.Book, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Book, error)
	DeleteByID(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Book, error) {
	var users []entity.Book

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Create(user entity.Book) (entity.Book, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindBYID(ID string) (entity.Book, error) {
	var user entity.Book

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Book, error) {
	var user entity.Book

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Book{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
