package application

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
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
	ApplyPromotion(cartId int, promotionId int) (*entity.Cart, error)
	Display(cartId int) (*dto.DisplayCartDTO, error)
	ResetCart(cartId int) (*entity.Cart, error)
	AddItem(cartId int, item *entity.Item) (*entity.Cart, error)
	UpdateCartPriceAndQuantity(cart *entity.Cart) (*entity.Cart, error)
	AddVasItem(vasItem *entity.VasItem) (*entity.Cart, error)
}

func (app *cartApp) ApplyPromotion(cartId int, promotionId int) (*entity.Cart, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		slog.Error("Cart not found. Error: ", err)
		return nil, fmt.Errorf("cart not found. CartID: %d", cartId)
	}

	if cart.AppliedPromotionId != 0 {
		return nil, fmt.Errorf("promotion already applied")
	}

	promotion, err := app.promotionApp.GetById(promotionId)
	if err != nil {
		slog.Error("Promotion not found. Error: ", err)
		return nil, fmt.Errorf("promotion is not found, promotionId: %d", promotionId)
	}

	items, _ := app.itemApp.ListByCartId(cartId)
	if len(items) > 0 {
		cart = calcCartPricesWithPromotion(cart, items, promotion)
		app.cartRepo.Update(cart)
	}

	cart.AppliedPromotionId = promotion.Id
	_, err = app.cartRepo.Update(cart)
	if err != nil {
		slog.Error("Failed to apply promotion. Error: ", err)
		return nil, fmt.Errorf("failed to apply promotion, cartId: %d, promotionId: %d", cartId, promotionId)
	}

	return cart, nil
}

func calcCartPricesWithPromotion(cart *entity.Cart, items []*entity.Item, promotion *entity.Promotion) *entity.Cart {
	totalDiscount := float64(0)
	firstSellerID := items[0].SellerID
	sameSeller := true
	sameSellerTotalDiscount := float64(0)

	for _, item := range items {
		if promotion.PromotionType == entity.CategoryPromotion && item.CategoryID == promotion.CategoryP.CategoryID && item.ItemType == entity.DefaultItem {
			itemDiscount := item.Price * (float64(promotion.CategoryP.DiscountRate) / 100)
			totalDiscount += float64(item.Quantity) * itemDiscount
		} else if promotion.PromotionType == entity.SameSellerPromotion {
			if item.SellerID != firstSellerID {
				sameSeller = false
				break
			}

			itemDiscount := item.Price * (float64(promotion.SameSellerP.DiscountRate) / 100)
			sameSellerTotalDiscount += float64(item.Quantity) * itemDiscount
		}
	}

	if promotion.PromotionType == entity.TotalPricePromotion {
		for _, rnge := range promotion.TotalPriceP {
			if cart.TotalPrice >= rnge.PriceRangeStart && cart.TotalPrice <= rnge.PriceRangeEnd {
				totalDiscount = rnge.DiscountAmount
				break
			}
		}
	} else if promotion.PromotionType == entity.SameSellerPromotion && sameSeller {
		totalDiscount = sameSellerTotalDiscount
	}

	cart.TotalDiscount = totalDiscount
	cart.TotalAmount = cart.TotalPrice - totalDiscount

	return cart
}

func (app *cartApp) Display(cartId int) (*dto.DisplayCartDTO, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		return nil, fmt.Errorf("cart not found. CartID: %d", cartId)
	}

	items, itemErr := app.itemApp.ListByCartId(cartId)
	if itemErr != nil {
		return nil, fmt.Errorf("failed to retrieve cart items. Error: %v", itemErr)
	}

	var itemDTOs []*dto.ItemDTO
	for _, item := range items {
		vasItems, vasItemErr := app.vasItemApp.ListByItemId(item.Id)
		if vasItemErr != nil {
			return nil, fmt.Errorf("failed to retrieve vas items. Error: %v", vasItemErr)
		}

		itemDTOs = append(itemDTOs, &dto.ItemDTO{
			ID:         item.Id,
			CategoryID: item.CategoryID,
			SellerID:   item.SellerID,
			CartID:     item.CartID,
			Price:      item.Price,
			ItemType:   item.ItemType,
			VasItems:   vasItems,
		})
	}

	displayCartDTO := &dto.DisplayCartDTO{
		ID:                 cart.Id,
		TotalPrice:         cart.TotalPrice,
		TotalDiscount:      cart.TotalDiscount,
		TotalAmount:        cart.TotalAmount,
		AppliedPromotionId: cart.AppliedPromotionId,
		Items:              itemDTOs,
	}

	return displayCartDTO, nil
}

