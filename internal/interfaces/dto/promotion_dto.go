package dto

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type PromotionRequest struct {
	PromotionType int                                   `json:"promotionType"`
	SameSellerP   *entity.SameSellerPromotionDiscount   `json:"sameSellerPromotion"`
	CategoryP     *entity.CategoryPromotionDiscount     `json:"categoryPromotion"`
	TotalPriceP   []*entity.TotalPricePromotionDiscount `json:"totalPricePromotions"`
}

func (request PromotionRequest) ToEntity() entity.Promotion {
	return entity.Promotion{
		PromotionType: entity.PromotionType(request.PromotionType),
		SameSellerP:   request.SameSellerP,
		CategoryP:     request.CategoryP,
		TotalPriceP:   request.TotalPriceP,
	}
}
