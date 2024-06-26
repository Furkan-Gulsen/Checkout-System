package repository

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
)

type CategoryRepositoryI interface {
	List() ([]*entity.Category, error)
	Create(category *entity.Category) (*entity.Category, error)
	GetByID(id int) (*entity.Category, error)
}
