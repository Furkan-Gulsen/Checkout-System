package application

import (
	"fmt"
	"sync"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

var _ CartAppInterface = &cartApp{}

type cartApp struct {
	cartRepo     repository.CartRepositoryI
	itemApp      ItemAppInterface
	vasItemApp   VasItemAppInterface
	promotionApp PromotionAppInterface
}

func NewCartApp(cartRepo repository.CartRepositoryI, itemApp ItemAppInterface, vasItemApp VasItemAppInterface, promotionApp PromotionAppInterface) *cartApp {
	return &cartApp{
		cartRepo:     cartRepo,
		itemApp:      itemApp,
		vasItemApp:   vasItemApp,
		promotionApp: promotionApp,
	}
}

type CartAppInterface interface {
	ApplyPromotion(cartId int, promotionId int) error
	Display(cartId int) (entity.Cart, error)
	ResetCart(cartId int) error
}

func (app *cartApp) ApplyPromotion(cartId int, promotionId int) error {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		return fmt.Errorf("failed to apply promotion: %v", err)
	}

	if cart.AppliedPromotionId != 0 {
		return fmt.Errorf("promotion already applied")
	}

	promotion, err := app.promotionApp.GetById(promotionId)
	if err != nil {
		return fmt.Errorf("failed to apply promotion: %v", err)
	}

	cart.AppliedPromotionId = promotion.Id

	_, err = app.cartRepo.Update(cart)
	if err != nil {
		return fmt.Errorf("failed to apply promotion: %v", err)
	}

	return nil
}

// TODO: Aggregate veya DTO ile değiştir...
func (app *cartApp) Display(cartId int) (entity.Cart, error) {
	return app.cartRepo.GetByID(cartId)
}

func (app *cartApp) ResetCart(cartId int) error {
	err := app.cartRepo.Delete(cartId)
	if err != nil {
		return fmt.Errorf("failed to reset cart: %v", err)
	}

	items, err := app.itemApp.ListByCartId(cartId)
	if err != nil {
		return fmt.Errorf("failed to reset cart: %v", err)
	}

	var wg sync.WaitGroup

	for _, item := range items {
		wg.Add(1)
		go func(item *entity.Item) {
			defer wg.Done()

			app.itemApp.Delete(item.Id)
			vasItems, _ := app.vasItemApp.ListByItemId(item.Id)
			for _, vasItem := range vasItems {
				app.vasItemApp.DeleteById(vasItem.Id)
			}
		}(item)
	}

	wg.Wait()

	return nil
}
