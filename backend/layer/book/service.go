package book

import (
	"backend/entity"
	"backend/helper"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Service interface {
	SShowAllBook() ([]BookFormat, error)
	SRegisterUser(book entity.BookInput) (BookFormat, error)
	SFindBookByID(bookID string) (BookFormat, error)
	SDeleteBookByID(bookID string) (interface{}, error)
	SUpdateBookByID(bookID string, input entity.UpdateBookInput) (BookFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SShowAllUserr() ([]BookFormat, error) {
	Books, err := s.repository.RShowAllBook()
	var formatBooks []BookFormat

	for _, book := range Books {
		formatBook := FormatBook(book)
		formatBooks = append(formatBooks, formatBook)

	}
	if err != nil {
		return formatBooks, err
	}

	return formatBooks, nil
}

func (s *service) SRegisterBook(book entity.BookInput) (BookFormat, error) {

	year, _ := strconv.Atoi(book.Year)

	var newBook = entity.Book{
		Tittle:    book.Tittle,
		Author:    book.Author,
		Year:      year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createBook, err := s.repository.RRegisterBook(newBook)
	formatBook := FormatBook(createBook)

	if err != nil {
		return formatBook, err
	}

	return formatBook, nil
}

func (s *service) SFindBookByID(bookID string) (BookFormat, error) {
	book, err := s.repository.RFindBookByID(bookID)

	if err != nil {
		return BookFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s not found", bookID)
		return BookFormat{}, errors.New(newError)
	}

	formatBook := FormatBook(book)

	return formatBook, nil
}

func (s *service) SDeleteBookByID(bookID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(bookID); err != nil {
		return nil, err
	}

	book, err := s.repository.RFindBookByID(bookID)

	if err != nil {
		return nil, err
	}
	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s not found", bookID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.RDeleteBookByID(bookID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete book ID : %s", bookID)

	formatDelete := FormatDeleteBook(msg)

	return formatDelete, nil
}

func (s *service) SUpdateBookByID(bookID string, input entity.UpdateBookInput) (BookFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(bookID); err != nil {
		return BookFormat{}, err
	}

	book, err := s.repository.RFindBookByID(bookID)

	if err != nil {
		return BookFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s not found", bookID)
		return BookFormat{}, errors.New(newError)
	}

	if input.Tittle != "" || len(input.Tittle) != 0 {
		dataUpdate["Tittle"] = input.Tittle
	}
	if input.Author != "" || len(input.Author) != 0 {
		dataUpdate["Author"] = input.Author
	}
	if input.Year != "" || len(input.Year) != 0 {
		dataUpdate["Year"] = input.Year
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	bookUpdated, err := s.repository.RUpdateBookByID(bookID, dataUpdate)

	if err != nil {
		return BookFormat{}, err
	}

	formatBook := FormatBook(bookUpdated)

	return formatBook, nil
}
