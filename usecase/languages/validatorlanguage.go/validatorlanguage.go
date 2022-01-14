package validatorlanguage

import (
	"go-mysql-api/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func LangValidateCreate(dto dto.CreateLanguageRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required.Error("Name is Required")),
	)

}
