package application

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type PromotionAppInterface interface {
	List() ([]*entity.Promotion, error)
	Create(*entity.Promotion) error
	GetById(int) (*entity.Promotion, error)
}

type promotionApp struct {
	promotionRepo repository.PromotionRepositoryI
}

func NewPromotionApp(promotionRepo repository.PromotionRepositoryI) *promotionApp {
	return &promotionApp{
		promotionRepo: promotionRepo,
	}
}

func (app *promotionApp) List() ([]*entity.Promotion, error) {
	return app.promotionRepo.List()
}

func (app *promotionApp) Create(promotion *entity.Promotion) error {
	return app.promotionRepo.Create(promotion)
}

func (app *promotionApp) GetById(promotionID int) (*entity.Promotion, error) {
	return app.promotionRepo.GetById(promotionID)
}
