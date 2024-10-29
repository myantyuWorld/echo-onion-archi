package persistence

import (
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence/model"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) FindAll(filter repository.BookFilter) ([]domain.Book, error) {
	panic("unimplemented")
}

func (b *bookRepository) FindByFilter(title *string, author *string) {
	panic("unimplemented")
}

func (b *bookRepository) Insert(book domain.Book) (domain.BookID, error) {
	input := model.ToModelBook(book)

	if err := b.db.Create(&input).Error; err != nil {
		return domain.BookID(""), apperr.NewApplicationError(apperr.ErrInternalError, "Failed to insert book", err)
	}

	return domain.BookID(input.ID), nil
}

func (b *bookRepository) Update(book domain.Book) (domain.BookID, error) {
	panic("unimplemented")
}
