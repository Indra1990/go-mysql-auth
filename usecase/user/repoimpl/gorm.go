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
	result := r.db.Preload("Book").Find(&ents)
	return ents, result.Error
}

func (r *GormRepository) FindById(id uint64) (entity.User, error) {
	var ent entity.User
	result := r.db.First(&ent, id)
	return ent, result.Error
}

func (r *GormRepository) Update(ent entity.User) error {
	result := r.db.Model(&ent).Updates(ent)
	return result.Error
}

func (r *GormRepository) Delete(id uint64) error {
	var ent entity.User
	result := r.db.Delete(&ent, id)
	return result.Error
}
