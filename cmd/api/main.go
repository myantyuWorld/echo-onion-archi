package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/config"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/db"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
)

func main() {
	e := echo.New()

	cfg := config.NewConfig()
	db, err := db.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	userRepo := persistence.NewUserRepository(db)
	profileRepo := persistence.NewProfileRepository(db)
	shoppingRepo := persistence.NewShoppingItemRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, profileRepo)
	shoppingUsecase := usecase.NewShoppingItemUsecase(shoppingRepo)

	authHandler := handler.NewAuthHandler(authUsecase, cfg.App)
	shoppingHandler := handler.NewShoppingHandler(shoppingUsecase)

	deps := &router.HandlerDependencies{
		AuthHandler:     authHandler,
		ShoppingHandler: shoppingHandler,
	}

	router.NewRouter(e, deps)
	e.Start(":8080")
}
