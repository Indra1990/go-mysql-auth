package controller

import (
	"go-mysql-api/config"
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookControoller interface {
	CreateBook(c *gin.Context)
	GetBooks(c *gin.Context)
}

var (
	dbBook *gorm.DB = config.SetupDatabaseConnection()
)

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
	var entBook entity.Book

	ctx.Bind(&dto)
	err := dto.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	titleExist := dbBook.Where("title = ?", dto.Title).Limit(1).First(&entBook)
	if titleExist.RowsAffected == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Already Title Exist"})
		return
	}

	book.service.CreateBook(dto)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "book created",
		"book":    dto,
	})
}
