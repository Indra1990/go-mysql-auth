package book

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	ListBook() ([]entity.Book, error)
	Create(ent entity.Book) (entity.Book, error)
	CheckTitleBook(title string) bool
}

type Service interface {
	CreateBook(dto dto.BookCreateRequest, iduser int) error
	GetBookList() ([]dto.GetBookResponse, error)
	ExistTitleBook(title string) bool
}
