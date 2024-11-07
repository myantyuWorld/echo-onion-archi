package model

import "errors"

type ShoppingCategory string

const (
	FoodCost  ShoppingCategory = "food"
	Commodity ShoppingCategory = "necessity"
)

func NewShoppingCategory(s string) (ShoppingCategory, error) {
	switch s {
	case string(FoodCost), string(Commodity):
		return ShoppingCategory(s), nil
	default:
		return "", errors.New("invalid shopping category")
	}
}

func (v ShoppingCategory) String() string {
	return string(v)
}
