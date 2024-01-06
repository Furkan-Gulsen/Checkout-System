package entity

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type PromotionType int

const (
	SameSellerPromotion PromotionType = iota
	CategoryPromotion
	TotalPricePromotion
)

type TotalPricePromotionDiscount struct {
	PriceRangeStart float64
	PriceRangeEnd   float64
	DiscountAmount  float64
}

type SameSellerPromotionDiscount struct {
	DiscountRate float64
}

type CategoryPromotionDiscount struct {
	DiscountRate float64
	CategoryID   int
}

type Promotion struct {
	ID            int           `json:"id" bson:"_id"`
	PromotionType PromotionType `json:"promotionType" bson:"promotionType" validate:"required"`
	Discount      interface{}   `json:"discount" bson:"discount" validate:"dive,required"`
}

func (p *Promotion) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		var customValidationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				customValidationErrors = append(customValidationErrors, fmt.Sprintf("%s is required.", err.StructField()))
			}
		}

		if len(customValidationErrors) > 0 {
			return errors.New("Validation errors: " + strings.Join(customValidationErrors, " "))
		}
	}

	switch p.PromotionType {
	case SameSellerPromotion:
		_, ok := p.Discount.(SameSellerPromotionDiscount)
		if !ok {
			return errors.New("discount is not SameSellerPromotionDiscount")
		}
	case CategoryPromotion:
		_, ok := p.Discount.(CategoryPromotionDiscount)
		if !ok {
			return errors.New("discount is not CategoryPromotionDiscount")
		}
	case TotalPricePromotion:
		_, ok := p.Discount.(TotalPricePromotionDiscount)
		if !ok {
			return errors.New("discount is not TotalPricePromotionDiscount")
		}
	}

	return nil
}
