package persistence

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/constants"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func setUpVasItemRepo(t *testing.T) (*VasItemRepository, func()) {
	db := utils.SetUpMockDatabase(t)
	repo := NewVasItemRepository(db, constants.MOCK_DB_NAME)

	return repo, func() {
		utils.CleanUpMockDatabase(db, "vasitems")
	}
}

func TestCreateVasItem_Success(t *testing.T) {
	repo, tearDown := setUpVasItemRepo(t)
	defer tearDown()

	vasItem := &entity.VasItem{
		CategoryId: 1111,
		ItemId:     2222,
		SellerId:   3333,
		Price:      340,
		Quantity:   3,
	}

	vasItem, err := repo.Create(vasItem)
	assert.Nil(t, err)

	assert.Equal(t, vasItem.CategoryId, 1111)
	assert.Equal(t, vasItem.ItemId, 2222)
	assert.Equal(t, vasItem.SellerId, 3333)
	assert.Equal(t, vasItem.Price, float64(340))
	assert.Equal(t, vasItem.Quantity, 3)
}

func TestGetVasItemById_Success(t *testing.T) {
	repo, tearDown := setUpVasItemRepo(t)
	defer tearDown()

	vasItem := &entity.VasItem{
		CategoryId: 1111,
		ItemId:     2222,
		SellerId:   3333,
		Price:      340,
		Quantity:   3,
	}

	vasItem, err := repo.Create(vasItem)
	if err != nil {
		t.Error("Error while saving vas item")
	}

	vasItem, err = repo.GetById(vasItem.Id)
	assert.Nil(t, err)

	assert.Equal(t, vasItem.CategoryId, 1111)
	assert.Equal(t, vasItem.ItemId, 2222)
	assert.Equal(t, vasItem.SellerId, 3333)
	assert.Equal(t, vasItem.Price, float64(340))
	assert.Equal(t, vasItem.Quantity, 3)
}

func TestListVasItemByItemId_Success(t *testing.T) {
	repo, tearDown := setUpVasItemRepo(t)
	defer tearDown()

	vasItem := &entity.VasItem{
		CategoryId: 1111,
		ItemId:     2222,
		SellerId:   3333,
		Price:      340,
		Quantity:   3,
	}

	vasItem, err := repo.Create(vasItem)
	if err != nil {
		t.Error("Error while saving vas item")
	}

	vasItems, err := repo.ListByItemId(vasItem.ItemId)
	assert.Nil(t, err)

	assert.Equal(t, vasItems[0].CategoryId, 1111)
	assert.Equal(t, vasItems[0].ItemId, 2222)
	assert.Equal(t, vasItems[0].SellerId, 3333)
	assert.Equal(t, vasItems[0].Price, float64(340))
	assert.Equal(t, vasItems[0].Quantity, 3)
}

func TestDeleteVasItemById_Success(t *testing.T) {
	repo, tearDown := setUpVasItemRepo(t)
	defer tearDown()

	vasItem := &entity.VasItem{
		CategoryId: 1111,
		ItemId:     2222,
		SellerId:   3333,
		Price:      340,
		Quantity:   3,
	}

	vasItem, err := repo.Create(vasItem)
	if err != nil {
		t.Error("Error while saving vas item")
	}

	err = repo.DeleteById(vasItem.Id)
	assert.Nil(t, err)

	vasItem, err = repo.GetById(vasItem.Id)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
}
