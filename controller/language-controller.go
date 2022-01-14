package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/languages/servicelanguages"
	"go-mysql-api/usecase/languages/validatorlanguage.go"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LanguagesController interface {
	CreateLang(c *gin.Context)
}

type languagesController struct {
	lang servicelanguages.Service
}

func NewLanguagesController(lang servicelanguages.Service) *languagesController {
	return &languagesController{lang}
}

func (c *languagesController) GetlistLang(ctx *gin.Context) {
	result, err := c.lang.LanguaageList()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *languagesController) CreateLang(ctx *gin.Context) {
	var dtoCreateLang dto.CreateLanguageRequest
	ctx.Bind(&dtoCreateLang)

	if valErr := validatorlanguage.LangValidateCreate(dtoCreateLang); valErr != nil {
		resErr := helper.APIResponse("create language failed", http.StatusUnprocessableEntity, "error", valErr)
		ctx.JSON(http.StatusUnprocessableEntity, resErr)
		return
	}

	saveErr := c.lang.LanguageCreate(dtoCreateLang)
	if saveErr != nil {
		resErr := helper.APIResponse("create language failed", http.StatusInternalServerError, "error", saveErr.Error())
		ctx.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res := helper.APIResponse("create language success", http.StatusOK, "success", dtoCreateLang)
	ctx.JSON(http.StatusOK, res)
}
