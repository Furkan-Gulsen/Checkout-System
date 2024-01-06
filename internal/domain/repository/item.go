package repository

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
)

type ItemRepositoryI interface {
	ListByCartId(cartId int) ([]*entity.Item, error)
	Create(item *entity.Item) error
	GetById(itemID int) (*entity.Item, error)
	Delete(itemID int) error
}
