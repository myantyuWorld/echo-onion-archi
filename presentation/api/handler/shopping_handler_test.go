package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/config"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
	"github.com/sakaguchi-0725/echo-onion-arch/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupShoppingItemHandler(t *testing.T) (*mocks.MockShoppingItemUsecase, *echo.Echo) {
	ctrl := gomock.NewController(t)
	authUsecase := mocks.NewMockAuthUsecase(ctrl)
	authHandler := handler.NewAuthHandler(authUsecase, config.NewConfig().App)
	usecase := mocks.NewMockShoppingItemUsecase(ctrl)
	handler := handler.NewShoppingHandler(usecase)

	e := echo.New()
	deps := router.HandlerDependencies{
		AuthHandler:     authHandler,
		ShoppingHandler: handler,
	}
	router.NewRouter(e, &deps)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return usecase, e
}

func TestShoppingItemHandler_FindAll_Success(t *testing.T) {
	usecase, e := setupShoppingItemHandler(t)

	rec, req := test.SetupRequest(e, http.MethodGet, "/shopping/1", "")
	usecase.EXPECT().FindAll(context.Background(), "1").Return([]*model.ShoppingItem{
		{
			ID:       0,
			OwnerID:  "1",
			Category: "necessity",
			Name:     "test",
			Picked:   false,
		},
		{
			ID:       1,
			OwnerID:  "1",
			Category: "food",
			Name:     "test",
			Picked:   true,
		},
	}, nil)

	e.ServeHTTP(rec, req)

	expectedRes := `[
		{
			"item_id" : 0,
			"category" : "necessity",
			"description" : "test",
			"picked" : false
		},
		{
			"item_id" : 1,
			"category" : "food",
			"description" : "test",
			"picked" : true
		}
	]`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())

}

func TestShoppingItemHandler_Create_Success(t *testing.T) {
	usecase, e := setupShoppingItemHandler(t)

	request := dto.ShoppingItemCreateRequest{
		OwnerID:     "1",
		Category:    "food",
		Description: "test",
	}
	reqBody, err := json.Marshal(request)
	require.NoError(t, err)

	rec, req := test.SetupRequest(e, http.MethodPost, "/shopping/1", string(reqBody))
	usecase.EXPECT().Create(context.Background(), request.OwnerID, request.Category, request.Description).Return(nil)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestShoppingItemHandler_Delete_Success(t *testing.T) {
	usecase, e := setupShoppingItemHandler(t)

	request := dto.ShoppingItemDeleteRequest{
		OwnerID: "1",
		ItemID:  1,
	}

	rec, req := test.SetupRequest(e, http.MethodDelete, "/shopping/1/1", "")
	usecase.EXPECT().Delete(context.Background(), request.OwnerID, request.ItemID).Return(nil)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
