package repository

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type CartRepositoryI interface {
	Create(cart *entity.Cart) (*entity.Cart, error)
	GetByID(id int) (*entity.Cart, error)
	Update(cart *entity.Cart) (*entity.Cart, error)
	Delete(id int) error
}
