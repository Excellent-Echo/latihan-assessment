package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/book"
	"github.com/marwanjuna/latihan-assessment/handler"
)

var (
	bookRepository = book.NewRepository(DB)
	bookService    = book.NewService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService, authService)
)

func BookRoute(r *gin.Engine) {
	r.GET("/books", handler.Middleware(userService, authService), bookHandler.ShowAllBooksHandler)
	r.POST("/books", handler.Middleware(userService, authService), bookHandler.CreateBookHandler)
	r.GET("/books/:id", handler.Middleware(userService, authService), bookHandler.ShowBookByIDhandler)
	r.PUT("/books/:id", handler.Middleware(userService, authService), bookHandler.UpdateBookByIDHandler)
	r.DELETE("/books/:id", handler.Middleware(userService, authService), bookHandler.DeleteByBookIDHandler)
}
