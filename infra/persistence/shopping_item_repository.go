package persistence

import (
	"errors"
	"fmt"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence/model"
	"gorm.io/gorm"
)

type shoppingItem struct {
	db *gorm.DB
}

// Delete implements repository.ShoppingItemRepository.
func (s *shoppingItem) Delete(userID domain.UserID, itemID int) error {
	result := s.db.Where("user_id = ?", userID.String()).Delete(&model.ShoppingItem{}, itemID)

	if result.Error != nil {
		return apperr.NewApplicationError(apperr.ErrInternalError, "Failed to delete shopping memo", result.Error)
	}

	// TODO : 削除扱いにしてもいいかも（フロントエンドから見た時は削除されたことと同義か）
	if result.RowsAffected == 0 {
		return apperr.NewApplicationError(apperr.ErrNotFound, fmt.Sprintf("Shopping Memo with ID %d not found", itemID), nil)
	}

	return nil
}

// FindAll implements repository.ShoppingItemRepository.
func (s *shoppingItem) FindAll(userID domain.UserID) ([]*domain.ShoppingItem, error) {
	var results []*model.ShoppingItem

	err := s.db.Where("user_id = ?", userID).Find(&results).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.ShoppingItem{}, apperr.NewApplicationError(apperr.ErrNotFound, fmt.Sprintf("ShoppingItems with userID %s not found", userID), err)
		}
		return []*domain.ShoppingItem{}, apperr.NewApplicationError(apperr.ErrInternalError, "Failed to FindAll Shopping Items", err)
	}

	var shoppingItems []*domain.ShoppingItem
	for _, v := range results {
		shoppingItems = append(shoppingItems, domain.ReCreateShoppingItem(
			v.ID,
			domain.UserID(v.UserID),
			v.Category,
			v.Description,
			v.Picked,
		))
	}

	return shoppingItems, nil
}

// FindByID implements repository.ShoppingItemRepository.
func (s *shoppingItem) FindByID(id string) (*domain.ShoppingItem, error) {
	panic("unimplemented")
}

// Insert implements repository.ShoppingItemRepository.
func (s *shoppingItem) Insert(v *domain.ShoppingItem) error {
	item := model.ToModelShoppingItem(v)

	if err := s.db.Create(&item).Error; err != nil {
		return apperr.NewApplicationError(apperr.ErrInternalError, "Failed to insert shopping item", err)
	}

	v.ID = item.ID

	return nil
}

// Update implements repository.ShoppingItemRepository.
func (s *shoppingItem) Update(v *domain.ShoppingItem) error {
	panic("unimplemented")
}

func NewShoppingItemRepository(db *gorm.DB) repository.ShoppingItemRepository {
	return &shoppingItem{db: db}
}
