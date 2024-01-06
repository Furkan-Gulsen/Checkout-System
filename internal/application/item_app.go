package application

import (
	"fmt"
	"log/slog"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type itemApp struct {
	itemRepo     repository.ItemRepositoryI
	categoryRepo repository.CategoryRepositoryI
}

type ItemAppInterface interface {
	List() ([]*entity.Item, error)
	Create(*entity.Item) error
	GetById(int) (*entity.Item, error)
	Delete(int) error
}

func NewItemApp(itemRepo repository.ItemRepositoryI, categoryRepo repository.CategoryRepositoryI) *itemApp {
	return &itemApp{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
	}
}

func (app *itemApp) List() ([]*entity.Item, error) {
	return app.itemRepo.List()
}

func (app *itemApp) Create(item *entity.Item) error {
	category, err := app.categoryRepo.GetByID(item.CategoryID)
	if err != nil {
		slog.Error("Item category not found. Error: ", err)
		return fmt.Errorf("item category not found. CategoryID: %d", item.CategoryID)
	}

	item.ItemType = category.ItemType
	fmt.Printf("itemType: %+v\n", item.ItemType)

	return app.itemRepo.Create(item)
}

func (app *itemApp) GetById(itemID int) (*entity.Item, error) {
	return app.itemRepo.GetById(itemID)
}

func (app *itemApp) Delete(itemID int) error {
	return app.itemRepo.Delete(itemID)
}
