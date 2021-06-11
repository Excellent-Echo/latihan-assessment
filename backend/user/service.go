package user

type UserService interface {
	GetAllUsers() ([]UserFormat, error)
}

type userService struct {
	repository UserRepository
}

func NewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetAllUsers() ([]UserFormat, error) {
	users, err := s.repository.GetAllUsers()

	var usersFormat []UserFormat

	for _, user := range users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}
