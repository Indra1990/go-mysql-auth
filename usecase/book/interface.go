package book

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	Create(ent entity.Book) error
}

type Service interface {
	CreateBook(dto dto.BookCreateRequest) error
}
