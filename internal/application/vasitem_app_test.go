package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type mockVasItemRepository struct{}

var (
	listVasItemRepo    func(itemId int) ([]*entity.VasItem, error)
	createVasItemRepo  func(vasItem *entity.VasItem) (*entity.VasItem, error)
	getByIDVasItemRepo func(vasItemID int) (*entity.VasItem, error)
	deleteVasItemRepo  func(vasItemID int) error
)

func (m *mockVasItemRepository) ListByItemId(itemId int) ([]*entity.VasItem, error) {
	return listVasItemRepo(itemId)
}

func (m *mockVasItemRepository) Create(vasItem *entity.VasItem) (*entity.VasItem, error) {
	return createVasItemRepo(vasItem)
}

func (m *mockVasItemRepository) GetById(vasItemID int) (*entity.VasItem, error) {
	return getByIDVasItemRepo(vasItemID)
}

func (m *mockVasItemRepository) DeleteById(vasItemID int) error {
	return deleteVasItemRepo(vasItemID)
}

var VasItemAppMock VasItemAppInterface = &mockVasItemRepository{}

func TestSaveVasItem_Success(t *testing.T) {
	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return &entity.VasItem{
			CategoryId: 1111,
			ItemId:     2222,
			SellerId:   3333,
			Price:      340,
			Quantity:   3,
		}, nil
	}

	vasItem := &entity.VasItem{
		CategoryId: 1111,
		ItemId:     2222,
		SellerId:   3333,
		Price:      340,
		Quantity:   3,
	}

	vasItem, err := VasItemAppMock.Create(vasItem)
	assert.Nil(t, err)

	assert.Equal(t, 1111, vasItem.CategoryId)
	assert.Equal(t, 2222, vasItem.ItemId)
	assert.Equal(t, 3333, vasItem.SellerId)
	assert.Equal(t, float64(340), vasItem.Price)
	assert.Equal(t, 3, vasItem.Quantity)
}

func TestGetVasItemByID_Success(t *testing.T) {
	getByIDVasItemRepo = func(id int) (*entity.VasItem, error) {
		return &entity.VasItem{
			Id:         1,
			CategoryId: 1111,
			ItemId:     2222,
			SellerId:   3333,
			Price:      340,
			Quantity:   3,
		}, nil
	}

	vasItem, err := VasItemAppMock.GetById(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, vasItem.Id)
	assert.Equal(t, 1111, vasItem.CategoryId)
	assert.Equal(t, 2222, vasItem.ItemId)
	assert.Equal(t, 3333, vasItem.SellerId)
	assert.Equal(t, float64(340), vasItem.Price)
	assert.Equal(t, 3, vasItem.Quantity)
}

func TestListVasItem_Success(t *testing.T) {
	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return []*entity.VasItem{
			{
				Id:         1,
				CategoryId: 1111,
				ItemId:     2222,
				SellerId:   3333,
				Price:      340,
				Quantity:   3,
			},
			{
				Id:         2,
				CategoryId: 4444,
				ItemId:     5555,
				SellerId:   6666,
				Price:      340,
				Quantity:   3,
			},
		}, nil
	}

	vasItems, err := VasItemAppMock.ListByItemId(1)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(vasItems))
	assert.Equal(t, 1, vasItems[0].Id)
	assert.Equal(t, 1111, vasItems[0].CategoryId)
	assert.Equal(t, 2222, vasItems[0].ItemId)
	assert.Equal(t, 3333, vasItems[0].SellerId)
	assert.Equal(t, float64(340), vasItems[0].Price)
	assert.Equal(t, 3, vasItems[0].Quantity)
	assert.Equal(t, 2, vasItems[1].Id)
	assert.Equal(t, 4444, vasItems[1].CategoryId)
	assert.Equal(t, 5555, vasItems[1].ItemId)
	assert.Equal(t, 6666, vasItems[1].SellerId)
	assert.Equal(t, float64(340), vasItems[1].Price)
	assert.Equal(t, 3, vasItems[1].Quantity)
}

func TestDeleteVasItemByID_Success(t *testing.T) {
	deleteVasItemRepo = func(id int) error {
		return nil
	}

	err := VasItemAppMock.DeleteById(1)
	assert.Nil(t, err)
}
