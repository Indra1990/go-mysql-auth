package repoauth

import (
	"go-mysql-api/entity"

	"gorm.io/gorm"
)

type AuthGormRepository struct {
	db *gorm.DB
}

func NewAuthGormRepository(db *gorm.DB) *AuthGormRepository {
	return &AuthGormRepository{db}
}
func (auth AuthGormRepository) UserRegister(ent entity.User) (entity.User, error) {
	err := auth.db.Create(&ent).Error
	if err != nil {
		return ent, err
	}
	return ent, err
}

func (auth AuthGormRepository) FindByID(email string) (entity.User, error) {
	var usr entity.User
	err := auth.db.Where("email", email).Find(&usr).Error
	if err != nil {
		return usr, err
	}

	return usr, nil
}
