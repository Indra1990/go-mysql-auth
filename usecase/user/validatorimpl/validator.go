package validatorimpl

import (
	"go-mysql-api/dto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateAuthLogin(dto dto.GetAuthUserRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Email, validation.Required.Error("is Required"), is.Email),
		validation.Field(&dto.Password, validation.Required.Error("is Required")),
	)
}

func ValidationCreateUser(dto dto.UserCreateRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required.Error("is Required")),
		validation.Field(&dto.Email, validation.Required.Error("is Required"), is.Email),
		validation.Field(&dto.Password, validation.Required.Error("is Required")),
	)
}

func ValidationUpdateUser(dto dto.UserUpdateRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required.Error("is Required")),
		validation.Field(&dto.Email, validation.Required.Error("is Required"), is.Email),
	)
}

func ValidationUserLanguage(userLang []dto.LanguageMany) error {
	err := validation.Validate(userLang)
	if err != nil {
		return err
	}
	return nil
}
