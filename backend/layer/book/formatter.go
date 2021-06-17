package book

import (
	"book-list/entity"
	"time"
)

type BookOutput struct {
	ID     int    `json:"id"`
	Tittle string `json:"tittle"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookOutputDetail struct {
	ID        int       `gorm:"primarykey" json:"id"`
	Tittle    string    `json:"tittle"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func BookOutputFormat(book entity.Books) BookOutput {
	var bookOuput = BookOutput{
		ID:     book.ID,
		Tittle: book.Tittle,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookOuput
}

func FormatOutputDetail(book entity.Books) BookOutputDetail {
	var bookOutput = BookOutputDetail{
		ID:        book.ID,
		Tittle:    book.Tittle,
		Author:    book.Author,
		Year:      book.Year,
		UserID:    book.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return bookOutput
}
