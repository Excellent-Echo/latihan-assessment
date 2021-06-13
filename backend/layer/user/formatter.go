package user

import (
	"backend/entity"
	"time"
)

type UserFormat struct {
	ID        int    `json:"id"`
	Name      string `json:"username"`
	Email     string `json:"email"`
	DateBirth string `json:"date_birth"`
	Address   string `json:"address"`
}

type UserDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func Format(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: StringDateFormat(user.DateBirth),
		Address:   user.Address,
	}

	return formatUser
}

func FormatDelete(msg string) UserDeleteFormat {
	var deleteFormat = UserDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}

func DateFormat(date string) time.Time {
	format := "2006-01-02"
	value := date
	dateFormat, _ := time.Parse(format, value)

	return dateFormat
}

func StringDateFormat(date time.Time) string {
	dateFormat := date.Format("2006-01-02")

	return dateFormat
}
