package validatorbook

import (
	"go-mysql-api/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func BookValidation(dto dto.BookCreateRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Title, validation.Required.Error("Title is Required")),
		validation.Field(&dto.Description, validation.Required.Error("Description is Required")),
	)
}
