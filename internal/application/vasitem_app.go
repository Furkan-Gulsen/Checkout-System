package application

import (
	"fmt"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type vasItemApp struct {
	vasItemRepo repository.VasItemRepositoryI
	itemApp     ItemAppInterface
}

type VasItemAppInterface interface {
	ListByItemId(int) ([]*entity.VasItem, error)
	GetById(int) (*entity.VasItem, error)
	Create(*entity.VasItem) (*entity.VasItem, error)
	DeleteById(int) error
}

func NewVasItemApp(vir repository.VasItemRepositoryI, iti ItemAppInterface) *vasItemApp {
	return &vasItemApp{
		vasItemRepo: vir,
		itemApp:     iti,
	}
}

func (app *vasItemApp) ListByItemId(itemId int) ([]*entity.VasItem, error) {
	return app.vasItemRepo.ListByItemId(itemId)
}

func (app *vasItemApp) GetById(vasItemId int) (*entity.VasItem, error) {
	return app.vasItemRepo.GetById(vasItemId)
}

func (app *vasItemApp) Create(vasItem *entity.VasItem) (*entity.VasItem, error) {
	// * Check Seller ID
	if vasItem.SellerId != 5003 {
		return nil, fmt.Errorf("vasItem seller id must be 5003. Seller ID: %d", vasItem.SellerId)
	}

	// * Check if item exists
	item, getItemErr := app.itemApp.GetById(vasItem.ItemId)
	if getItemErr != nil {
		return nil, fmt.Errorf("item not found. Item ID: %d", vasItem.ItemId)
	}
	fmt.Println("item: ", item.CategoryID)

	// * Check if item category id is 1001 or 3004
	if item.CategoryID != 1001 && item.CategoryID != 3004 {
		return nil, fmt.Errorf("vasItem cannot be added to this product. Category ID: %d", vasItem.CategoryId)
	}

	// * Check if vasItem quantity is more than 3
	vasItems, listVasItemErr := app.ListByItemId(vasItem.ItemId)
	if listVasItemErr != nil {
		return nil, fmt.Errorf("error while listing vasItems. Item ID: %d", vasItem.ItemId)
	}
	vasItemsQuantity := vasItem.Quantity
	for _, vasItem := range vasItems {
		vasItemsQuantity += vasItem.Quantity
	}
	if vasItemsQuantity > 3 {
		return nil, fmt.Errorf("vasItem quantity cannot be more than 3. Item ID: %d", vasItem.ItemId)
	}

	return app.vasItemRepo.Create(vasItem)
}

func (app *vasItemApp) DeleteById(vasItemId int) error {
	return app.vasItemRepo.DeleteById(vasItemId)
}
