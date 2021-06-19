package book

import "github.com/marwanjuna/latihan-assessment/entity"

type BookFormat struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func FormattingBook(book entity.Book) BookFormat {
	bookFormat := BookFormat{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookFormat
}
