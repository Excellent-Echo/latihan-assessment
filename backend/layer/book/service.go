package book

import (
	"book-list/entity"
	"errors"
	"strconv"
	"time"
)

type Service interface {
	GetAllBooks() ([]BookOutput, error)
	CreateNewBook(book entity.BookInput) (BookOutput, error)
	GetBookByID(ID string) (entity.Books, error)
	UpdateBookByID(ID string, dataInput entity.BookInput) (entity.Books, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllBooks() ([]BookOutput, error) {
	books, err := s.repo.FindAllBook()

	var booksOutput []BookOutput

	for _, book := range books {
		bookOutput := BookOutputFormat(book)
		booksOutput = append(booksOutput, bookOutput)
	}

	if err != nil {
		return booksOutput, err
	}

	return booksOutput, nil

}

func (s *service) CreateNewBook(book entity.BookInput) (BookOutput, error) {
	year, _ := strconv.Atoi(book.Year)

	var newBook = entity.Books{
		Tittle:    book.Tittle,
		Author:    book.Author,
		Year:      year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createBook, err := s.repo.CreateBook(newBook)
	bookOutput := BookOutputFormat(createBook)
	if err != nil {
		return bookOutput, err
	}

	return bookOutput, nil
}

func (s *service) GetBookByID(ID string) (entity.Books, error) {
	book, err := s.repo.FindBookByID(ID)

	if err != nil {
		return entity.Books{}, err
	}

	if book.ID == 0 {
		return entity.Books{}, errors.New("Book ID Not Found")
	}

	return book, nil
}

func (s *service) UpdateBookByID(ID string, dataInput entity.BookInput) (entity.Books, error) {
	var dataUpdate = map[string]interface{}{}

	book, err := s.repo.FindBookByID(ID)

	if err != nil {
		return entity.Books{}, err
	}

	if book.ID == 0 {
		return entity.Books{}, errors.New("User ID not found")
	}

	if dataInput.Tittle != "" || len(dataInput.Tittle) != 0 {
		dataUpdate["tittle"] = dataInput.Tittle
	}
	if dataInput.Author != "" || len(dataInput.Author) != 0 {
		dataUpdate["author"] = dataInput.Author
	}
	if dataInput.Year != "" || len(dataInput.Year) != 0 {
		dataUpdate["year"] = dataInput.Year
	}

	dataUpdate["updated_at"] = time.Now()

	bookUpdated, err := s.repo.UpdateByID(ID, dataUpdate)

	if err != nil {
		return entity.Books{}, err
	}

	return bookUpdated, nil
}
