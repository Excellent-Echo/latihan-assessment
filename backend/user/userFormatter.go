package user

import (
	"latihan-assessment/backend/entity"
	"time"
)

type UserFormat struct {
	ID        uint      `json: "id"`
	Name      string    `json: "name"`
	Email     string    `json: "email"`
	Address   string    `json: "address"`
	DateBirth time.Time `json: "date_birth`
}

func UserFormatting(user entity.Users) UserFormat {
	UserFormat := UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		DateBirth: user.DateBirth,
	}

	return UserFormat
}
