package services

import (
	"book_list/entities"
	"book_list/inputs"
	"book_list/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUpUser(inputSignUp inputs.UserInputSignUp) (entities.User, error)
	SignInUser(inputSignIn inputs.UserInputSignIn) (entities.User, error)
	UpdateUserById(inputId inputs.InputUserId, inputData inputs.UpdateUserInput) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserById(inputId inputs.InputUserId) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) SignUpUser(inputSignUp inputs.UserInputSignUp) (entities.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputSignUp.Password), bcrypt.MinCost)
	userData := entities.User{
		Name:     inputSignUp.Name,
		Email:    inputSignUp.Email,
		Password: string(passwordHash),
	}

	if err != nil {
		return userData, err
	}

	newUserData, err := s.userRepository.InsertUser(userData)
	if err != nil {
		return newUserData, err
	}

	return newUserData, nil
}

func (s *userService) SignInUser(inputSignIn inputs.UserInputSignIn) (entities.User, error) {
	userData := entities.User{
		Email: inputSignIn.Email,
		Password: inputSignIn.Password,
	}
	userData, err := s.userRepository.FindUserByEmail(userData.Email)
	if err != nil {
		return userData, err
	}

	if userData.Id == 0 {
		return userData, errors.New("no user found with this email")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userData.Password)); err != nil {
		return userData, err
	}

	return userData, nil
}

func (s *userService) UpdateUserById(inputId inputs.InputUserId, inputUpdateData inputs.UpdateUserInput) (entities.User, error) {
	userData, err := s.userRepository.FindUserById(inputId.Id)
	if err != nil {
		return userData, nil
	}

	userData.Name = inputUpdateData.Name
	userData.Email = inputUpdateData.Email

	if userData.Id == 0 {
		return userData, errors.New("not the user")
	}

	updateDataUser, err := s.userRepository.UpdateUser(userData)
	if err != nil {
		return updateDataUser, nil
	}

	return updateDataUser, nil
}

func (s *userService) GetAllUsers() ([]entities.User, error) {
	var userData []entities.User
	userData, err := s.userRepository.FindAllUsers()

	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (s *userService) GetUserById(inputId inputs.InputUserId) (entities.User, error) {
	userData, err := s.userRepository.FindUserById(inputId.Id)
	if err != nil {
		return userData, err
	}

	if userData.Id == 0 {
		return userData, errors.New("there aren't users with this id")
	}

	return userData, nil
}