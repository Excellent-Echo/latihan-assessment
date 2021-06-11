package book

import (
	"errors"
	"fmt"

	"github.com/marwanjuna/latihan-assessment/entity"
	"github.com/marwanjuna/latihan-assessment/helper"
	"github.com/marwanjuna/latihan-assessment/user"
)

type BookService interface {
	GetAllBooks() ([]BookFormat, error)
	CreateNewBook(book entity.BookInput) (entity.Book, error)
	GetBookByID(id string) (entity.Book, error)
	UpdateBookByID(id string, dataInput entity.BookInput) (entity.Book, error)
	DeleteBookByID(id string) (interface{}, error)
}

type bookService struct {
	repository BookRepository
}

func NewService(repository BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) GetAllBooks() ([]BookFormat, error) {
	books, err := s.repository.GetAllBooks()

	var booksFormat []BookFormat

	for _, book := range books {
		var bookFormat = FormattingBook(book)
		booksFormat = append(booksFormat, bookFormat)
	}

	if err != nil {
		return booksFormat, err
	}

	return booksFormat, nil
}

func (s *bookService) CreateNewBook(book entity.BookInput) (entity.Book, error) {
	var newBook = entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	createBook, err := s.repository.PostBook(newBook)

	if err != nil {
		return createBook, err
	}

	return createBook, nil
}

func (s *bookService) GetBookByID(id string) (entity.Book, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.Book{}, err
	}

	book, err := s.repository.GetOneBook(id)

	if err != nil {
		return entity.Book{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s is not found", id)
		return entity.Book{}, errors.New(newError)
	}

	return book, nil
}

func (s *bookService) UpdateBookByID(id string, dataInput entity.BookInput) (entity.Book, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.Book{}, err
	}

	book, err := s.repository.GetOneBook(id)

	if err != nil {
		return entity.Book{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s is not found", id)
		return entity.Book{}, errors.New(newError)
	}

	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Author != "" || len(dataInput.Author) != 0 {
		dataUpdate["author"] = dataInput.Author
	}
	if dataInput.Year != 0 {
		dataUpdate["year"] = dataInput.Year
	}

	bookUpdated, err := s.repository.UpdateBook(id, dataUpdate)

	if err != nil {
		return entity.Book{}, err
	}

	return bookUpdated, nil
}

func (s *bookService) DeleteBookByID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	book, err := s.repository.GetOneBook(id)

	if err != nil {
		return nil, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("book id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteBook(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("book id %s success deleted", id)

	formatDelete := user.FormatDelete(msg)

	return formatDelete, nil
}
