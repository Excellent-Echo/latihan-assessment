package book

import (
	"errors"
	"fmt"
	"latihanassesment/entity"
	"latihanassesment/helper"
	"time"
)

type Service interface {
	GetAllBook() ([]entity.Book, error)
	SaveNewBook(book entity.BookInput) (entity.Book, error)
	GetBookByID(ID string) (entity.Book, error)
	UpdateBookByID(bookID string, dataInput entity.BookInput) (entity.Book, error)
	DeleteBookByID(bookID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllBook() ([]entity.Book, error) {
	books, err := s.repository.FindAll()

	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetBookByID(ID string) (entity.Book, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Book{}, err
	}

	book, err := s.repository.FindBYID(ID)

	if err != nil {
		return entity.Book{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", ID)
		return entity.Book{}, errors.New(newError)
	}

	return book, nil
}

func (s *service) SaveNewBook(book entity.BookInput) (entity.Book, error) {

	var newBook = entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
		UserID: book.UserID,
	}

	createBook, err := s.repository.Create(newBook)

	if err != nil {
		return createBook, err
	}

	return createBook, nil

}

func (s *service) UpdateBookByID(bookID string, dataInput entity.BookInput) (entity.Book, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(bookID); err != nil {
		return entity.Book{}, err
	}

	user, err := s.repository.FindBYID(bookID)

	if err != nil {
		return entity.Book{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("book id %s not found", bookID)
		return entity.Book{}, errors.New(newError)
	}

	if dataInput.Author != "" || len(dataInput.Author) != 0 {
		dataUpdate["author"] = dataInput.Author
	}
	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Year == 0 {
		dataUpdate["year"] = dataInput.Year
	}
	if dataInput.UserID == 0 {
		dataUpdate["user_id"] = dataInput.Title
	}

	dataUpdate["update_at"] = time.Now()

	fmt.Println(dataUpdate)

	bookUpdated, err := s.repository.UpdateByID(bookID, dataUpdate)

	if err != nil {
		return entity.Book{}, err
	}

	return bookUpdated, nil
}

func (s *service) DeleteBookByID(bookID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(bookID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindBYID(bookID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("book id %s not found", bookID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(bookID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", bookID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}
