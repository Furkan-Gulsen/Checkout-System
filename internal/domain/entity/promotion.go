package entity

import (
	"fmt"

	"github.com/go-playground/validator"
)

type PromotionType int

const (
	SameSellerPromotion PromotionType = iota + 1
	CategoryPromotion
	TotalPricePromotion
)

type SameSellerPromotionDiscount struct {
	DiscountRate float64 `json:"discountRate" bson:"discountRate" validate:"required,min=0,max=100"`
}

type CategoryPromotionDiscount struct {
	DiscountRate float64 `json:"discountRate" bson:"discountRate" validate:"required,min=0,max=100"`
	CategoryID   int     `json:"categoryID" bson:"categoryID" validate:"required"`
}

type TotalPricePromotionDiscount struct {
	PriceRangeStart float64 `json:"priceRangeStart" bson:"priceRangeStart" validate:"required"`
	PriceRangeEnd   float64 `json:"priceRangeEnd" bson:"priceRangeEnd" validate:"required"`
	DiscountAmount  float64 `json:"discountAmount" bson:"discountAmount" validate:"required"`
}

type Promotion struct {
	Id            int                            `json:"id" bson:"_id"`
	PromotionType PromotionType                  `json:"promotionType" bson:"promotionType" validate:"required,oneof=1 2 3"`
	SameSellerP   *SameSellerPromotionDiscount   `json:"sameSellerPromotion,omitempty" bson:"sameSellerPromotion,omitempty"`
	CategoryP     *CategoryPromotionDiscount     `json:"categoryPromotion,omitempty" bson:"categoryPromotion,omitempty"`
	TotalPriceP   []*TotalPricePromotionDiscount `json:"totalPricePromotions,omitempty" bson:"totalPricePromotions,omitempty"`
}

func (p *Promotion) Validate() error {
	// * Check if only one of sameSellerP, categoryP or totalPriceP exists
	if (p.SameSellerP != nil && p.CategoryP != nil) ||
		(p.SameSellerP != nil && p.TotalPriceP != nil) ||
		(p.CategoryP != nil && p.TotalPriceP != nil) {
		return fmt.Errorf("only one of sameSellerP, categoryP, or totalPriceP can be set")
	}

	// * Check if promotion type is valid
	validate := validator.New()
	switch p.PromotionType {
	case SameSellerPromotion:
		if p.SameSellerP == nil {
			return fmt.Errorf("sameSellerP is required")
		}
		if err := validate.Struct(p.SameSellerP); err != nil {
			return translateValidationErrors(err)
		}

	case CategoryPromotion:
		if p.CategoryP == nil {
			return fmt.Errorf("categoryP is required")
		}
		if err := validate.Struct(p.CategoryP); err != nil {
			return translateValidationErrors(err)
		}

	case TotalPricePromotion:
		if p.TotalPriceP == nil {
			return fmt.Errorf("totalPriceP is required")
		} else if len(p.TotalPriceP) == 0 {
			return fmt.Errorf("totalPriceP must have at least one element")
		}

		for _, totalPriceP := range p.TotalPriceP {
			if err := validate.Struct(totalPriceP); err != nil {
				return translateValidationErrors(err)
			}
		}

	default:
		return fmt.Errorf("invalid promotion type")
	}

	return nil
}

func translateValidationErrors(err error) error {
	for _, validationErr := range err.(validator.ValidationErrors) {
		switch validationErr.Tag() {
		case "required":
			return fmt.Errorf("%s is required", validationErr.StructField())
		case "min":
			return fmt.Errorf("quantity must be greater than or equal to 0")
		case "max":
			return fmt.Errorf("quantity must be less than or equal to 100")
		}
	}
	return err
}
