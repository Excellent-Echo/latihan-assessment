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

func (s *service) GetAllUser() ([]entity.User, error) {
	Users, err := s.repository.GetAll()

	if err != nil {
		return Users, err
	}

	return Users, nil
}
