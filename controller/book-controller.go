package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/usecase/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookControoller interface {
	CreateBook(c *gin.Context)
}

type bookControoller struct {
	service book.Service
}

func NewBookController(book book.Service) *bookControoller {
	return &bookControoller{book}
}

func (book bookControoller) CreateBook(ctx *gin.Context) {
	var dto dto.BookCreateRequest

	ctx.Bind(&dto)
	book.service.CreateBook(dto)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "book created",
		"book":    dto,
	})
}