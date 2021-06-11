package book

import (
	"book-list/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllBook() ([]entity.Books, error)
	CreateBook(book entity.Books) (entity.Books, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Books, error) {
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
