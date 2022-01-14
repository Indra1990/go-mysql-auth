package book

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	ListBook() ([]entity.Book, error)
	Create(ent entity.Book) (entity.Book, error)
	CreateMultiple(ents []entity.Book) error
	CheckTitleBook(title string) bool
	FindByIDBook(idbook int) (entity.Book, error)
	UpdateBook(ent entity.Book) error
	DeleteBook(ent entity.Book) error
	DeleteBookMultiple(ent []entity.Book, idbooks []uint) error
}

type Service interface {
	CreateBook(dto dto.BookCreateRequest, iduser int) error
	BookCreateMultipleRequest(dto []dto.BookCreateMultipleRequest, iduser int) error
	GetBookList() ([]dto.GetBookResponse, error)
	ExistTitleBook(title string) bool
	BookFindByID(idbook int) (dto.GetBookResponse, error)
	BookUpdated(idbook int, currentLogged int, dto dto.BookUpdateRequest) error
	BookUpdatedMultiple(dto []dto.BookUpdateMultipleRequest) error
	BookDelete(idbook int) error
	BookDeleteMultiple(dto []dto.BookDeleteMultiple) (bool, error)
}
