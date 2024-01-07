package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type mockItemRepository struct{}

var (
	listItemRepo    func(id int) ([]*entity.Item, error)
	createItemRepo  func(item *entity.Item) (*entity.Item, error)
	updateItemRepo  func(item *entity.Item) (*entity.Item, error)
	getByIDItemRepo func(id int) (*entity.Item, error)
	deleteItemRepo  func(id int) error
)

func (m *mockItemRepository) ListByCartId(cartId int) ([]*entity.Item, error) {
	return listItemRepo(cartId)
}

func (m *mockItemRepository) Create(item *entity.Item) (*entity.Item, error) {
	return createItemRepo(item)
}

func (m *mockItemRepository) Update(item *entity.Item) (*entity.Item, error) {
	return updateItemRepo(item)
}

func (m *mockItemRepository) GetById(id int) (*entity.Item, error) {
	return getByIDItemRepo(id)
}

func (m *mockItemRepository) Delete(id int) error {
	return deleteItemRepo(id)
}

var ItemAppMock ItemAppInterface = &mockItemRepository{}

func TestSaveItem_Success(t *testing.T) {
	createItemRepo = func(item *entity.Item) (*entity.Item, error) {
		return &entity.Item{
			Id:         1111,
			CategoryID: 1111,
			SellerID:   1111,
			Price:      100,
			Quantity:   5,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	item := &entity.Item{
		CategoryID: 1111,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     1111,
		ItemType:   entity.DefaultItem,
	}

	item, err := ItemAppMock.Create(item)
	assert.Nil(t, err)
	assert.Equal(t, 1111, item.Id)
	assert.Equal(t, 1111, item.CategoryID)
	assert.Equal(t, 1111, item.SellerID)
	assert.Equal(t, float64(100), item.Price)
	assert.Equal(t, 5, item.Quantity)
	assert.Equal(t, 1111, item.CartID)
	assert.Equal(t, entity.DefaultItem, item.ItemType)

}

func TestGetItemByID_Success(t *testing.T) {
	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return &entity.Item{
			Id:         1111,
			CategoryID: 1111,
			SellerID:   1111,
			Price:      100,
			Quantity:   5,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	item, err := ItemAppMock.GetById(1111)
	assert.Nil(t, err)
	assert.Equal(t, 1111, item.Id)
	assert.Equal(t, 1111, item.CategoryID)
	assert.Equal(t, 1111, item.SellerID)
	assert.Equal(t, float64(100), item.Price)
	assert.Equal(t, 5, item.Quantity)
	assert.Equal(t, 1111, item.CartID)
	assert.Equal(t, entity.DefaultItem, item.ItemType)
}

func TestGetItemByID_Fail(t *testing.T) {
	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return nil, nil
	}

	item, err := ItemAppMock.GetById(1111)
	assert.Nil(t, item)
	assert.Nil(t, err)
}

func TestListItemsByCartID_Success(t *testing.T) {
	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			{
				Id:         1111,
				CategoryID: 1111,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     1111,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         2222,
				CategoryID: 2222,
				SellerID:   2222,
				Price:      200,
				Quantity:   10,
				CartID:     2222,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	items, err := ItemAppMock.ListByCartId(1111)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, 1111, items[0].Id)
	assert.Equal(t, 1111, items[0].CategoryID)
	assert.Equal(t, 1111, items[0].SellerID)
	assert.Equal(t, float64(100), items[0].Price)
	assert.Equal(t, 5, items[0].Quantity)
	assert.Equal(t, 1111, items[0].CartID)
	assert.Equal(t, entity.DefaultItem, items[0].ItemType)
	assert.Equal(t, 2222, items[1].Id)
	assert.Equal(t, 2222, items[1].CategoryID)
	assert.Equal(t, 2222, items[1].SellerID)
	assert.Equal(t, float64(200), items[1].Price)
	assert.Equal(t, 10, items[1].Quantity)
	assert.Equal(t, 2222, items[1].CartID)
	assert.Equal(t, entity.DefaultItem, items[1].ItemType)
}

func TestListItemsByCartID_Fail(t *testing.T) {
	listItemRepo = func(id int) ([]*entity.Item, error) {
		return nil, nil
	}

	items, err := ItemAppMock.ListByCartId(1111)
	assert.Nil(t, items)
	assert.Nil(t, err)
}
