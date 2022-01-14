package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/user"
	"go-mysql-api/usecase/user/validatorimpl"
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
		resErr := helper.APIResponse("user all failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	if err != nil {
		resErr := helper.APIResponse("user all failed", http.StatusInternalServerError, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	ctx.Bind(&dto)
	res := helper.APIResponse("get user all", http.StatusOK, "success", dto)

	ctx.JSON(http.StatusOK, res)
}

func (u userController) FindByIdUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		resErr := helper.APIResponse("user detail failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	requestDTO := dto.GetUserByIDRequest{
		ID: uint64(id),
	}

	responseDTO, errFindUser := u.service.UserFindById(requestDTO)
	if errFindUser != nil {
		resErr := helper.APIResponse("user detail failed", http.StatusBadRequest, "error", errFindUser.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}
	res := helper.APIResponse("user detail", http.StatusOK, "success", responseDTO)
	ctx.JSON(http.StatusOK, res)
}

func (u userController) CreateUser(ctx *gin.Context) {
	var dto dto.UserCreateRequest
	ctx.Bind(&dto)

	if err := validatorimpl.ValidationCreateUser(dto); err != nil {
		errValidate := helper.APIResponse("user failed create", http.StatusUnprocessableEntity, "error", err)
		ctx.JSON(http.StatusUnprocessableEntity, errValidate)
		return
	}

	if validEmailExist := u.service.CheckEmailExist(dto.Email); validEmailExist {
		errValidate := helper.APIResponse("user failed create", http.StatusUnprocessableEntity, "error", gin.H{"email": "Already email exist"})
		ctx.JSON(http.StatusUnprocessableEntity, errValidate)
		return
	}

	for _, lang := range dto.LanguageMany {
		errLang := u.service.UserLanguageFindByID(int(lang.ID))
		if errLang != nil {
			errValidateLang := helper.APIResponse("user failed create", http.StatusUnprocessableEntity, "error", gin.H{"error": errLang.Error() + " id: " + strconv.FormatUint(uint64(lang.ID), 10)})
			ctx.JSON(http.StatusUnprocessableEntity, errValidateLang)
			return
		}
	}

	userCreated, userCreatedErr := u.service.CreateUser(dto)
	if userCreatedErr != nil {
		errValidate := helper.APIResponse("user failed create", http.StatusUnprocessableEntity, "error", userCreatedErr)
		ctx.JSON(http.StatusUnprocessableEntity, errValidate)
		return
	}

	res := helper.APIResponse("created user", http.StatusOK, "success", userCreated)
	ctx.JSON(http.StatusOK, res)
}

func (u userController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		resErr := helper.APIResponse("user update failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	var dto dto.UserUpdateRequest
	ctx.Bind(&dto)
	dto.ID = id
	u.service.UpdateUser(dto)

	res := helper.APIResponse("updated user", http.StatusOK, "success", dto)
	ctx.JSON(http.StatusOK, res)
}

func (u userController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		resErr := helper.APIResponse("user delete failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	var dto dto.GetUserByIDRequest
	dto.ID = id
	u.service.DeleteUser(dto)

	res := helper.APIResponse("updated user", http.StatusOK, "success", gin.H{"delted": "successfully deleted"})
	ctx.JSON(http.StatusOK, res)
}
