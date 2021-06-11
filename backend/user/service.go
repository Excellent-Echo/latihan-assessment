package user

type Service interface {
	GetAllUser() ([]UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	users, err := s.repository.GetAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}
