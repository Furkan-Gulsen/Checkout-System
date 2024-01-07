package dto

import "github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"

type CategoryRequest struct {
	Name     string `json:"name"`
	ItemType int    `json:"itemType"`
}

func (request CategoryRequest) ToEntity() entity.Category {
	return entity.Category{
		Name:     request.Name,
		ItemType: entity.ItemType(request.ItemType),
	}
}
