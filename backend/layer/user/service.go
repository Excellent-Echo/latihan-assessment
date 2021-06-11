package user

type Service interface {
	GetAllUsers() ([]UserOutput, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUsers() ([]UserOutput, error) {
	users, err := s.repo.FindAll()

	var usersOutput []UserOutput

	for _, user := range users {
		userOutput := UserOutputFormat(user)
		usersOutput = append(usersOutput, userOutput)
	}

	if err != nil {
		return usersOutput, err
	}

	return usersOutput, nil

}
