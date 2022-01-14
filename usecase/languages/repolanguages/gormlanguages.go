package repolanguages

import (
	"go-mysql-api/entity"

	"gorm.io/gorm"
)

type GormRepositoryLanguages struct {
	db *gorm.DB
}

func NewGormRepositoryLanguages(db *gorm.DB) *GormRepositoryLanguages {
	return &GormRepositoryLanguages{db}
}

func (r *GormRepositoryLanguages) GetListLanguage() ([]entity.Languages, error) {
	var ents []entity.Languages
	err := r.db.Preload("User").Find(&ents).Error
	if err != nil {
		return ents, err
	}
	return ents, nil
}

func (r *GormRepositoryLanguages) Create(ent entity.Languages) error {
	createLanguage := r.db.Create(&ent)
	return createLanguage.Error
}
