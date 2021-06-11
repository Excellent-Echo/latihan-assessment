package user

import (
	"time"

	"github.com/marwanjuna/latihan-assessment/entity"
)

type UserFormat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	DateBirth time.Time `json:"date_birth"`
	Address   string    `json:"address"`
}

type UserInputFormat struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: user.DateBirth,
		Address:   user.Address,
	}

	return userFormat
}

func FormattingUserInput(user entity.User) UserInputFormat {
	userFormat := UserInputFormat{
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  user.Password,
	}

	return userFormat
}
