package user

import (
	"latihan-assessment/entity"
	"time"
)

type UserFormat struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	DateBirth string `json:"date_birth"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: user.DateBirth,
		Email:     user.Email,
		Password:  user.Password,
	}
	return formatUser
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDeleteUser(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
