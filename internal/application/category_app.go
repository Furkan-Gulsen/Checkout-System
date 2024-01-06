package application

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type categoryApp struct {
	repo repository.CategoryRepositoryI
}

var _ CategoryAppInterface = &categoryApp{}

type CategoryAppInterface interface {
	List() ([]entity.Category, error)
	Create(entity.Category) error
	GetByID(id int) (entity.Category, error)
}

func (app *categoryApp) List() ([]entity.Category, error) {
	return app.repo.List()
}

func (app *categoryApp) Create(category entity.Category) error {
	return app.repo.Create(category)
}

func (app *categoryApp) GetByID(id int) (entity.Category, error) {
	return app.repo.GetByID(id)
}
