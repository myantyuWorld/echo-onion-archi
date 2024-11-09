package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
)

type ShoppingHandler interface {
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

type shoppingHandler struct {
	usecase usecase.ShoppingItemUsecase
}

// Create implements ShoppingHandler.
func (s *shoppingHandler) Create(c echo.Context) error {
	var req dto.ShoppingItemCreateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	err := s.usecase.Create(c.Request().Context(), req.OwnerID, req.Category, req.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "shopping memo create successful"})
}

// Delete implements ShoppingHandler.
func (s *shoppingHandler) Delete(c echo.Context) error {
	var req dto.ShoppingItemDeleteRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	err := s.usecase.Delete(c.Request().Context(), req.OwnerID, req.ItemID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "shopping memo delete successful"})
}

// FindAll implements ShoppingHandler.
func (s *shoppingHandler) FindAll(c echo.Context) error {
	var req dto.ShoppingItemFindAllRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return apperr.NewApplicationError(apperr.ErrBadReqeust, "invalid input", err)
	}

	items, err := s.usecase.FindAll(c.Request().Context(), req.OwnerID)
	if err != nil {
		return err
	}

	response := make([]dto.ShoppingItemFindAllResponse, len(items))
	for i, v := range items {
		res := dto.ShoppingItemFindAllResponse{
			ItemID:      v.ID,
			Category:    v.Category.String(),
			Description: v.Name.String(),
			Picked:      v.Picked,
		}
		response[i] = res
	}

	return c.JSON(http.StatusOK, response)
}

func NewShoppingHandler(usecase usecase.ShoppingItemUsecase) ShoppingHandler {
	return &shoppingHandler{usecase}
}
