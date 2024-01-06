package application

import (
	"fmt"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type vasItemApp struct {
	vasItemRepo  repository.VasItemRepositoryI
	categoryRepo repository.CategoryRepositoryI
	itemRepo     repository.ItemRepositoryI
}

type VasItemAppInterface interface {
	ListByItemId(int) ([]*entity.VasItem, error)
	GetById(int) (*entity.VasItem, error)
	Create(*entity.VasItem) error
}

func NewVasItemApp(vir repository.VasItemRepositoryI, cr repository.CategoryRepositoryI, ir repository.ItemRepositoryI) *vasItemApp {
	return &vasItemApp{
		vasItemRepo:  vir,
		categoryRepo: cr,
		itemRepo:     ir,
	}
}

func (app *vasItemApp) ListByItemId(itemId int) ([]*entity.VasItem, error) {
	return app.vasItemRepo.ListByItemId(itemId)
}

func (app *vasItemApp) GetById(vasItemId int) (*entity.VasItem, error) {
	return app.vasItemRepo.GetById(vasItemId)
}

func (app *vasItemApp) Create(vasItem *entity.VasItem) error {
	_, err := app.categoryRepo.GetByID(vasItem.CategoryId)
	if err != nil {
		return fmt.Errorf("category not found. Category ID: %d", vasItem.CategoryId)
	}

	_, err = app.itemRepo.GetById(vasItem.ItemId)
	if err != nil {
		return fmt.Errorf("item not found. Item ID: %d", vasItem.ItemId)
	}

	return app.vasItemRepo.Create(vasItem)
}
