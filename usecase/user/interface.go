package user

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	Create(ent entity.User) error
	List() ([]entity.User, error)
	FindById(id uint64) (entity.User, error)
}

type Service interface {
	CreateUser(dto dto.UserCreateRequest) error
	GetUserList() ([]dto.GetUserResponse, error)
	UserFindById(dto dto.GetUserByIDRequest) (dto.GetUserResponse, error)
}
