package repoimplbook

import (
	"go-mysql-api/entity"

	"gorm.io/gorm"
)

type GormRepositoryBook struct {
	db *gorm.DB
}

func NewGormRepositoryBook(db *gorm.DB) *GormRepositoryBook {
	return &GormRepositoryBook{db}
}

func (r *GormRepositoryBook) Create(ent entity.Book) error {
	result := r.db.Create(&ent)
	return result.Error
}
