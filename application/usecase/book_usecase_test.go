package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
)

func setupBookUsecase(t *testing.T) (usecase.BookUsecase, *mocks.MockBookRepository, *mocks.MockProfileRepository) {
	ctrl := gomock.NewController(t)

	bookRepo := mocks.NewMockBookRepository(ctrl)
	profileRepo := mocks.NewMockProfileRepository(ctrl)
	bookUsecase := usecase.NewBookUsecase(bookRepo, profileRepo)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return bookUsecase, bookRepo, profileRepo
}

func TestBookUsecase_CreateBook_Success(t *testing.T) {
	bookUsecase, bookRepo, profileRepo := setupBookUsecase(t)
	userID := model.GenerateNewUserID()

	bookID := model.GenerateNewBookID()
	bookRepo.EXPECT().Insert(gomock.Any()).Return(bookID, nil)
	profileRepo.EXPECT().FindByID(userID).Return(
		model.Profile{
			Role: model.Admin,
		}, nil,
	)

	input := dto.BookInput{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     model.Available.String(),
	}

	id, err := bookUsecase.CreateBook(userID.String(), input)
	assert.NoError(t, err)
	assert.Equal(t, bookID.String(), id)
}

func TestBookUsecase_CreateBook_NoPermission(t *testing.T) {
	bookUsecase, _, profileRepo := setupBookUsecase(t)

	userID := model.GenerateNewUserID()
	profileRepo.EXPECT().FindByID(userID).Return(
		model.Profile{
			Role: model.General,
		}, nil,
	)

	input := dto.BookInput{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     "invalid_status",
	}

	id, err := bookUsecase.CreateBook(userID.String(), input)

	assert.Error(t, err)
	assert.Equal(t, "", id)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrForbidden, appErr.Code)
	assert.Equal(t, "No permission", appErr.Message)
}

func TestBookUsecase_CreateBook_InvalidRequest(t *testing.T) {
	bookUsecase, _, profileRepo := setupBookUsecase(t)

	userID := model.GenerateNewUserID()
	profileRepo.EXPECT().FindByID(userID).Return(
		model.Profile{
			Role: model.Admin,
		}, nil,
	)

	input := dto.BookInput{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     "invalid_status",
	}

	id, err := bookUsecase.CreateBook(userID.String(), input)

	assert.Error(t, err)
	assert.Equal(t, "", id)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrBadReqeust, appErr.Code)
	assert.Equal(t, "invalid book status", appErr.Message)
}

func TestBookUsecase_CreateBook_InsertFailure(t *testing.T) {
	bookUsecase, bookRepo, profileRepo := setupBookUsecase(t)

	userID := model.GenerateNewUserID()
	profileRepo.EXPECT().FindByID(userID).Return(
		model.Profile{
			Role: model.Admin,
		}, nil,
	)

	bookRepo.EXPECT().Insert(gomock.Any()).Return(model.BookID(""), apperr.NewApplicationError(apperr.ErrInternalError, "Failed to insert book", errors.New("error")))

	input := dto.BookInput{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     model.Available.String(),
	}

	id, err := bookUsecase.CreateBook(userID.String(), input)

	assert.Error(t, err)
	assert.Equal(t, "", id)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrInternalError, appErr.Code)
	assert.Equal(t, "Failed to insert book", appErr.Message)
}
