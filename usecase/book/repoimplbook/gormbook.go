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

func (r *GormRepositoryBook) ListBook() ([]entity.Book, error) {
	var ents []entity.Book
	result := r.db.Preload("User").Find(&ents)
	return ents, result.Error
}
