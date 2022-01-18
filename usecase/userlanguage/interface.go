package userlanguage

import "go-mysql-api/entity"

type Repository interface {
	UpdateOrCreateMany(manyUserLanguage []entity.User_Language) ([]entity.User_Language, error)
}
