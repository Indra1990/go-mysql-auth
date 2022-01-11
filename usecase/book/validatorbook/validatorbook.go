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

func BookValidationUpdate(dto dto.BookUpdateRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Title, validation.Required.Error("Title is Required")),
		validation.Field(&dto.Description, validation.Required.Error("Description is Required")),
	)
}

func BookValidationUpdateMultiple(dto dto.BookUpdateMultipleRequest) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ID, validation.Required.Error("ID book is Required")),
		validation.Field(&dto.Title, validation.Required.Error("Title is Required")),
		validation.Field(&dto.Description, validation.Required.Error("Description is Required")),
	)
}
