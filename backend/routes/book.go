package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/book"
	"github.com/marwanjuna/latihan-assessment/handler"
)

var (
	bookRepository = book.NewRepository(DB)
	bookService    = book.NewService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService)
)

func BookRoute(r *gin.Engine) {
	r.GET("/books", bookHandler.ShowAllBooksHandler)
	r.POST("/books", bookHandler.CreateBookHandler)
	r.GET("/books/:id", bookHandler.ShowBookByIDhandler)
	r.PUT("/books/:id", bookHandler.UpdateBookByIDHandler)
	r.DELETE("/books/:id", bookHandler.DeleteByBookIDHandler)
}
