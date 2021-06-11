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
	SFindUserByID(userID string) (BookFormat, error)
	SDeleteUserByID(userID string) (interface{}, error)
	SUpdateUserByID(userID string, input entity.UpdateBookInput) (BookFormat, error)
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

	strconv.Itoa(book.Year)

	var newBook = entity.Book{
		Tittle:    book.Tittle,
		Author:    book.Author,
		Year:      book.Year,
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

func (s *service) SFindUserByID(userID string) (UserFormat, error) {
	user, err := s.repository.RFindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := Format(user)

	return formatUser, nil
}

func (s *service) SDeleteUserByID(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return nil, err
	}

	user, err := s.repository.RFindUserByID(userID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.RDeleteUserByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}

func (s *service) SUpdateUserByID(userID string, input entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.RFindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	if input.Name != "" || len(input.Name) != 0 {
		dataUpdate["Name"] = input.Name
	}
	if input.Address != "" || len(input.Address) != 0 {
		dataUpdate["Address"] = input.Address
	}

	if input.DateBirth != "" || len(input.DateBirth) != 0 {
		dataUpdate["DateBirth"] = input.DateBirth
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userUpdated, err := s.repository.RUpdateUserByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := Format(userUpdated)

	return formatUser, nil
}
