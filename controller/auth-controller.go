package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControoller interface {
	Login(c *gin.Context)
	CekToken(c *gin.Context)
	// GetBooks(c *gin.Context)
}

type authControoller struct {
	service auth.Service
}

func NewAuthController(auth auth.Service) *authControoller {
	return &authControoller{auth}
}

func (u authControoller) Login(ctx *gin.Context) {
	var dto dto.GetAuthUserRequest
	ctx.Bind(&dto)

	if validateLogin := dto.ValidateAuthLogin(); validateLogin != nil {
		res := helper.APIResponse("get user failed", http.StatusUnprocessableEntity, "error", validateLogin)

		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	authUser, err := u.service.DoLogin(dto)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Email or Password does not match",
		})
		return
	}

	// return
	token, errToken := u.service.CreateToken(authUser.ID)
	if errToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errToken.Error(),
		})
		return
	}
	// tkn := u.service.ExtractToken(ctx.Request)
	res := helper.APIResponse("login success", http.StatusOK, "success", authUser)

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Login",
		"access_token": token,
		"users":        res,
	})
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
