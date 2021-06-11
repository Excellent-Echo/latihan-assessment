package user

import (
	"book-list/entity"
	"time"
)

type UserOutput struct {
	ID        int       `gorm:"Primarykey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	DateBirth time.Time `json:"date_birth"`
	Address   string    `json:"address"`
}

func UserOutputFormat(user entity.Users) UserOutput {
	var userOutput = UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: user.DateBirth,
		Address:   user.Address,
	}
	return userOutput
}
