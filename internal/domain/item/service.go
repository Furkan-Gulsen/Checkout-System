package item

import (
	"context"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
)

type ItemService struct {
	itemRepository ItemRepository
}

func NewItemService(itemRepository ItemRepository) *ItemService {
	return &ItemService{
		itemRepository: itemRepository,
	}
}

func (p *ItemService) List() ([]entities.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.itemRepository.List(ctx)
}

func (p *ItemService) Create(item entities.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.itemRepository.Create(ctx, item)
}

func (p *ItemService) GetById(itemID int64) (entities.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.itemRepository.GetById(ctx, itemID)
}

func (p *ItemService) Delete(itemID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.itemRepository.Delete(ctx, itemID)
}
