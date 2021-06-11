package book

import (
	"gorm.io/gorm"
	"latihan-assessment/entity"
)

type Repository interface {
	GetAllBooks() ([]entity.Books, error)
	CreateBooks(book entity.Books) (entity.Books, error)
	GetBooksById(ID string) (entity.Books, error)
	DeleteBookById(ID string) (string, error)
	UpdateBookById(ID string, bookUpdate map[string]interface{}) (entity.Books, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllBooks() ([]entity.Books, error) {
	var books []entity.Books

	if err := r.db.Find(&books).Error; err != nil {
		return books, nil
	}

	return books, nil
}

func (r *repository) CreateBooks(book entity.Books) (entity.Books, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, nil
	}

	return book, nil
}

func (r *repository) GetBooksById(ID string) (entity.Books, error) {
	var book entity.Books

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, nil
	}

	return book, nil
}

func (r *repository) DeleteBookById(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Books{}).Error; err != nil {
		return "error", nil
	}

	return "success", nil
}

func (r *repository) UpdateBookById(ID string, bookUpdate map[string]interface{}) (entity.Books, error) {
	var bookWillUpdate entity.Books

	if err := r.db.Model(&bookWillUpdate).Where("id = ?", ID).Updates(bookUpdate); err != nil {
		return bookWillUpdate, nil
	}

	if err := r.db.Where("id = ?", ID).Find(bookWillUpdate).Error; err != nil {
		return bookWillUpdate, nil
	}

	return bookWillUpdate, nil

}
