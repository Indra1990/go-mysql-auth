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
	result := r.db.Preload("Books").Preload("Languages").Find(&ents)
	return ents, result.Error
}

func (r *GormRepository) FindById(id uint64) (entity.User, error) {
	var ent entity.User
	result := r.db.Model(&ent).Where("id = ? ", id).Preload("Languages").Find(&ent)
	return ent, result.Error
}

func (r *GormRepository) Update(ent entity.User) (entity.User, error) {
	result := r.db.Save(&ent).Error
	if result != nil {
		return ent, result
	}
	return ent, nil
}

func (r *GormRepository) Delete(id uint64) error {
	var ent entity.User
	dataDelte, err := r.FindById(id)
	if err != nil {
		return err
	}
	// delete assosiation
	errAsso := r.db.Model(&dataDelte).Association("Languages").Delete(dataDelte.Languages)
	if err != nil {
		return errAsso
	}
	// delete user find
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

func (r *GormRepository) FindIDUserLanguage(id int) (entity.Languages, error) {
	var ent entity.Languages
	err := r.db.Where("id", id).Find(&ent).Error
	if err != nil {
		return ent, err
	}
	return ent, nil
}

func (r *GormRepository) CheckManyUserLanguage(iduser int, idlanguage int) bool {
	var userLang entity.User_Language
	if row := r.db.Where("user_id = ? AND languages_id = ?", iduser, idlanguage).First(&userLang); row.RowsAffected == 0 {
		return true
	}
	return false
}

func (r *GormRepository) DeleteUserLanguages(iduser uint, idlanguage uint) error {
	var userLang entity.User_Language
	err := r.db.Where("user_id = ? AND languages_id = ?", iduser, idlanguage).Delete(&userLang)
	return err.Error
}
