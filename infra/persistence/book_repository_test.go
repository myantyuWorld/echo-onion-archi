package persistence_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupBookRepository() repository.BookRepository {
	cleanUpTables(testDB, "books")
	cleanUpTables(testDB, "categories")
	return persistence.NewBookRepository(testDB)
}

func TestBookReposiotry_Insert_Success(t *testing.T) {
	bookRepo := setupBookRepository()

	category := model.Category{
		ID:   1,
		Name: "Programming",
	}

	err := testDB.Create(&category).Error
	require.NoError(t, err)

	bookID := domain.GenerateNewBookID()
	book := domain.Book{
		ID:         bookID,
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     domain.Available,
	}

	result, err := bookRepo.Insert(book)

	require.NoError(t, err)
	assert.Equal(t, book.ID, result)
}

func TestBookRepository_Insert_Failure(t *testing.T) {
	bookRepo := setupBookRepository()

	bookID := domain.GenerateNewBookID()
	book := domain.Book{
		ID:         bookID,
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     domain.Available,
	}

	result, err := bookRepo.Insert(book)

	require.Error(t, err)
	assert.Equal(t, domain.BookID(""), result)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrInternalError, appErr.Code)
	assert.Equal(t, "Failed to insert book", appErr.Message)
}
