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
func (s *Service) GetUserList() ([]dto.GetUserResponse, error) {
	UserList, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	userDto, err := s.mapUserEntitiesToGetResponseDTOs(UserList)
	return userDto, err
}

// map get list user entity to dto
func (s *Service) mapUserEntitiesToGetResponseDTOs(ents []entity.User) ([]dto.GetUserResponse, error) {
	result := []dto.GetUserResponse{}
	for _, usr := range ents {
		listUser := s.mapUserEntityToGetUserByIDDTO(usr)
		result = append(result, listUser)
	}

	return result, nil
}

// find id user from db
func (s *Service) UserFindById(dto dto.GetUserByIDRequest) (dto.GetUserResponse, error) {
	userId, err := s.repo.FindById(dto.ID)
	usr := s.mapUserEntityToGetUserByIDDTO(userId)
	return usr, err
}

func (s *Service) mapUserEntityToGetUserByIDDTO(ent entity.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:    ent.ID,
		Name:  ent.Name,
		Email: ent.Email,
	}
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

// map dto to entity create user
func (s *Service) mapUserCreateRequestDTOtoEntity(dto dto.UserCreateRequest) entity.User {
	return entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (s *Service) UpdateUser(dto dto.UserUpdateRequest) error {
	ent := s.mapUserUpdateRequestDTOtoEntity(dto)
	err := s.repo.Update(ent)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) mapUserUpdateRequestDTOtoEntity(dto dto.UserUpdateRequest) entity.User {
	return entity.User{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
	}
}

func (s *Service) DeleteUser(dto dto.GetUserByIDRequest) error {
	err := s.repo.Delete(dto.ID)
	return err
}
