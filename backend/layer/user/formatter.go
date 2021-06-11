package user

import (
	"book-list/entity"
	"time"
)

type UserOutput struct {
	ID        int    `gorm:"Primarykey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	DateBirth string `json:"date_birth"`
	Address   string `json:"address"`
}

type UserInput struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	DateBirth string `json:"date_birth"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func UserOutputFormat(user entity.Users) UserOutput {
	var userOutput = UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: FormatStringDateBirth(user.DateBirth),
		Address:   user.Address,
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
