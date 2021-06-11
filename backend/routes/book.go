package routes

import (
	"latihanassesment/book"
	"latihanassesment/handler"

	"github.com/gin-gonic/gin"
)

var (
	bookRepository = book.NewRepository(DB)
	bookService    = book.NewService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService, authService)
)

func BookRoute(r *gin.Engine) {
	r.GET("/books", handler.Middleware(userService, authService), bookHandler.ShowBookHandler)
	r.POST("/books/register", bookHandler.CreateBookHandler)
	r.GET("/books/:book_id", handler.Middleware(userService, authService), bookHandler.GetBookByIDHandler)
	r.DELETE("/books/:book_id", handler.Middleware(userService, authService), bookHandler.DeleteBookByIDHandler)
	r.PUT("/books/:book_id", handler.Middleware(userService, authService), bookHandler.UpdateBookByIDHandler)
}
