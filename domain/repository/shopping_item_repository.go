//go:generate mockgen -source=loan_repository.go -destination=../../mocks/domain/repository/mock_loan_repository.go -package=mocks
package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type ShoppingItemRepository interface {
	FindAll() ([]*model.ShoppingItem, error)
	FindByID(id uint) (*model.ShoppingItem, error)
	Insert(v model.ShoppingItem) error
	Update(v model.ShoppingItem) error
	Delete(id uint) error
}
