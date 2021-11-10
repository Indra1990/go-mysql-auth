package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/usecase/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookControoller interface {
	CreateBook(c *gin.Context)
	GetBooks(c *gin.Context)
}

type bookControoller struct {
	service book.Service
}

func NewBookController(book book.Service) *bookControoller {
	return &bookControoller{book}
}

func (book bookControoller) GetBooks(ctx *gin.Context) {
	dtoBook, err := book.service.GetBookList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "List Book Error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "book list",
		"books":   dtoBook,
	})
}

func (book bookControoller) CreateBook(ctx *gin.Context) {
	var dto dto.BookCreateRequest

	err := dto.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			// "book":    dto,
		})
		return
	}
	ctx.Bind(&dto)
	book.service.CreateBook(dto)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "book created",
		"book":    dto,
	})
}
