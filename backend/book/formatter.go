package book

import (
	"latihan-assessment/entity"
	"time"
)

type BookFormat struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatBook(book entity.Books) BookFormat {
	var formatBook = BookFormat{
		ID:     book.ID,
		UserID: book.UserID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	return formatBook
}

func FormatDeleteBook(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
