package serviceimplbook

import (
	"errors"
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

func (s *Service) BookCreateMultipleRequest(dto []dto.BookCreateMultipleRequest, iduser int) error {
	books := []entity.Book{}
	for _, dtobook := range dto {
		result := entity.Book{
			Title:       dtobook.Title,
			Description: dtobook.Description,
			Slug:        slug.Make(dtobook.Title),
			UserID:      uint(iduser),
		}
		books = append(books, result)
	}

	err := s.repo.CreateMultiple(books)
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

func (s *Service) BookFindByID(idbook int) (dto.GetBookResponse, error) {
	var dto dto.GetBookResponse
	book, err := s.repo.FindByIDBook(idbook)
	dtoBook := s.mapBookFindIDentitytoDTO(book)
	if err != nil {
		return dto, err
	}
	return dtoBook, nil
}

func (s *Service) BookUpdated(idbook int, currentLogged int, dto dto.BookUpdateRequest) error {
	updateBook, err := s.repo.FindByIDBook(int(idbook))
	if err != nil {
		return err
	}

	if uint64(updateBook.UserID) != uint64(currentLogged) {
		return errors.New("permission denied")
	}
	updateBook.ID = uint64(idbook)
	updateBook.Title = dto.Title
	updateBook.Description = dto.Description
	updateBook.Slug = slug.Make(dto.Title)
	updatedErr := s.repo.UpdateBook(updateBook)
	if updatedErr != nil {
		return updatedErr
	}
	return nil
}

func (s *Service) BookUpdatedMultiple(dto []dto.BookUpdateMultipleRequest) error {
	for _, updRequest := range dto {
		updMultiple, err := s.repo.FindByIDBook(updRequest.ID)
		if err != nil {
			return err
		}
		updMultiple.Title = updRequest.Title
		updMultiple.Description = updRequest.Description
		updMultiple.Slug = slug.Make(updRequest.Title)
		updMultiple.UserID = uint(updRequest.UserID)
		updatedBookErr := s.repo.UpdateBook(updMultiple)
		if updatedBookErr != nil {
			return updatedBookErr
		}
	}
	return nil
}

func (s *Service) BookDelete(idbook int) error {
	book, err := s.repo.FindByIDBook(idbook)
	if err != nil {
		return err
	}
	delBook := s.repo.DeleteBook(book)
	if delBook != nil {
		return delBook
	}
	return nil
}

func (s *Service) BookDeleteMultiple(dto []dto.BookDeleteMultiple) (bool, error) {
	for _, del := range dto {

		findBook, err := s.repo.FindByIDBook(int(del.ID))
		if err != nil {
			return false, err
		}

		resultErr := s.repo.DeleteBook(findBook)
		if resultErr != nil {
			return false, resultErr
		}
	}

	return true, nil
}

func (s *Service) mapBookFindIDentitytoDTO(ent entity.Book) dto.GetBookResponse {
	return dto.GetBookResponse{
		ID:          ent.ID,
		Title:       ent.Title,
		Description: ent.Description,
		Slug:        ent.Slug,
		UserID:      int(ent.ID),
		User: dto.User{
			ID:    ent.User.ID,
			Name:  ent.User.Name,
			Email: ent.User.Email,
		},
	}
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
