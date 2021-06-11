package user

import "latihan-assessment/entity"

type Service interface {
	GetAllUser() ([]entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service)GetAllUser() ([]entity.User, error)   {
	Users, err := s.repository.GetAll()
}