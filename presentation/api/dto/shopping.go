package dto

type ShoppingItemCreateRequest struct {
	OwnerID     string `json:"owner_id" param:"owner_id" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Description string `json:"name" validate:"required"`
}

type ShoppingItemDeleteRequest struct {
	OwnerID string `json:"owner_id" param:"owner_id" validate:"required"`
	ItemID  int    `json:"item_id" param:"item_id" validate:"required"`
}

type ShoppingItemFindAllRequest struct {
	OwnerID string `json:"owner_id" param:"owner_id" validate:"required"`
}

type ShoppingItemFindAllResponse struct {
	ItemID      int    `json:"item_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Picked      bool   `json:"picked"`
}
