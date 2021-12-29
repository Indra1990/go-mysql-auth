package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/auth"
	"go-mysql-api/usecase/auth/validatorauth"
	"go-mysql-api/usecase/user"
	"go-mysql-api/usecase/user/validatorimpl"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControoller interface {
	Login(c *gin.Context)
	CekToken(c *gin.Context)
	Register(c *gin.Context)
}

type authControoller struct {
	service auth.Service
}

func NewAuthController(auth auth.Service) *authControoller {
	return &authControoller{auth}
}

func (u authControoller) Register(ctx *gin.Context) {
	var dto dto.UserCreateRequest
	err := ctx.Bind(&dto)

	if err != nil {
		res := helper.APIResponse("register failed", http.StatusBadRequest, "error", err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if errvalidate := validatorauth.ValidateRegister(dto); errvalidate != nil {
		res := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errvalidate)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	usr, errRegister := u.service.RegisterUserInput(dto)
	if errRegister != nil {
		res := helper.APIResponse("register failed", http.StatusBadRequest, "error", errRegister)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, errToken := u.service.CreateToken(usr.ID)
	if errToken != nil {
		res := helper.APIResponse("login failed generate token", http.StatusBadRequest, "error", gin.H{"error": errToken.Error()})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	formater := user.FormatUserLoginRegister(usr, token)
	res := helper.APIResponse("login success", http.StatusOK, "success", formater)

	ctx.JSON(http.StatusOK, res)
}

func (u authControoller) Login(ctx *gin.Context) {
	var dto dto.GetAuthUserRequest

	ctx.Bind(&dto)

	if validateLogin := validatorimpl.ValidateAuthLogin(dto); validateLogin != nil {
		res := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", validateLogin)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	authUser, err := u.service.DoLogin(dto)
	if err != nil {
		res := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, errToken := u.service.CreateToken(authUser.ID)
	if errToken != nil {
		res := helper.APIResponse("login failed generate token", http.StatusBadRequest, "error", gin.H{"error": errToken.Error()})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	formater := user.FormatUserLoginRegister(authUser, token)
	res := helper.APIResponse("login success", http.StatusOK, "success", formater)

	ctx.JSON(http.StatusOK, res)
}

func (u authControoller) CekToken(ctx *gin.Context) {
	tkn, err := u.service.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Mzk1NTIzNjQsInVzZXJfaWQiOjUzfQ.TLsi50GjrgwENGUIfOjGbR2dqb-JrFBFpXj1DAOl9TA")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Invalid",
		})
		return
	}

	if tkn.Valid {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "TOKEN VALID",
		})
		return
	}
}
