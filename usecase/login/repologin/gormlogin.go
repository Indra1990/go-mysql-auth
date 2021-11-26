package repologin

import "gorm.io/gorm"

type GormRepositoryBook struct {
	db *gorm.DB
}

func NewGormRepositoryBook(db *gorm.DB) *GormRepositoryBook {
	return &GormRepositoryBook{db}
}

// func (g *GormRepositoryBook) LoginRepo(email string, password string) error {

// }
