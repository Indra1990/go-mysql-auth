package auth

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"

	"github.com/dgrijalva/jwt-go"
)

// for function db gorm
type Repository interface {
	FindByID(email string) (entity.User, error)
	UserRegister(ent entity.User) (entity.User, error)
}

// for function to controller
type Service interface {
	DoLogin(dto dto.GetAuthUserRequest) (dto.GetUserResponse, error)
	RegisterUserInput(dto dto.UserCreateRequest) (dto.GetUserResponse, error)
	CreateToken(userid uint64) (string, error)
	ValidateToken(tokenEncoded string) (*jwt.Token, error)
}
