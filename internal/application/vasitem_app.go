package application

import (
	"fmt"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type vasItemApp struct {
	vasItemRepo repository.VasItemRepositoryI
	categoryApp CategoryAppInterface
	itemApp     ItemAppInterface
}

type VasItemAppInterface interface {
	ListByItemId(int) ([]*entity.VasItem, error)
	GetById(int) (*entity.VasItem, error)
	Create(*entity.VasItem) error
	DeleteById(int) error
}

func NewVasItemApp(vir repository.VasItemRepositoryI, cai CategoryAppInterface, iti ItemAppInterface) *vasItemApp {
	return &vasItemApp{
		vasItemRepo: vir,
		categoryApp: cai,
		itemApp:     iti,
	}
}

func (app *vasItemApp) ListByItemId(itemId int) ([]*entity.VasItem, error) {
	return app.vasItemRepo.ListByItemId(itemId)
}

func (app *vasItemApp) GetById(vasItemId int) (*entity.VasItem, error) {
	return app.vasItemRepo.GetById(vasItemId)
}

func (app *vasItemApp) Create(vasItem *entity.VasItem) error {
	_, err := app.categoryApp.GetByID(vasItem.CategoryId)
	if err != nil {
		return fmt.Errorf("category not found. Category ID: %d", vasItem.CategoryId)
	}

	_, err = app.itemApp.GetById(vasItem.ItemId)
	if err != nil {
		return fmt.Errorf("item not found. Item ID: %d", vasItem.ItemId)
	}

	return app.vasItemRepo.Create(vasItem)
}

func (app *vasItemApp) DeleteById(vasItemId int) error {
	return app.vasItemRepo.DeleteById(vasItemId)
}
