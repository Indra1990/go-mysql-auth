package user

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	Create(ent entity.User) error
	List() ([]entity.User, error)
}

type Service interface {
	CreateUser(dto dto.UserCreateRequest) error
	GetUserList() ([]entity.User, error)
}
