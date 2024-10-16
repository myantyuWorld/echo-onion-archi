//go:generate mockgen -source=book_usecase.go -destination=../../mocks/application/usecase/mock_book_usecase.go -package=mocks
package usecase

import (
	"github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/auth"
)

type BookUsecase interface {
	CreateBook(idStr string, input dto.BookInput) (string, error)
}

type bookUsecase struct {
	bookRepo    repository.BookRepository
	profileRepo repository.ProfileRepository
}

func NewBookUsecase(bookRepo repository.BookRepository, profileRepo repository.ProfileRepository) BookUsecase {
	return &bookUsecase{bookRepo, profileRepo}
}

func (b *bookUsecase) CreateBook(idStr string, input dto.BookInput) (string, error) {
	bookID := model.GenerateNewBookID()
	userID, err := model.NewUserID(idStr)
	if err != nil {
		return "", apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid userID", err)
	}

	profile, err := b.profileRepo.FindByID(userID)
	if !auth.IsAdmin(profile) {
		return "", apperr.NewApplicationError(apperr.ErrForbidden, "No permission", err)
	}

	status, err := model.NewBookStatus(input.Status)
	if err != nil {
		return "", apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid book status", err)
	}

	book, err := model.NewBook(bookID, input.Title, input.Author, input.CategoryID, status)
	if err != nil {
		return "", apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid request", err)
	}

	createdID, err := b.bookRepo.Insert(book)
	if err != nil {
		return "", err
	}

	return createdID.String(), nil
}
