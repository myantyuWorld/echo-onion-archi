//go:generate mockgen -source=$GOFILE -destination=../../mocks/application/usecase/mock_$GOFILE -package=mocks
package usecase

import (
	"context"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
)

type ShoppingItemUsecase interface {
	FindAll(ctx context.Context, userID string) ([]*domain.ShoppingItem, error)
	Delete(ctx context.Context, userID string, itemID int) error
	Create(ctx context.Context, userID string, category string, description string) error
}

type shoppingUsecase struct {
	shoppingRepo repository.ShoppingItemRepository
}

// Create implements ShoppingItemUsecase.
func (s *shoppingUsecase) Create(ctx context.Context, userID string, category string, description string) error {
	domainUserID, err := domain.NewUserID(userID)
	if err != nil {
		return err
	}

	newShoppingItem, err := domain.NewShoppingItem(domainUserID, category, description)
	if err != nil {
		return err
	}

	err = s.shoppingRepo.Insert(newShoppingItem)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements ShoppingItemUsecase.
func (s *shoppingUsecase) Delete(ctx context.Context, userID string, itemID int) error {
	domainUserID, err := domain.NewUserID(userID)
	if err != nil {
		return err
	}
	err = s.shoppingRepo.Delete(domainUserID, itemID)

	if err != nil {
		return err
	}
	return nil
}

// FindAll implements ShoppingItemUsecase.
func (s *shoppingUsecase) FindAll(ctx context.Context, userID string) ([]*domain.ShoppingItem, error) {
	domainUserID, err := domain.NewUserID(userID)
	if err != nil {
		return []*domain.ShoppingItem{}, err
	}

	items, err := s.shoppingRepo.FindAll(domainUserID)

	if err != nil {
		return []*domain.ShoppingItem{}, err
	}
	return items, nil
}

func NewShoppingItemUsecase(shoppingRepo repository.ShoppingItemRepository) ShoppingItemUsecase {
	return &shoppingUsecase{
		shoppingRepo: shoppingRepo,
	}
}
