package serviceimpl

import (
	"errors"
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

		resultLanguage := []dto.UserLanguageResponse{}
		for _, lang := range usr.Languages {
			listLanguage := dto.UserLanguageResponse{
				ID:   lang.ID,
				Name: lang.Name,
			}
			resultLanguage = append(resultLanguage, listLanguage)
		}

		listUser := dto.GetUserResponse{
			ID:        usr.ID,
			Name:      usr.Name,
			Email:     usr.Email,
			Books:     resultbook,
			Languages: resultLanguage,
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
	resultBook := []dto.BookResponse{}
	for _, valueBook := range ent.Books {
		result := dto.BookResponse{
			ID:          valueBook.ID,
			Title:       valueBook.Title,
			Description: valueBook.Description,
			UserID:      valueBook.UserID,
		}

		resultBook = append(resultBook, result)
	}
	return dto.GetUserResponse{
		ID:    ent.ID,
		Name:  ent.Name,
		Email: ent.Email,
		Books: resultBook,
	}
}

// save to db
func (s *Service) CreateUser(dto dto.UserCreateRequest) (dto.GetUserResponse, error) {
	arr := []entity.Languages{}
	for _, val := range dto.LanguageMany {
		result := entity.Languages{
			ID: val.ID,
		}
		arr = append(arr, result)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)
	userCreate := entity.User{
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  string(password),
		Languages: arr,
	}

	user, err := s.repo.Create(userCreate)
	ent, errMap := s.mapUserCreateEntityTODTO(user)
	if err != nil {
		return ent, errMap
	}
	return ent, nil
}

func (s *Service) UserLanguageFindByID(id int) error {
	data, errlang := s.repo.FindIDUserLanguage(id)
	if errlang != nil {
		return errlang
	}
	if data.ID == 0 {
		return errors.New("not found language")
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
func (s *Service) mapUserCreateEntityTODTO(ent entity.User) (dto.GetUserResponse, error) {
	var strSlice = []dto.UserLanguageResponse{}
	for _, lang := range ent.Languages {
		findLanguage, _ := s.repo.FindIDUserLanguage(int(lang.ID))

		result := dto.UserLanguageResponse{
			ID:   lang.ID,
			Name: findLanguage.Name,
		}
		strSlice = append(strSlice, result)
	}
	return dto.GetUserResponse{
		ID:        ent.ID,
		Name:      ent.Name,
		Email:     ent.Email,
		Languages: strSlice,
	}, nil

}

func (s *Service) UpdateUser(dto dto.UserUpdateRequest, id int64) (dto.GetUserResponse, error) {
	arr := []entity.Languages{}
	for _, lang := range dto.LanguageMany {
		if checkMany := s.repo.CheckManyUserLanguage(int(id), int(lang.ID)); checkMany {
			resultMany := entity.Languages{
				ID: uint(lang.ID),
			}
			arr = append(arr, resultMany)
		}

	}

	usr, err := s.repo.FindById(uint64(id))
	dtoUser := s.mapUserEntityToGetUserByIDDTO(usr)
	if err != nil {
		return dtoUser, err
	}
	usr.Name = dto.Name
	usr.Email = dto.Email
	usr.Languages = arr
	dtoUpdate, dtoUpdateErr := s.repo.Update(usr)
	if dtoUpdateErr != nil {
		return dtoUser, dtoUpdateErr
	}
	ent := s.mapUserUpdateRequestDTOtoEntity(dtoUpdate)
	return ent, nil

}

func (s *Service) mapUserUpdateRequestDTOtoEntity(ent entity.User) dto.GetUserResponse {
	var dtoLanguages []dto.UserLanguageResponse
	for _, lang := range ent.Languages {
		langFind, _ := s.repo.FindIDUserLanguage(int(lang.ID))
		resultLang := dto.UserLanguageResponse{
			ID:   lang.ID,
			Name: langFind.Name,
		}
		dtoLanguages = append(dtoLanguages, resultLang)
	}
	return dto.GetUserResponse{
		ID:        ent.ID,
		Name:      ent.Name,
		Email:     ent.Email,
		Languages: dtoLanguages,
	}
}

func (s *Service) DeleteUser(dto dto.GetUserByIDRequest) error {
	err := s.repo.Delete(dto.ID)
	return err
}

func (auth *Service) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
