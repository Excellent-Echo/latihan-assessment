package book

import "book-list/entity"

type BookOutput struct {
	ID     int    `json:"id"`
	Tittle string `json:"tittle"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func BookOutputFormat(book entity.Books) BookOutput {
	var bookOuput = BookOutput{
		Tittle: book.Tittle,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookOuput
}
