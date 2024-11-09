package model

type ShoppingItem struct {
	ID       int
	OwnerID  UserID
	Category ShoppingCategory
	Name     ShoppingName
	Picked   bool
}

func NewShoppingItem(ownerID UserID, category string, name string) (*ShoppingItem, error) {
	categoryValueObject, err := NewShoppingCategory(category)
	if err != nil {
		return nil, err
	}
	nameValueObject, err := NewShoppingName(name)
	if err != nil {
		return nil, err
	}

	return &ShoppingItem{
		OwnerID:  ownerID,
		Category: categoryValueObject,
		Name:     nameValueObject,
		Picked:   false,
	}, nil
}

func ReCreateShoppingItem(id int, ownerID UserID, category string, name string, picked bool) *ShoppingItem {
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
