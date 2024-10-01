package persistence

import (
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"gorm.io/gorm"
)

type shoppingItem struct {
	db *gorm.DB
}

// Delete implements repository.ShoppingItemRepository.
func (s *shoppingItem) Delete(id uint) error {
	panic("unimplemented")
}

// FindAll implements repository.ShoppingItemRepository.
func (s *shoppingItem) FindAll() ([]*model.ShoppingItem, error) {
	panic("unimplemented")
}

// FindByID implements repository.ShoppingItemRepository.
func (s *shoppingItem) FindByID(id uint) (*model.ShoppingItem, error) {
	panic("unimplemented")
}

// Insert implements repository.ShoppingItemRepository.
func (s *shoppingItem) Insert(v model.ShoppingItem) error {
	panic("unimplemented")
}

// Update implements repository.ShoppingItemRepository.
func (s *shoppingItem) Update(v model.ShoppingItem) error {
	panic("unimplemented")
}

func NewShoppingItemRepository(db *gorm.DB) repository.ShoppingItemRepository {
	return &shoppingItem{db: db}
}
