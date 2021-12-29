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

func (r *GormRepository) Create(ent entity.User) (entity.User, error) {
	resultErr := r.db.Create(&ent)
	if resultErr != nil {
		return ent, resultErr.Error
	}
	return ent, nil
}

func (r *GormRepository) List() ([]entity.User, error) {
	var ents []entity.User
	result := r.db.Preload("Books").Find(&ents)
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

func (r *GormRepository) EmailExist(email string) bool {
	var ent entity.User
	if err := r.db.Where("email = ?", email).Limit(1).First(&ent); err.RowsAffected == 1 {
		return true
	}
	return false
}
