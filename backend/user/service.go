package user

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}

}

func (s *service) GetAllUsers() ([]entiity.User, error) {
	Users, err := s.repository.GetAll()

	if err != nil {
		return Users, err
	}

	return Users, nil
}
