package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var key = os.Getenv("DB_SECRET_KEY")

type Service interface {
	GenerateToken(ID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
