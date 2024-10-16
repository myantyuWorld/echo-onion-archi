package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	input "github.com/sakaguchi-0725/echo-onion-arch/application/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/auth"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/dto"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupBookHandler(t *testing.T) (*mocks.MockBookUsecase, *echo.Echo) {
	ctrl := gomock.NewController(t)
	bookUsecase := mocks.NewMockBookUsecase(ctrl)
	bookHandler := handler.NewBookHandler(bookUsecase)

	e := echo.New()
	deps := NewMockHandlersDependency(ctrl, SetBookHandler(bookHandler))
	router.NewRouter(e, deps, cfg)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return bookUsecase, e
}

func TestBookHandler_CreateBook_Success(t *testing.T) {
	bookUsecase, e := setupBookHandler(t)

	userID := model.GenerateNewUserID()
	token, err := auth.GenerateToken(userID, cfg.JWTSecret)
	require.NoError(t, err)

	request := dto.CreateBookRequest{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     "available",
	}

	reqBody, err := json.Marshal(request)
	require.NoError(t, err)

	rec, req := SetupRequest(e, http.MethodPost, "/books", string(reqBody))
	req.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: token,
	})

	bookID := model.GenerateNewBookID()
	bookUsecase.EXPECT().CreateBook(userID.String(), input.BookInput{
		Title:      "スターティングGo言語",
		Author:     "松尾 愛賀",
		CategoryID: 1,
		Status:     "available",
	}).Return(bookID.String(), nil)

	e.ServeHTTP(rec, req)

	expectedRes := fmt.Sprintf(`{
		"id": "%s"
	}`, bookID.String())

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}
