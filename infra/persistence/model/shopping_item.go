package model

import (
	"time"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type ShoppingItem struct {
	ID          int       `gorm:"primaryKey;not null"`
	UserID      string    `gorm:"not null"`
	Category    string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Picked      bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func ToModelShoppingItem(shoppingItem *domain.ShoppingItem) ShoppingItem {
	return ShoppingItem{
		ID:          shoppingItem.ID,
		UserID:      shoppingItem.OwnerID.String(),
		Category:    shoppingItem.Category.String(),
		Description: shoppingItem.Name.String(),
		Picked:      false,
	}
}
