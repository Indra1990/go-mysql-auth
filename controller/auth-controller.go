package controller

import (
	"go-mysql-api/dto"
	"go-mysql-api/usecase/auth"
	"net/http"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

type AuthControoller interface {
	Login(c *gin.Context)
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": validateLogin.Error(),
		})
		return
	}

	if checkEmail := checkmail.ValidateFormat(dto.Email); checkEmail != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": checkEmail.Error()})
		return
	}

	authUser, err := u.service.DoLogin(dto)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Email or Password does not match",
		})
		return
	}

	token, errToken := u.service.CreateToken(authUser.ID)
	if errToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errToken.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Login",
		"users":        authUser,
		"access_token": token,
		"time":         time.Now().Unix(),
	})
}
