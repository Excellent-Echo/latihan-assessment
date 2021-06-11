package handler

import (
	"net/http"
	"strconv"

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

	response, err := h.bookService.CreateNewBook(inputBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

func (h *bookHandler) ShowBookByIDhandler(c *gin.Context) {
	id := c.Param("id")

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) UpdateBookByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var updateBookInput entity.BookInput

	if err := c.ShouldBindJSON(&updateBookInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": splitError,
		})
		return
	}

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	book, err := h.bookService.UpdateBookByID(id, updateBookInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) DeleteByBookIDHandler(c *gin.Context) {
	id := c.Param("id")

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	book, err := h.bookService.DeleteBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
