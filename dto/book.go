package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BookCreateRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Slug        string `form:"slug" json:"-"`
	UserID      int64  `form:"user_id"`
}

func (req BookCreateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required.Error("Title is Required")),
		validation.Field(&req.Description, validation.Required.Error("Description is Required")),
	)
}

type GetBookResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	User        User   `json:"user"`
}

type User struct {
	ID    uint64 `json:"iduser"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
