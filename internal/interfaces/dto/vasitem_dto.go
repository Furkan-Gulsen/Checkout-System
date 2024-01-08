package dto

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type VasItemCreateRequest struct {
	ItemId     int     `json:"itemId"`
	CategoryId int     `json:"categoryId"`
	SellerId   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

func (dto VasItemCreateRequest) ToEntity() *entity.VasItem {
	return &entity.VasItem{
		ItemId:     dto.ItemId,
		CategoryId: dto.CategoryId,
		SellerId:   dto.SellerId,
		Price:      dto.Price,
		Quantity:   dto.Quantity,
	}
}
