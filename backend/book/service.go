package book

import (
	"errors"
	"fmt"
	"latihan-assessment/entity"
	"latihan-assessment/helper"
	"time"
)

type Service interface {
	AllBooks() ([]BookFormat, error)
	NewBook(book entity.BookInput) (BookFormat, error)
	BookById(bookID string) (BookFormat, error)
	DeleteById(bookID string) (interface{}, error)
	UpdateById(bookID string, dataUpdate entity.BookInputUpdate) (BookFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AllBooks() ([]BookFormat, error) {
	books, err := s.repository.GetAllBooks()
	var formatBooks []BookFormat

	for _, book := range books {
		formatbook := FormatBook(book)
		formatBooks = append(formatBooks, formatbook)
	}

	if err != nil {
		return formatBooks, err
	}
	return formatBooks, nil
}

func (s *service) NewBook(book entity.BookInput) (BookFormat, error) {
	var newBook = entity.Books{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	createBook, err := s.repository.CreateBooks(newBook)

	formatBook := FormatBook(createBook)

	if err != nil {
		return formatBook, err
	}

	return formatBook, nil
}

func (s *service) BookById(bookID string) (BookFormat, error) {
	if err := helper.ValidateIDNumber(bookID); err != nil {
		return BookFormat{}, err
	}

	book, err := s.repository.GetBooksById(bookID)

	if err != nil {
		return BookFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id is not found : %s ", bookID)
		return BookFormat{}, errors.New(newError)
	}

	formatBook := FormatBook(book)

	return formatBook, nil
}

func (s *service) DeleteById(bookID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(bookID); err != nil {
		return BookFormat{}, err
	}

	book, err := s.repository.GetBooksById(bookID)

	if err != nil {
		return nil, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book ID %s not found", bookID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteBookById(bookID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("success delete book by ID : %s", bookID)

	formatDelete := FormatDeleteBook(message)

	return formatDelete, nil

}

func (s *service) UpdateById(bookID string, dataUpdate entity.BookInputUpdate) (BookFormat, error) {
	var BookUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(bookID); err != nil {
		return BookFormat{}, err
	}

	book, err := s.repository.GetBooksById(bookID)

	if err != nil {
		return BookFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id not found : %s", bookID)
		return BookFormat{}, errors.New(newError)
	}

	if dataUpdate.Title != "" || len(dataUpdate.Title) != 0 {
		BookUpdate["title"] = dataUpdate.Title
	}
	if dataUpdate.Author != "" || len(dataUpdate.Author) != 0 {
		BookUpdate["author"] = dataUpdate.Author
	}

	if dataUpdate.Year != 0 {
		BookUpdate["year"] = dataUpdate.Year
	}

	BookUpdate["updated_at"] = time.Now()

	bookUpdated, err := s.repository.UpdateBookById(bookID, BookUpdate)

	if err != nil {
		return BookFormat{}, err
	}

	formatbook := FormatBook(bookUpdated)

	return formatbook, nil
}
