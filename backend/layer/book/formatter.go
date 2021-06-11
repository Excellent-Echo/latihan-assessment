package book

import (
	"backend/entity"
	"time"
)

type BookFormat struct {
	ID     int    `json:"id"`
	Tittle string `json:"tittle"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatBook(Book entity.Book) BookFormat {
	var formatBook = BookFormat{
		ID:     Book.ID,
		Tittle: Book.Tittle,
		Author: Book.Author,
		Year:   Book.Year,
	}

	return formatBook
}

func FormatDeleteBook(msg string) BookDeleteFormat {
	var deleteFormat = BookDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
