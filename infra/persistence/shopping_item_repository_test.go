package persistence_test

import (
	"fmt"
	"testing"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupShoppingItemRepositoryTest() repository.ShoppingItemRepository {
	cleanUpTables(testDB, "shopping_memos")
	return persistence.NewShoppingItemRepository(testDB)
}

func TestShoppingItem_Insert_Success(t *testing.T) {
	repo := setupShoppingItemRepositoryTest()

	item, _ := domain.NewShoppingItem(domain.GenerateNewUserID(), "food", "納豆")

	err := repo.Insert(item)
	require.NoError(t, err)
}

func TestShoppingItem_FindAll_Success(t *testing.T) {
	repo := setupShoppingItemRepositoryTest()
	userID, _ := domain.NewUserID("1")

	for i := 0; i < 10; i++ {
		item, _ := domain.NewShoppingItem(userID, "food", fmt.Sprintf("item%d", i))
		repo.Insert(item)
	}

	items, err := repo.FindAll(userID)

	require.NoError(t, err)
	assert.Equal(t, len(items), 10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, items[i].Name, fmt.Sprintf("item%d", i))
		assert.Equal(t, items[i].Category, "food")
	}
}

func TestShoppingItem_Delete_Success(t *testing.T) {
	repo := setupShoppingItemRepositoryTest()
	userID, _ := domain.NewUserID("1")
	item, _ := domain.NewShoppingItem(userID, "food", "納豆")

	repo.Insert(item)
	items, _ := repo.FindAll(userID)

	err := repo.Delete(userID, items[0].ID)
	require.NoError(t, err)

	newItems, _ := repo.FindAll(userID)
	assert.Equal(t, len(newItems), 0)
}
