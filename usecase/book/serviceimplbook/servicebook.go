package serviceimplbook

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/book"
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
