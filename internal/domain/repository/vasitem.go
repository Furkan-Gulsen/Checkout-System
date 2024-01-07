package repository

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type VasItemRepositoryI interface {
	ListByItemId(int) ([]*entity.VasItem, error)
	GetById(int) (*entity.VasItem, error)
	Create(*entity.VasItem) (*entity.VasItem, error)
	DeleteById(int) error
}
