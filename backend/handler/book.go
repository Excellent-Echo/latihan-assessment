package handler

import (
	"latihanassesment/auth"
	"latihanassesment/book"
	"latihanassesment/entity"
	"latihanassesment/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
	authService auth.Service
}

func NewBookHandler(bookService book.Service, authService auth.Service) *bookHandler {
	return &bookHandler{bookService, authService}
}

func (h *bookHandler) ShowBookHandler(c *gin.Context) {
	users, err := h.bookService.GetAllBook()

	if err != nil {

		c.JSON(500, users)
		return
	}

	c.JSON(200, users)
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var inputUser entity.BookInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SpilitErrorInformation(err)

		c.JSON(400, splitError)
		return
	}

	newUser, err := h.bookService.SaveNewBook(inputUser)
	if err != nil {
		c.JSON(500, newUser)
		return
	}

	c.JSON(201, newUser)
}

func (h *bookHandler) GetBookByIDHandler(c *gin.Context) {
	id := c.Params.ByName("book_id")

	user, err := h.bookService.GetBookByID(id)
	if err != nil {

		c.JSON(400, user)
		return
	}

	c.JSON(200, user)
}

func (h *bookHandler) DeleteBookByIDHandler(c *gin.Context) {
	id := c.Params.ByName("book_id")

	user, err := h.bookService.DeleteBookByID(id)

	if err != nil {
		c.JSON(400, user)
		return
	}

	c.JSON(200, user)
}

func (h *bookHandler) UpdateBookByIDHandler(c *gin.Context) {
	id := c.Params.ByName("book_id")

	var updateUserInput entity.BookInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		c.JSON(400, updateUserInput)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	UserData := int(c.MustGet("currentUser").(int))

	if idParam != UserData {

		c.JSON(401, updateUserInput)
		return
	}

	user, err := h.bookService.UpdateBookByID(id, updateUserInput)
	if err != nil {

		c.JSON(500, user)
		return
	}
	c.JSON(200, user)
}
