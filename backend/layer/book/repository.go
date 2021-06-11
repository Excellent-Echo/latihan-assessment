package user

import (
	"backend/entity"

	"gorm.io/gorm"
)

type Repository interface {
	RShowAllBook() ([]entity.Book, error)
	RRegisterBook(Book entity.Book) (entity.Book, error)
	RFindBookByID(ID string) (entity.Book, error)
	RDeleteBookByID(ID string) (string, error)
	RUpdateBookByID(ID string, dataUpdate map[string]interface{}) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RShowAllBook() ([]entity.Book, error) {
	var Books []entity.Book

	err := r.db.Find(&Books).Error
	if err != nil {
		return Books, err
	}

	return Books, nil
}

func (r *repository) RRegisterBook(Book entity.Book) (entity.Book, error) {

	err := r.db.Create(&Book).Error
	if err != nil {
		return Book, err
	}

	return Book, nil
}

func (r *repository) RFindBookByID(ID string) (entity.Book, error) {
	var Book entity.Book

	err := r.db.Where("id = ?", ID).Find(&Book).Error
	if err != nil {
		return Book, err
	}

	return Book, nil
}

func (r *repository) RDeleteBookByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Book{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) RUpdateBookByID(ID string, dataUpdate map[string]interface{}) (entity.Book, error) {

	var Book entity.Book

	if err := r.db.Model(&Book).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return Book, err
	}

	if err := r.db.Where("id = ?", ID).Find(&Book).Error; err != nil {
		return Book, err
	}

	return Book, nil
}
