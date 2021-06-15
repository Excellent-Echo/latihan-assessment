package handler

import (
	"book-list/entity"
	"book-list/layer/book"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

// GetBooks godoc
// @Security Auth
// @Summary get all book
// @Description Get All book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} book.BookOutput
// @Failure 500 {object} helper.Response
// @Router /books [get]
func (h *bookHandler) ShowAllBook(c *gin.Context) {
	books, err := h.service.GetAllBooks()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, books)
}

// CreateBook godoc
// @Security Auth
// @Summary Create new book
// @Description Create new book
// @Tags books
// @Param book body entity.BookInput true "book data"
// @Success 200 {object} book.BookOutputDetail
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /books [post]
func (h *bookHandler) CreateNewBooks(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))
	id := strconv.Itoa(userData)

	var input entity.BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBook, err := h.service.CreateNewBook(input, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)

}

// GetBookByID godoc
// @Security Auth
// @Summary Get Book by ID
// @Description Get a Book By ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} entity.Books
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /books/{id} [get]
func (h *bookHandler) ShowBookByID(c *gin.Context) {
	id := c.Param("book_id")

	book, err := h.service.GetBookByID(id)
	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

// UpdateBook godoc
// @Security Auth
// @Summary Update book
// @Description update book by id
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Param book body entity.BookInput true "update book"
// @Success 200 {object} entity.Books
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /books/{id} [put]
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

// DeleteBook godoc
// @Security Auth
// @Summary Delete book
// @Description Delete a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} helper.DeleteFormat
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /books/{id} [delete]
func (h *bookHandler) DeleterBookByID(c *gin.Context) {
	id := c.Param("book_id")

	user, err := h.service.DeleteBookByID(id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
