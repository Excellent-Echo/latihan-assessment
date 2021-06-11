package user

import (
	"latihan-assessment/backend/entity"

	"latihan-assessment/backend/user"
)

type Service interface {
	GetAllUser() ([]entity.Users, error)
}

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) *service {
	return &service(repository)
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	Users, err := s.repository.GetAll()

	var usersFormat []UserFormat

	for _, user := range Users {
		var userFormat = UserFormatting(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}
