package dto

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type ItemCreateRequest struct {
	CategoryID int             `json:"categoryId"`
	SellerID   int             `json:"sellerId"`
	CartID     int             `json:"cartId"`
	Price      float64         `json:"price"`
	Quantity   int             `json:"quantity"`
	ItemType   entity.ItemType `json:"itemType"`
}

func (dto ItemCreateRequest) ToEntity() *entity.Item {
	return &entity.Item{
		CategoryID: dto.CategoryID,
		SellerID:   dto.SellerID,
		CartID:     dto.CartID,
		Price:      dto.Price,
		Quantity:   dto.Quantity,
		ItemType:   dto.ItemType,
	}
}

type ItemDTO struct {
	ID         int               `json:"id"`
	CategoryID int               `json:"categoryId"`
	SellerID   int               `json:"sellerId"`
	CartID     int               `json:"cartId"`
	Price      float64           `json:"price"`
	ItemType   entity.ItemType   `json:"itemType"`
	VasItems   []*entity.VasItem `json:"vasItem"`
}
