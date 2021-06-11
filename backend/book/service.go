package book

import (
	"github.com/marwanjuna/latihan-assessment/entity"
)

type BookService interface {
	GetAllBooks() ([]BookFormat, error)
	CreateNewBook(book entity.BookInput) (BookFormat, error)
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

func (s *bookService) CreateNewBook(book entity.BookInput) (BookFormat, error) {
	var newBook = entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	createBook, err := s.repository.PostBook(newBook)
	formatBook := FormattingBook(createBook)

	if err != nil {
		return formatBook, err
	}

	return formatBook, nil
}
