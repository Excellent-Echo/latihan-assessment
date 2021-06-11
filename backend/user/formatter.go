package user

import (
	"books-project/entity"
	"time"
)

type UserFormat struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	DateBirth time.Time `json:"date_birth"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		DateBirth: user.DateBirth,
	}

	return userFormat
}
