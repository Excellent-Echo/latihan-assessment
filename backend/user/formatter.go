package user

import (
	"latihanassesment/entity"
	"time"
)

type UserFormat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: user.DateBirth,
	}

	return formatUser
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
