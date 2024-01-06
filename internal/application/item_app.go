package application

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type itemApp struct {
	repo repository.ItemRepositoryI
}

var _ ItemAppInterface = &itemApp{}

type ItemAppInterface interface {
	List() ([]*entity.Item, error)
	Create(*entity.Item) error
	GetById(int) (*entity.Item, error)
	Delete(int) error
}

func (app *itemApp) List() ([]*entity.Item, error) {
	return app.repo.List()
}

func (app *itemApp) Create(item *entity.Item) error {
	return app.repo.Create(item)
}

func (app *itemApp) GetById(itemID int) (*entity.Item, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	return app.repo.GetById(itemID)
}

func (app *itemApp) Delete(itemID int) error {
	return app.repo.Delete(itemID)
}
