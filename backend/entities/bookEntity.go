package entities

import "time"

type Book struct {
	Id        int
	UserId    int
	Title     string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