func (app *cartApp) ResetCart(cartId int) (*entity.Cart, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		return nil, fmt.Errorf("failed to reset cart: %v", err)
	}

	err = app.cartRepo.Delete(cartId)
	if err != nil {
		return nil, fmt.Errorf("failed to reset cart: %v", err)
	}

	items, err := app.itemApp.ListByCartId(cartId)
	if err != nil {
		return nil, fmt.Errorf("failed to reset cart: %v", err)
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

	cart.TotalAmount = 0
	cart.TotalDiscount = 0
	cart.TotalPrice = 0
	cart.AppliedPromotionId = 0
	app.cartRepo.Update(cart)
	// TODO: Rollback if any error occurs

	return cart, nil
}

func (app *cartApp) AddItem(cartId int, item *entity.Item) (*entity.Cart, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if cart == nil || err != nil {
		cart = &entity.Cart{
			Id: utils.GenerateID(),
		}
		_, err := app.cartRepo.Create(cart)
		if err != nil {
			return nil, fmt.Errorf("failed create cart error: %v", err)
		}
	}
	item.CartID = cartId

	itemValidateErr := item.Validate()
	if itemValidateErr != nil {
		return nil, fmt.Errorf("failed to add item: %v", itemValidateErr)
	}

	item, err = app.itemApp.Create(item)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %v", err)
	}

	cart, updCartErr := app.UpdateCartPriceAndQuantity(cart)
	if updCartErr != nil {
		app.itemApp.Delete(item.Id) // * Rollback
		slog.Error("Failed to update cart price and quantity. Error: ", updCartErr)
		return nil, fmt.Errorf("failed to update cart price and quantity. Error: %v", updCartErr)
	}

	return cart, nil
}

func (app *cartApp) UpdateCartPriceAndQuantity(cart *entity.Cart) (*entity.Cart, error) {
	items, listItemsErr := app.itemApp.ListByCartId(cart.Id)
	if listItemsErr != nil {
		return nil, fmt.Errorf("list items error: %v", listItemsErr)
	}

	var totalPrice float64
	var totalQuantity int

	for _, item := range items {
		totalQuantity += item.Quantity
		totalPrice += item.Price * float64(item.Quantity)
		vasItems, listVasItemErr := app.vasItemApp.ListByItemId(item.Id)
		if len(vasItems) > 0 && listVasItemErr == nil {
			for _, vasItem := range vasItems {
				totalPrice += vasItem.Price * float64(vasItem.Quantity)
			}
		}

	}

	cart.TotalPrice = totalPrice

	if totalQuantity > 30 {
		return nil, fmt.Errorf("total quantity can not be more than 30. Total Quantity: %d", totalQuantity)
	}

	if cart.AppliedPromotionId != 0 {
		promotion, getPromErr := app.promotionApp.GetById(cart.AppliedPromotionId)
		if getPromErr != nil {
			return nil, fmt.Errorf("get promotion error: %v", getPromErr)
		}
		cart = calcCartPricesWithPromotion(cart, items, promotion)
	} else {
		cart.TotalAmount = totalPrice
		cart.TotalDiscount = 0
	}

	if cart.TotalAmount > 500000 {
		return nil, fmt.Errorf("total amount can not be more than 500000. Total Amount: %f", cart.TotalAmount)
	}

	_, err := app.cartRepo.Update(cart)
	if err != nil {
		return nil, fmt.Errorf("update cart error: %v", err)
	}

	return cart, nil
}

func (app *cartApp) AddVasItem(vasItem *entity.VasItem) (*entity.Cart, error) {
	vasItem, err := app.vasItemApp.Create(vasItem)
	if err != nil {
		return nil, fmt.Errorf("failed to add vas item: %v", err)
	}

	cart, err := app.cartRepo.GetByID(vasItem.ItemId)
	if err != nil {
		app.vasItemApp.DeleteById(vasItem.Id) // * Rollback
		return nil, fmt.Errorf("cart not found. CartID: %d", vasItem.ItemId)
	}

	cart, err = app.UpdateCartPriceAndQuantity(cart)
	if err != nil {
		app.vasItemApp.DeleteById(vasItem.Id) // * Rollback
		return nil, fmt.Errorf("failed to update cart price and quantity. Error: %v", err)
	}

	return cart, nil
}
