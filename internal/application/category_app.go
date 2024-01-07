package application

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

var _ CategoryAppInterface = &categoryApp{}

type categoryApp struct {
	categoryRepo repository.CategoryRepositoryI
}

func NewCategoryApp(categoryRepo repository.CategoryRepositoryI) *categoryApp {
	return &categoryApp{
		categoryRepo: categoryRepo,
	}
}

type CategoryAppInterface interface {
	List() ([]*entity.Category, error)
	Create(category *entity.Category) (*entity.Category, error)
	GetByID(id int) (*entity.Category, error)
}

func (app *categoryApp) List() ([]*entity.Category, error) {
	return app.categoryRepo.List()
}

func (app *categoryApp) Create(category *entity.Category) (*entity.Category, error) {
	return app.categoryRepo.Create(category)
}

func (app *categoryApp) GetByID(id int) (*entity.Category, error) {
	return app.categoryRepo.GetByID(id)
}
