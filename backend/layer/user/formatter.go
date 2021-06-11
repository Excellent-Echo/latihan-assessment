package user

import (
	"book-list/entity"
	"time"
)

type UserLoginOutput struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"string"`
	DateBirth     string `json:"date_birth"`
	Email         string `json:"email"`
	Authorization string `json:"authorization"`
}

func UserOutputFormat(user entity.Users) entity.UserOutput {
	var userOutput = entity.UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		DateBirth: FormatStringDateBirth(user.DateBirth),
		Email:     user.Email,
	}
	return userOutput
}

func FormatDateBirth(date string) time.Time {
	layoutFormat := "2006-01-02"
	value := date
	dateBirth, _ := time.Parse(layoutFormat, value)

	return dateBirth
}

func FormatStringDateBirth(date time.Time) string {
	dateBirth := date.Format("2006-01-02")

	return dateBirth
}
