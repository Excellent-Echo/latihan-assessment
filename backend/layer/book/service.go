package book

import (
	"book-list/entity"
	"time"
)

type Service interface {
	GetAllBooks() ([]BookOutput, error)
	CreateNewBook(book entity.Books) (BookOutput, error)
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

func (s *service) CreateNewBook(book entity.Books) (BookOutput, error) {

	var newBook = entity.Books{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: FormatDateBirth(user.DateBirth),
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repo.CreateUser(newUser)
	userOutput := UserOutputFormat(createUser)
	if err != nil {
		return userOutput, err
	}

	return userOutput, nil
}
