package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/book"
	"github.com/marwanjuna/latihan-assessment/entity"
	"github.com/marwanjuna/latihan-assessment/helper"
)

type bookHandler struct {
	bookService book.BookService
}

func NewBookHandler(bookService book.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) ShowAllBooksHandler(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var inputBook entity.BookInput

	if err := c.ShouldBindJSON(&inputBook); err != nil {
		splitError := helper.SplitErrorInformation(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": splitError,
		})
		return
	}

}
