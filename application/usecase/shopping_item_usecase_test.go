package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
)

func setupShoppingUsecase(t *testing.T) (usecase.ShoppingItemUsecase, *mocks.MockShoppingItemRepository) {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockShoppingItemRepository(ctrl)
	authUsecase := usecase.NewShoppingItemUsecase(repo)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return authUsecase, repo
}

func TestShoppingUsecase_FindAll_Success(t *testing.T) {
	usecase, repo := setupShoppingUsecase(t)

	domainUserID := domain.GenerateNewUserID()

	repo.EXPECT().FindAll(domainUserID).Return([]*domain.ShoppingItem{
		{
			ID:       0,
			OwnerID:  "",
			Category: "",
			Name:     "",
			Picked:   false,
		},
		{
			ID:       0,
			OwnerID:  "",
			Category: "",
			Name:     "",
			Picked:   false,
		},
	}, nil)

	items, err := usecase.FindAll(context.TODO(), domainUserID.String())
	assert.NoError(t, err)
	assert.Equal(t, len(items), 2)

}
func TestShoppingUsecase_Insert_SUCCESS(t *testing.T) {
	usecase, repo := setupShoppingUsecase(t)

	domainUserID := domain.GenerateNewUserID()

	repo.EXPECT().Insert(&domain.ShoppingItem{
		ID:       0,
		OwnerID:  domain.UserID(domainUserID.String()),
		Category: "food",
		Name:     "test",
		Picked:   false,
	}).Return(nil)

	err := usecase.Create(context.TODO(), domainUserID.String(), "food", "test")

	assert.NoError(t, err)
}

func TestShoppingUsecase_Delete(t *testing.T) {
	usecase, repo := setupShoppingUsecase(t)

	domainUserID := domain.GenerateNewUserID()

	repo.EXPECT().Delete(domainUserID, 1).Return(nil)

	err := usecase.Delete(context.TODO(), domainUserID.String(), 1)

	assert.NoError(t, err)
}
