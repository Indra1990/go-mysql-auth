package repouserlanguage

import (
	"go-mysql-api/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db}
}

func (g GormRepository) UpdateOrCreateMany(manyUserLanguage []entity.User_Language) ([]entity.User_Language, error) {
	result := g.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"languages_id"}),
	}).Create(&manyUserLanguage)

	return manyUserLanguage, result.Error
}
