package serviceimpl

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/user"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	// "github.com/go-playground/validator"
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

		resultbook := []dto.BookResponse{}
		for _, bk := range usr.Books {
			listBook := dto.BookResponse{
				ID:          bk.ID,
				Title:       bk.Title,
				Description: bk.Description,
				UserID:      bk.UserID,
			}
			resultbook = append(resultbook, listBook)
		}

		listUser := dto.GetUserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
			Books: resultbook,
		}
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

// check email already exist
func (s *Service) CheckEmailExist(email string) bool {
	errEmail := s.repo.EmailExist(email)
	if errEmail {
		return errEmail
	}
	return false
}

// map dto to entity create user
func (s *Service) mapUserCreateRequestDTOtoEntity(dto dto.UserCreateRequest) entity.User {
	hash, _ := s.hashPassword(dto.Password)
	return entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hash,
	}
}

// hash password user
func (s *Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
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

func (auth *Service) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
