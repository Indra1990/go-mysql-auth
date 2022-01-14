package languages

import (
	"go-mysql-api/dto"
	"go-mysql-api/entity"
)

type Repository interface {
	GetListLanguage() ([]entity.Languages, error)
	Create(ent entity.Languages) error
}

type Service interface {
	LanguaageList() ([]dto.GetLanguageListResponse, error)
	LanguageCreate(dto dto.CreateLanguageRequest) error
}
