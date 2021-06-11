package user

import (
	"time"

	"github.com/marwanjuna/latihan-assessment/entity"
)

type UserFormat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	DateBirth time.Time `json:"date_birth"`
	Address   string    `json:"address"`
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
