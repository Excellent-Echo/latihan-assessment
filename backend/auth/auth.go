package auth

import (
	"github.com/dgrijalva/jwt-go"
)

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

func (s *jwtService) GenerateToken(ID int) (string, error) {
	claim := jwt.MapClaims{
		"id": ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}
}
