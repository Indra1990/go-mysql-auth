package serviceimplbook

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/book"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Service struct {
	repo book.Repository
}

func NewServiceBook(repo book.Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateBook(dto dto.BookCreateRequest) error {
	mapDTO := s.mapBookCreateRequestDTOtoEntity(dto)
	insertBook := s.repo.Create(mapDTO)
	if insertBook != nil {
		return insertBook

	}
	return nil
}

func (s *Service) mapBookCreateRequestDTOtoEntity(dto dto.BookCreateRequest) entity.Book {
	return entity.Book{
		Title:       dto.Title,
		Description: dto.Description,
		UserID:      uint(dto.UserID),
	}
}

func (s *Service) ValidateRequest(dto dto.BookCreateRequest) error {
	// var dto dto.BookCreateRequest
	err := validation.ValidateStruct(&dto,
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.Description, validation.Required),
	)

	if err != nil {
		return err
	}

	return nil
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
