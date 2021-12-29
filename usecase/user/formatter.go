package user

import (
	"go-mysql-api/dto"
)

type UserFormaterLoginRegister struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUserLoginRegister(dto dto.GetUserResponse, token string) UserFormaterLoginRegister {
	format := UserFormaterLoginRegister{
		ID:    uint64(dto.ID),
		Name:  dto.Name,
		Email: dto.Email,
		Token: token,
	}

	return format
}
