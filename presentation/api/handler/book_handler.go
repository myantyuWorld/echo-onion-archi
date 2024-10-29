package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	input "github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
)

type BookHandler interface {
	CreateBook(c echo.Context) error
}

type bookHandler struct {
	usecase usecase.BookUsecase
}

func NewBookHandler(usecase usecase.BookUsecase) BookHandler {
	return &bookHandler{usecase}
}

func (b *bookHandler) CreateBook(c echo.Context) error {
	userID := c.Get("userID").(string)

	var req dto.CreateBookRequest
	if err := c.Bind(&req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid request", err)
	}

	if err := c.Validate(req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid request", err)
	}

	input := input.BookInput{
		Title:      req.Title,
		Author:     req.Author,
		CategoryID: req.CategoryID,
		Status:     req.Status,
	}

	bookID, err := b.usecase.CreateBook(userID, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto.BookIDResponse{ID: bookID})
}
