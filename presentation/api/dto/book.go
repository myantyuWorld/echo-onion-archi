package dto

type CreateBookRequest struct {
	Title      string `json:"title" validate:"required"`
	Author     string `json:"author"`
	CategoryID uint   `json:"category_id" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type BookIDResponse struct {
	ID string `json:"id"`
}
