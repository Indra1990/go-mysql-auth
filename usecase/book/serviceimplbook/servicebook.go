package serviceimplbook

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/book"

	"github.com/gosimple/slug"
)

type Service struct {
	repo book.Repository
}

func NewServiceBook(repo book.Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateBook(dto dto.BookCreateRequest, iduser int) error {
	book := entity.Book{}
	book.Title = dto.Title
	book.Description = dto.Description
	book.Slug = slug.Make(dto.Title)
	book.UserID = uint(iduser)
	book, err := s.repo.Create(book)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ExistTitleBook(title string) bool {
	errBool := s.repo.CheckTitleBook(title)
	return errBool
}

func (s *Service) GetBookList() ([]dto.GetBookResponse, error) {
	bookList, err := s.repo.ListBook()
	if err != nil {
		return nil, err
	}
	bookDto, err := s.mapBookGetResponseEntityToDTO(bookList)
	return bookDto, err
}

func (s *Service) mapBookGetResponseEntityToDTO(ents []entity.Book) ([]dto.GetBookResponse, error) {
	result := []dto.GetBookResponse{}
	for _, book := range ents {
		dataBook := dto.GetBookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Slug:        book.Slug,
			Description: book.Description,
			User: dto.User{
				ID:    book.User.ID,
				Name:  book.User.Name,
				Email: book.User.Email,
			},
		}

		result = append(result, dataBook)
	}

	return result, nil
}
