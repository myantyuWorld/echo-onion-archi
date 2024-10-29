package model

type ShoppingItem struct {
	ID       uint
	OwnerID  UserID
	Category ShoppingCategory
	Name     ShoppingName
	Picked   bool
}

func NewShoppingItem(id uint, ownerID UserID, category string, name string) (*ShoppingItem, error) {
	categoryValueObject, err := NewShoppingCategory(category)
	if err != nil {
		return nil, err
	}
	nameValueObject, err := NewShoppingName(name)
	if err != nil {
		return nil, err
	}

	return &ShoppingItem{
		ID:       id,
		OwnerID:  ownerID,
		Category: categoryValueObject,
		Name:     nameValueObject,
		Picked:   false,
	}, nil
}

func ReCreate(id uint, ownerID UserID, category string, name string, picked bool) *ShoppingItem {
	categoryValueObject, _ := NewShoppingCategory(category)
	nameValueObject, _ := NewShoppingName(name)

	return &ShoppingItem{
		ID:       id,
		OwnerID:  ownerID,
		Category: categoryValueObject,
		Name:     nameValueObject,
		Picked:   picked,
	}
}
