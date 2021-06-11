package user

import (
	"github.com/marwanjuna/latihan-assessment/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]UserFormat, error)
	CreateNewUser(user entity.UserInput) (UserInputFormat, error)
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

func (s *userService) CreateNewUser(user entity.UserInput) (UserInputFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserInputFormat{}, err
	}

	var newUser = entity.User{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  string(genPassword),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUserInput(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}
