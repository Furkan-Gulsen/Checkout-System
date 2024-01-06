package repository

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type PromotionRepositoryI interface {
	List() ([]*entity.Promotion, error)
	Create(promotion *entity.Promotion) error
	GetById(promotionID int) (*entity.Promotion, error)
}
