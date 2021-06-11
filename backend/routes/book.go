package routes

import (
	"book-list/handler"
	"book-list/layer/book"

	"github.com/gin-gonic/gin"
)

var (
	bookRepository = book.NewRepository(DB)
	bookService    = book.NewService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService)
)

func RouteBook(r *gin.Engine) {
	r.GET("/books", handler.Middleware(authService), bookHandler.ShowAllBook)
	r.GET("/books/:book_id", handler.Middleware(authService), bookHandler.ShowBookByID)
	r.POST("/books", handler.Middleware(authService), bookHandler.CreateNewBooks)
	r.PUT("/books/:book_id", handler.Middleware(authService), bookHandler.UpdateBookByIDHandler)
}
