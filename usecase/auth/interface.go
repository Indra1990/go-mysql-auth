package auth

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

// for function db gorm
type Repository interface {
	AuthLogin(email string, password string) (entity.User, error)
}

// for function to controller
type Service interface {
	DoLogin(dto dto.GetAuthUserRequest) (dto.GetUserResponse, error)
	CreateToken(userid uint64) (string, error)
}
