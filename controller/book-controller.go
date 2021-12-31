package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/book"
	"go-mysql-api/usecase/book/validatorbook"
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

	ctx.Bind(&dto)
	err := validatorbook.BookValidation(dto)
	if err != nil {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "error", err)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	if titleExist := book.service.ExistTitleBook(dto.Title); !titleExist {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "errro", gin.H{"title": "title already exist"})
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	errBook := book.service.CreateBook(dto, ctx.MustGet("iduser").(int))
	if errBook != nil {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "error", errBook)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	res := helper.APIResponse("user detail", http.StatusOK, "success", dto)
	ctx.JSON(http.StatusOK, res)
}
