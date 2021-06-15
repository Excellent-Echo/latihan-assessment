package book

import (
	"book-list/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllBook() ([]entity.Books, error)
	CreateBook(book entity.Books) (entity.Books, error)
	FindBookByID(ID string) (entity.Books, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Books, error)
	DeleteByID(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllBook() ([]entity.Books, error) {
	var books []entity.Books

	if err := r.db.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (r *repository) CreateBook(book entity.Books) (entity.Books, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) FindBookByID(ID string) (entity.Books, error) {
	var book entity.Books

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Books, error) {
	var book entity.Books

	if err := r.db.Model(&book).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return book, err
	}

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Books{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
