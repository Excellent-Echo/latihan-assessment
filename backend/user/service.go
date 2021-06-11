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
	users, err := s.repository.GetAllUsers()
	var formatUsers []UserFormat

	for _, user := range users {
		formatuser := FormatUser(user)
		formatUsers = append(formatUsers, formatuser)
	}

	if err != nil {
		return formatUsers, err
	}
	return formatUsers, nil
}
