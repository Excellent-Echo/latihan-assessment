package user

import (
	"errors"
	"fmt"

	"github.com/marwanjuna/latihan-assessment/entity"
	"github.com/marwanjuna/latihan-assessment/helper"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]UserFormat, error)
	CreateNewUser(user entity.UserInput) (UserFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
	GetUserByID(id string) (UserFormat, error)
	UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error)
	DeleteUserByID(id string) (interface{}, error)
}

type userService struct {
	repository UserRepository
}

func NewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetAllUsers() ([]UserFormat, error) {
	users, err := s.repository.GetAllUsers()

	var usersFormat []UserFormat

	for _, user := range users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}

func (s *userService) CreateNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  string(genPassword),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userService) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := ("Email/password is incorrect")
		return user, errors.New(newError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) GetUserByID(id string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	userFormat := FormattingUser(user)

	return userFormat, nil
}

func (s *userService) UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Address != "" || len(dataInput.Address) != 0 {
		dataUpdate["address"] = dataInput.Address
	}
	if dataInput.DateBirth != "" || len(dataInput.DateBirth) != 0 {
		dataUpdate["date_birth"] = dataInput.DateBirth
	}

	userUpdated, err := s.repository.UpdateUser(id, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormattingUser(userUpdated)

	return formatUser, nil
}

func (s *userService) DeleteUserByID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteUser(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("user id %s success deleted", id)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}
