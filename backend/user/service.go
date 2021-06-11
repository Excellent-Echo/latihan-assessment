package user

import (
	"errors"
	"fmt"
	"latihan-assessment/entity"
	"latihan-assessment/helper"
)

type Service interface {
	AllUser() ([]UserFormat, error)
	NewUser(user entity.UserInput) (UserFormat, error)
	UserById(userID string) (UserFormat, error)
	DeleteById(userID string) (interface{}, error)
	UpdateById(userID string, dataUpdate entity.UserInputUpdate) (UserFormat, error)
	LoginUser(input entity.UserLoginInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AllUser() ([]UserFormat, error) {
	users, err := s.repository.GetAllUsers()
	var formatUsers []UserFormat

	for _, user := range users {
		formatuser := FormatUser(user)
		formatUsers = append(formatUsers, formatuser)
	}

	if err != nil {
		return formatUsers, err
	}
	return formatUsers, nil
}

func (s *service) NewUser(user entity.UserInput) (UserFormat, error) {
	var newUser = entity.User{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  user.Password,
	}

	createUser, err := s.repository.CreateUsers(newUser)

	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) UserById(userID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetUserById(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprint("user id is not found : %s ", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)

	return formatUser, nil
}

func (s *service) DeleteById(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetUserById(userID)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprint("user ID %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteUserById(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("success delete user by ID : %s", userID)

	formatDelete := FormatDeleteUser(message)

	return formatDelete, nil

}

//func (s *service) UpdateById(userID string, dataUpdate entity.UserInputUpdate) (UserFormat, error) {
//
//}
//
//func (s *service) LoginUser(input entity.UserLoginInput) (entity.User, error) {
//
//}
