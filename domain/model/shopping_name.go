package model

import "errors"

type ShoppingName string

func NewShoppingName(s string) (ShoppingName, error) {
	if s == "" {
		return "", errors.New("value is required in shopping name")
	}
	if len(s) > 30 {
		return "", errors.New("limit input to 30 chars")
	}
	return ShoppingName(s), nil
}

func (v ShoppingName) String() string {
	return string(v)
}
