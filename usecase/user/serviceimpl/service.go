package serviceimpl

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/user"
)

type Service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *Service {
	return &Service{repo}
}

// from db get list user
func (s *Service) GetUserList() ([]entity.User, error) {
	ents, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	return ents, nil
}

// save to db
func (s *Service) CreateUser(dto dto.UserCreateRequest) error {
	ent := s.mapUserCreateRequestDTOtoEntity(dto)
	err := s.repo.Create(ent)
	if err != nil {
		return err
	}

	return nil
}

// map dto to entity
func (s *Service) mapUserCreateRequestDTOtoEntity(dto dto.UserCreateRequest) entity.User {
	return entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
