package user

import (
	"book-list/entity"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]entity.UserOutput, error)
	CreateNewUser(user entity.UserInput) (entity.UserOutput, error)
	LoginUser(loginInput entity.LoginInput) (entity.Users, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUsers() ([]entity.UserOutput, error) {
	users, err := s.repo.FindAll()

	fmt.Print(users)
	var usersOutput []entity.UserOutput

	for _, user := range users {
		userOutput := UserOutputFormat(user)
		usersOutput = append(usersOutput, userOutput)
	}

	if err != nil {
		return usersOutput, err
	}

	return usersOutput, nil

}

func (s *service) CreateNewUser(user entity.UserInput) (entity.UserOutput, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	var newUser = entity.Users{
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

func (s *service) LoginUser(loginInput entity.LoginInput) (entity.Users, error) {
	user, err := s.repo.FindUserByEmail(loginInput.Email)

	if err != nil {
		return user, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}
