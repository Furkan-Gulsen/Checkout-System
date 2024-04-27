package persistence

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/utils"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func setUpItemRepo(t *testing.T) (*ItemRepository, func()) {
	db := utils.SetUpMockDatabase(t)
	repo := NewItemRepository(db, constants.MOCK_DB_NAME)

	return repo, func() {
		utils.CleanUpMockDatabase(db, "items")
	}
}

func TestCreateItem_Success(t *testing.T) {
	repo, tearDown := setUpItemRepo(t)
	defer tearDown()

	item := &entity.Item{
		Id:         1,
		CategoryID: 1,
		SellerID:   1,
		Price:      100,
		Quantity:   5,
		CartID:     1,
		ItemType:   entity.DefaultItem,
	}

	item, err := repo.Create(item)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestCreateItem_Failure(t *testing.T) {
	item := &entity.Item{
		Id:         1,
		CategoryID: 1,
		SellerID:   1,
		Price:      100,
		Quantity:   5,
		CartID:     1,
		ItemType:   4,
	}

	err := item.Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Validation errors: ItemType must be one of 1 2.")
}

func TestListItemsByCartID(t *testing.T) {
	repo, tearDown := setUpItemRepo(t)
	defer tearDown()

	_, _ = repo.Create(&entity.Item{Id: 1, CategoryID: 1, SellerID: 1, Price: 100, Quantity: 5, CartID: 1, ItemType: entity.DefaultItem})
	_, _ = repo.Create(&entity.Item{Id: 2, CategoryID: 2, SellerID: 2, Price: 200, Quantity: 5, CartID: 1, ItemType: entity.DefaultItem})

	items, err := repo.ListByCartId(1)
	assert.Nil(t, err)
	assert.Len(t, items, 2)
}

func TestListItemsByCartID_Failure(t *testing.T) {
	repo, tearDown := setUpItemRepo(t)
	defer tearDown()

	items, err := repo.ListByCartId(1)
	assert.Nil(t, items)
	assert.Nil(t, err)
}

func TestGetItemByID_Success(t *testing.T) {

	repo, tearDown := setUpItemRepo(t)
	defer tearDown()

	expectedItem := &entity.Item{Id: 1, CategoryID: 1, SellerID: 1, Price: 100, Quantity: 5, CartID: 1, ItemType: entity.DefaultItem}
	item, err := repo.Create(expectedItem)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	item, err = repo.GetById(item.Id)
	assert.Nil(t, err)
	assert.Equal(t, expectedItem, item)
}

func TestGetItemByID_NotFound(t *testing.T) {
	repo, tearDown := setUpItemRepo(t)
	defer tearDown()

	item, err := repo.GetById(1)
	assert.Nil(t, item)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "could not find item with id 1: mongo: no documents in result")
}
