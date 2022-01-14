package servicelanguages

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/languages"
)

type Service struct {
	repo languages.Repository
}

func NewServiceLanguages(repo languages.Repository) *Service {
	return &Service{repo}
}

func (s Service) LanguaageList() ([]dto.GetLanguageListResponse, error) {
	result, err := s.repo.GetListLanguage()
	if err != nil {
		return nil, err
	}
	resultList := s.mapLanguageListEntityTODTO(result)
	return resultList, nil
}

func (s *Service) mapLanguageListEntityTODTO(ents []entity.Languages) []dto.GetLanguageListResponse {
	var languages []dto.GetLanguageListResponse
	for _, lang := range ents {
		usrLang := []dto.UserLanguages{}
		for _, usr := range lang.User {
			resultUsr := dto.UserLanguages{
				ID:    int(usr.ID),
				Name:  usr.Name,
				Email: usr.Email,
			}

			usrLang = append(usrLang, resultUsr)
		}
		resultlang := dto.GetLanguageListResponse{
			ID:   int(lang.ID),
			Name: lang.Name,
			User: usrLang,
		}

		languages = append(languages, resultlang)
	}
	return languages
}

func (s *Service) LanguageCreate(dto dto.CreateLanguageRequest) error {
	language := entity.Languages{}
	language.Name = dto.Name
	err := s.repo.Create(language)
	if err != nil {
		return err
	}
	return nil
}
