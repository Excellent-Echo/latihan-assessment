package auth

var SECRET_KEY = []byte("rahasia")

type Service interface {
	GenerateToken(ID int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}
