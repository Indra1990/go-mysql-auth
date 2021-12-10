package user

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	Create(ent entity.User) error
	List() ([]entity.User, error)
	FindById(id uint64) (entity.User, error)
	Update(ent entity.User) error
	Delete(id uint64) error
	EmailExist(email string) bool
}

type Service interface {
	GetUserList() ([]dto.GetUserResponse, error)
	UserFindById(dto dto.GetUserByIDRequest) (dto.GetUserResponse, error)
	CreateUser(dto dto.UserCreateRequest) error
	UpdateUser(dto dto.UserUpdateRequest) error
	DeleteUser(dto dto.GetUserByIDRequest) error
	CheckEmailExist(email string) bool
}
