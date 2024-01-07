package application

import (
	"fmt"
	"log/slog"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type itemApp struct {
	itemRepo repository.ItemRepositoryI
	// categoryApp CategoryAppInterface
}

type ItemAppInterface interface {
	ListByCartId(int) ([]*entity.Item, error)
	Create(*entity.Item) (*entity.Item, error)
	Update(*entity.Item) (*entity.Item, error)
	GetById(int) (*entity.Item, error)
	Delete(int) error
}

func NewItemApp(itemRepo repository.ItemRepositoryI) *itemApp {
	return &itemApp{
		itemRepo: itemRepo,
		// categoryApp: categoryApp,
	}
}

func (app *itemApp) ListByCartId(cartId int) ([]*entity.Item, error) {
	return app.itemRepo.ListByCartId(cartId)
}

func (app *itemApp) Create(item *entity.Item) (*entity.Item, error) {
	// category, err := app.categoryApp.GetByID(item.CategoryID)
	// if err != nil || category == nil {
	// 	slog.Error("Item category not found. Error: ", err)
	// 	return nil, fmt.Errorf("item category not found. CategoryID: %d", item.CategoryID)
	// }
	// item.ItemType = category.ItemType

	if item.ItemType == entity.DigitalItem && item.Quantity > 5 {
		slog.Error("Digital item quantity can not be more than 5.")
		return nil, fmt.Errorf("digital item quantity can not be more than 5. Quantity: %d", item.Quantity)
	} else if item.ItemType == entity.DefaultItem && item.Quantity > 10 {
		slog.Error("Default item quantity can not be more than 10.")
		return nil, fmt.Errorf("default item quantity can not be more than 10. Quantity: %d", item.Quantity)
	}

	return app.itemRepo.Create(item)
}

func (app *itemApp) GetById(itemID int) (*entity.Item, error) {
	return app.itemRepo.GetById(itemID)
}

func (app *itemApp) Delete(itemID int) error {
	return app.itemRepo.Delete(itemID)
}

func (app *itemApp) Update(item *entity.Item) (*entity.Item, error) {
	return app.itemRepo.Update(item)
}
