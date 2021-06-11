package handler

import (
	"book-list/entity"
	"book-list/layer/book"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) ShowAllBook(c *gin.Context) {
	books, err := h.service.GetAllBooks()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, books)
}

func (h *bookHandler) CreateNewBooks(c *gin.Context) {
	var input entity.BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBook, err := h.service.CreateNewBook(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)

}

func (h *bookHandler) ShowBookByID(c *gin.Context) {
	id := c.Param("book_id")

	book, err := h.service.GetBookByID(id)
	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

func (h *bookHandler) UpdateBookByIDHandler(c *gin.Context) {
	id := c.Param("book_id")

	var updateUserInput entity.BookInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {

		c.JSON(400, gin.H{"error": err})
		return
	}

	book, err := h.service.UpdateBookByID(id, updateUserInput)
	if err != nil {

		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, book)
}
