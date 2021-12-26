package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserCreateRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"-"`
}

func (user UserCreateRequest) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required.Error("is Required")),
		validation.Field(&user.Email, validation.Required.Error("is Required")),
	)
}

type GetAuthUserResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	ID    uint64         `json:"id"`
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Books []BookResponse `json:"books,omitempty"`
}

type BookResponse struct {
	ID          uint64 `json:"idbook"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type GetUserByIDRequest struct {
	ID uint64 `json:"id"`
}

type UserUpdateRequest struct {
	ID    uint64
	Name  string `form:"name"`
	Email string `form:"email"`
}

type GetAuthUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (user GetAuthUserRequest) ValidateAuthLogin() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required.Error("is Required"), is.Email),
		validation.Field(&user.Password, validation.Required.Error("is Required")),
	)
}
