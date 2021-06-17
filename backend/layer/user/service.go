package user

import (
	"book-list/entity"
	"book-list/helper"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]entity.UserOutput, error)
	CreateNewUser(user entity.UserInput) (entity.UserOutput, error)
	LoginUser(loginInput entity.LoginInput) (entity.Users, error)
	GetUserByID(ID string) (entity.UserOutput, error)
	UpdateUserByID(ID string, dataInput entity.UserUpdateInput) (entity.UserOutput, error)
	DeleteUserByID(ID string) (interface{}, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUsers() ([]entity.UserOutput, error) {
	users, err := s.repo.FindAll()

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

func (s *service) GetUserByID(ID string) (entity.UserOutput, error) {
	user, err := s.repo.FindUserByID(ID)

	if err != nil {
		return entity.UserOutput{}, err
	}

	if user.ID == 0 {
		return entity.UserOutput{}, errors.New("User ID Not Found")
	}

	FormatOutput := UserOutputFormat(user)
	return FormatOutput, nil
}

func (s *service) UpdateUserByID(ID string, dataInput entity.UserUpdateInput) (entity.UserOutput, error) {
	var dataUpdate = map[string]interface{}{}

	user, err := s.repo.FindUserByID(ID)

	if err != nil {
		return entity.UserOutput{}, err
	}

	if user.ID == 0 {
		return entity.UserOutput{}, errors.New("User ID not found")
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

	dataUpdate["updated_at"] = time.Now()

	userUpdated, err := s.repo.UpdateByID(ID, dataUpdate)

	if err != nil {
		return entity.UserOutput{}, err
	}

	formatUser := UserOutputFormat(userUpdated)

	return formatUser, nil
}

func (s *service) DeleteUserByID(ID string) (interface{}, error) {
	user, err := s.repo.FindUserByID(ID)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user id not found")
	}

	status, err := s.repo.DeleteByID(ID)

	if status == "error" {
		return nil, errors.New("error in internal server")
	}

	data := fmt.Sprintf("user id %s success deleted", ID)

	formatDelete := helper.FormatDelete(data)

	return formatDelete, nil
}
