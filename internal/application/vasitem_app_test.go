package application

import (
	"fmt"
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
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

var VasItemAppMock repository.VasItemRepositoryI = &mockVasItemRepository{}

func TestSaveVasItem_Success(t *testing.T) {

	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return &entity.VasItem{
			Id:         1234,
			CategoryId: 3242,
			ItemId:     2222,
			SellerId:   5003.,
			Price:      340,
			Quantity:   1,
		}, nil
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return &entity.Item{
			Id:         2222,
			CategoryID: 3004,
			SellerID:   1111,
			Price:      340,
			Quantity:   3,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return []*entity.VasItem{
			{
				Id:         1,
				CategoryId: 3242,
				ItemId:     2222,
				SellerId:   5003,
				Price:      340,
				Quantity:   1,
			},
		}, nil
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)
	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5003,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, err)
	assert.Equal(t, 1234, vasItem.Id)
	fmt.Println("vasItem: ", vasItem)
	assert.Equal(t, 3242, vasItem.CategoryId)
	assert.Equal(t, 2222, vasItem.ItemId)
	assert.Equal(t, 5003, vasItem.SellerId)
	assert.Equal(t, float64(340), vasItem.Price)
	assert.Equal(t, 1, vasItem.Quantity)
}

func TestSaveVasItem_Fail(t *testing.T) {

	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return nil, fmt.Errorf("error while creating vasItem")
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return nil, fmt.Errorf("error while getting item")
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return nil, fmt.Errorf("error while listing vasItems")
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)

	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5003,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
}

func TestSaveVasItem_Fail_SellerID(t *testing.T) {

	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return nil, fmt.Errorf("error while creating vasItem")
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return &entity.Item{
			Id:         2222,
			CategoryID: 3004,
			SellerID:   1111,
			Price:      340,
			Quantity:   3,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return nil, fmt.Errorf("error while listing vasItems")
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)

	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5004,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
	assert.Equal(t, "vasItem seller id must be 5003. Seller ID: 5004", err.Error())
}

func TestSaveVasItem_Fail_ItemID(t *testing.T) {

	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return nil, fmt.Errorf("error while creating vasItem")
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return nil, fmt.Errorf("error while getting item")
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return nil, fmt.Errorf("error while listing vasItems")
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)

	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5003,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
	assert.Equal(t, "item not found. Item ID: 2222", err.Error())
}

func TestSaveVasItem_Fail_CategoryID(t *testing.T) {
	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return nil, fmt.Errorf("error while creating vasItem")
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return &entity.Item{
			Id:         2222,
			CategoryID: 1001,
			SellerID:   1111,
			Price:      340,
			Quantity:   3,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return nil, fmt.Errorf("error while listing vasItems")
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)

	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5003,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
	assert.Equal(t, "error while listing vasItems. Item ID: 2222", err.Error())
}

func TestSaveVasItem_Fail_Quantity(t *testing.T) {
	createVasItemRepo = func(vasItem *entity.VasItem) (*entity.VasItem, error) {
		return nil, fmt.Errorf("error while creating vasItem")
	}

	getByIDItemRepo = func(id int) (*entity.Item, error) {
		return &entity.Item{
			Id:         2222,
			CategoryID: 3004,
			SellerID:   1111,
			Price:      340,
			Quantity:   3,
			CartID:     1111,
			ItemType:   entity.DefaultItem,
		}, nil
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return []*entity.VasItem{
			{
				Id:         1,
				CategoryId: 3242,
				ItemId:     2222,
				SellerId:   5003,
				Price:      340,
				Quantity:   1,
			},
			{
				Id:         2,
				CategoryId: 3242,
				ItemId:     2222,
				SellerId:   5003,
				Price:      340,
				Quantity:   1,
			},
			{
				Id:         3,
				CategoryId: 3242,
				ItemId:     2222,
				SellerId:   5003,
				Price:      340,
				Quantity:   1,
			},
		}, nil
	}

	app := NewVasItemApp(VasItemAppMock, ItemAppMock)

	vasItem := &entity.VasItem{
		CategoryId: 5003,
		ItemId:     2222,
		SellerId:   5003,
		Price:      340,
		Quantity:   2,
	}

	vasItem, err := app.Create(vasItem)
	assert.Nil(t, vasItem)
	assert.NotNil(t, err)
	assert.Equal(t, "vasItem quantity cannot be more than 3. Item ID: 2222", err.Error())
}
