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

func (r *GormRepositoryBook) Create(ent entity.Book) (entity.Book, error) {
	resultErr := r.db.Create(&ent)
	if resultErr != nil {
		return ent, resultErr.Error
	}
	return ent, nil
}

func (r *GormRepositoryBook) ListBook() ([]entity.Book, error) {
	var ents []entity.Book
	result := r.db.Preload("User").Find(&ents)
	return ents, result.Error
}

func (r *GormRepositoryBook) CheckTitleBook(title string) bool {
	var ent []entity.Book
	if err := r.db.Where("title = ?", title).Limit(1).First(&ent); err.RowsAffected == 1 {
		return false
	}
	return true
}
