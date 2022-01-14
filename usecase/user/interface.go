package user

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"net/http"
)

type Repository interface {
	Create(ent entity.User) (entity.User, error)
	List() ([]entity.User, error)
	FindById(id uint64) (entity.User, error)
	Update(ent entity.User) error
	Delete(id uint64) error
	EmailExist(email string) bool
	FindIDUserLanguage(id int) (entity.Languages, error)
}

type Service interface {
	GetUserList() ([]dto.GetUserResponse, error)
	UserFindById(dto dto.GetUserByIDRequest) (dto.GetUserResponse, error)
	CreateUser(dto dto.UserCreateRequest) (dto.GetUserResponse, error)
	UpdateUser(dto dto.UserUpdateRequest) error
	DeleteUser(dto dto.GetUserByIDRequest) error
	CheckEmailExist(email string) bool
	ExtractToken(r *http.Request) string
	UserLanguageFindByID(id int) error
}
