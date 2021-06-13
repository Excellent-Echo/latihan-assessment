package user

import (
	"backend/entity"
	"backend/helper"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SShowAllUser() ([]UserFormat, error)
	SRegisterUser(user entity.UserInput) (UserFormat, error)
	SFindUserByID(userID string) (UserFormat, error)
	SDeleteUserByID(userID string) (interface{}, error)
	SUpdateUserByID(userID string, input entity.UpdateUserInput) (UserFormat, error)
	SLoginUser(input entity.LoginUserInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SLoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.RFindUserByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	//ngecek password sama atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}

func (s *service) SShowAllUser() ([]UserFormat, error) {
	Users, err := s.repository.RShowAllUser()
	var formatUsers []UserFormat

	for _, user := range Users {
		formatUser := Format(user)
		formatUsers = append(formatUsers, formatUser)

	}
	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SRegisterUser(user entity.UserInput) (UserFormat, error) {
	//bcryp password jadi hash agar susah diretas
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: DateFormat(user.DateBirth),
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.RRegisterUser(newUser)
	formatUser := Format(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
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
