package user

import (
	"book-list/entity"
	"time"
)

type Service interface {
	GetAllUsers() ([]UserOutput, error)
	CreateNewUser(user UserInput) (UserOutput, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUsers() ([]UserOutput, error) {
	users, err := s.repo.FindAll()

	var usersOutput []UserOutput

	for _, user := range users {
		userOutput := UserOutputFormat(user)
		usersOutput = append(usersOutput, userOutput)
	}

	if err != nil {
		return usersOutput, err
	}

	return usersOutput, nil

}

func (s *service) CreateNewUser(user UserInput) (UserOutput, error) {
	var newUser = entity.Users{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  user.Password,
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
