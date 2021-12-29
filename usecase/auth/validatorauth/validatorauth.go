package validatorauth

import (
	"go-mysql-api/dto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateRegister(dto dto.UserCreateRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required.Error("is Required")),
		validation.Field(&dto.Email, validation.Required.Error("is Required"), is.Email),
		validation.Field(&dto.Password, validation.Required.Error("is Required")),
	)
}
