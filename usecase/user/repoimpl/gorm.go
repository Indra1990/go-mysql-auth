package repoimpl

import (
	"go-mysql-api/entity"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db}
}

func (r *GormRepository) Create(ent entity.User) error {
	result := r.db.Create(&ent)
	return result.Error
}

func (r *GormRepository) List() ([]entity.User, error) {
	var ents []entity.User
	result := r.db.Find(&ents)
	return ents, result.Error
}
