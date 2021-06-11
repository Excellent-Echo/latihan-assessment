package book

import (
	"github.com/marwanjuna/latihan-assessment/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]entity.Book, error)
	PostBook(book entity.Book) (entity.Book, error)
	GetOneBook(id string) (entity.Book, error)
	UpdateBook(id string, dataUpdate map[string]interface{}) (entity.Book, error)
	DeleteBook(id string) (string, error)
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

func (r *Repository) GetOneBook(id string) (entity.Book, error) {
	var book entity.Book

	if err := r.db.Where("id = ?", id).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *Repository) UpdateBook(id string, dataUpdate map[string]interface{}) (entity.Book, error) {
	var book entity.Book

	if err := r.db.Model(&book).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return book, err
	}

	if err := r.db.Where("id = ?", id).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *Repository) DeleteBook(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Book{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
