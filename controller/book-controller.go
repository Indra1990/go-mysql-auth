package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/book"
	"go-mysql-api/usecase/book/validatorbook"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
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
		resErr := helper.APIResponse("Get book failed", http.StatusInternalServerError, "error", err)
		ctx.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res := helper.APIResponse("Get book success", http.StatusOK, "success", dtoBook)
	ctx.JSON(http.StatusOK, res)
}

func (book bookControoller) CreateBook(ctx *gin.Context) {
	var bookRequest dto.BookCreateRequest
	ctx.Bind(&bookRequest)

	err := validatorbook.BookValidation(bookRequest)
	if err != nil {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "error", err)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	if titleExist := book.service.ExistTitleBook(bookRequest.Title); !titleExist {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "errro", gin.H{"title": "title already exist"})
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	user := ctx.MustGet("currentUser").(dto.GetUserResponse)
	errBook := book.service.CreateBook(bookRequest, int(user.ID))
	if errBook != nil {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "error", errBook)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	res := helper.APIResponse("user detail", http.StatusOK, "success", bookRequest)
	ctx.JSON(http.StatusOK, res)
}

func (book bookControoller) CreateBookMultiple(ctx *gin.Context) {
	var bookRequestMultiple []dto.BookCreateMultipleRequest
	err := ctx.Bind(&bookRequestMultiple)
	validation.Validate(bookRequestMultiple)
	if err != nil {
		resErr := helper.APIResponse("create book failed", http.StatusInternalServerError, "error", err)
		ctx.JSON(http.StatusInternalServerError, resErr)
		return
	}

	user := ctx.MustGet("currentUser").(dto.GetUserResponse)
	errReqMulti := book.service.BookCreateMultipleRequest(bookRequestMultiple, int(user.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errReqMulti)
		return
	}

	res := helper.APIResponse("create book success", http.StatusOK, "success", bookRequestMultiple)
	ctx.JSON(http.StatusOK, res)

}

func (book bookControoller) FindByIdBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resErr := helper.APIResponse("Get Detail book failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	res, errBook := book.service.BookFindByID(id)
	if errBook != nil {
		resErr := helper.APIResponse("Get Detail book failed", http.StatusBadRequest, "error", errBook.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	response := helper.APIResponse("user detail", http.StatusOK, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (book *bookControoller) UpdateBook(ctx *gin.Context) {
	var dto dto.BookUpdateRequest
	ctx.Bind(&dto)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resErr := helper.APIResponse("Update book failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	errValidator := validatorbook.BookValidationUpdate(dto)
	if errValidator != nil {
		resErr := helper.APIResponse("create book failed", http.StatusUnprocessableEntity, "error", errValidator)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	resultErr := book.service.BookUpdated(int(id), dto)
	if resultErr != nil {
		resErr := helper.APIResponse("Update book failed", http.StatusBadRequest, "error", resultErr.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (book *bookControoller) DeleteBook(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resErr := helper.APIResponse("Delete book failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}
	book.service.BookDelete(int(id))

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})

}
