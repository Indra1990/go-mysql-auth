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

func (r *GormRepositoryBook) CreateMultiple(ents []entity.Book) error {
	// var ents []entity.Book
	err := r.db.Create(&ents)
	if err != nil {
		return err.Error
	}
	return nil
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

func (r *GormRepositoryBook) FindByIDBook(idbook int) (entity.Book, error) {
	var ent entity.Book
	findBook := r.db.Preload("User").First(&ent, idbook)
	return ent, findBook.Error
}

func (r *GormRepositoryBook) UpdateBook(ent entity.Book) error {
	updateBook := r.db.Model(&ent).Updates(ent)
	return updateBook.Error
}

func (r *GormRepositoryBook) DeleteBook(ent entity.Book) error {
	delete := r.db.Delete(&ent)
	return delete.Error
}

func (r *GormRepositoryBook) DeleteBookMultiple(ents []entity.Book, idbooks []uint) error {
	delete := r.db.Delete(&ents, idbooks)
	return delete.Error
}
