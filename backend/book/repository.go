package book

import (
	"github.com/marwanjuna/latihan-assessment/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]entity.Book, error)
	PostBook(book entity.Book) (entity.Book, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllBooks() ([]entity.Book, error) {
	var Books []entity.Book

	err := r.db.Find(&Books).Error

	if err != nil {
		return Books, err
	}

	return Books, nil
}

func (r *Repository) PostBook(book entity.Book) (entity.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
