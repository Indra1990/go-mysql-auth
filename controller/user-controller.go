package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/usecase/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	FindByIdUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	service user.Service
}

func NewUserController(svc user.Service) UserController {
	return &userController{svc}
}

func (u userController) GetUser(ctx *gin.Context) {
	dto, err := u.service.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user all",
		"users":   dto,
	})
}

func (u userController) FindByIdUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	requestDTO := dto.GetUserByIDRequest{
		ID: uint64(id),
	}

	responseDTO, err := u.service.UserFindById(requestDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get data user by id",
		"data":    responseDTO,
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

func (u userController) UpdateUser(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var dto dto.UserUpdateRequest
	ctx.Bind(&dto)
	dto.ID = id
	u.service.UpdateUser(dto)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"data":    dto,
	})
}

func (u userController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var dto dto.GetUserByIDRequest
	dto.ID = id
	u.service.DeleteUser(dto)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
		// "data":    dto,
	})

}
