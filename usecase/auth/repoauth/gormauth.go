package repoauth

import (
	"go-mysql-api/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthGormRepository struct {
	db *gorm.DB
}

func NewAuthGormRepository(db *gorm.DB) *AuthGormRepository {
	return &AuthGormRepository{db}
}

func (auth AuthGormRepository) AuthLogin(email string, password string) (entity.User, error) {
	var usr entity.User
	result := auth.db.Where("email = ?", email).First(&usr)
	if result.RowsAffected == 1 {
		if checkHash := checkPassword(usr.Password, password); checkHash != nil {
			return usr, checkHash
		}
	}
	return usr, result.Error
}

func checkPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
