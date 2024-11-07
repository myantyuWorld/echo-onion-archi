//go:generate mockgen -source=$GOFILE -destination=../../mocks/domain/repository/mock_$GOFILE -package=mocks
package repository

import (
	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type ShoppingItemRepository interface {
	FindAll(userID domain.UserID) ([]*domain.ShoppingItem, error)
	FindByID(id string) (*domain.ShoppingItem, error)
	Insert(v *domain.ShoppingItem) error
	Update(v *domain.ShoppingItem) error
	Delete(userID domain.UserID, itemID int) error
}
