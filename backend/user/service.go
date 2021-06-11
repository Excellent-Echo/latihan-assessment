package user

import (
	"books-project/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetAllUser() ([]UserFormat, error)
}

type userService struct {
	repository UserRepository
}

func UserNewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Address:   user.Address,
		DateBirth: user.DateBirth,
	}
	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userService) GetAllUser() ([]UserFormat, error) {
	users, err := s.repository.ReadAllUser()
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
