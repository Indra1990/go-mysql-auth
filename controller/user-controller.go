package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userController struct {
	service user.Service
}

func NewUserController(svc user.Service) UserController {
	return &userController{svc}
}

func (u userController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user all",
	})
}

func (u userController) CreateUser(ctx *gin.Context) {
	var dto dto.UserCreateRequest
	ctx.Bind(&dto)
	u.service.CreateUser(dto)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"data":    dto,
	})
}
