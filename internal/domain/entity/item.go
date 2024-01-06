package entity

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type ItemType int

const (
	DigitalItem PromotionType = iota + 1
	DefaultItem
)

type Item struct {
	Id         int      `json:"id" bson:"_id"`
	CategoryID int      `json:"categoryId" bson:"categoryId" validate:"required"`
	SellerID   int      `json:"sellerId" bson:"sellerId" validate:"required"`
	CartID     int      `json:"cartId" bson:"cartId" validate:"required"`
	Price      float64  `json:"price" bson:"price" validate:"required"`
	Quantity   int      `json:"quantity" bson:"quantity" validate:"required,max=10"`
	ItemType   ItemType `json:"itemType,omitempty" bson:"itemType,omitempty"`
}

func (item Item) Validate() error {
	validate := validator.New()
	err := validate.Struct(item)

	if err == nil {
		return nil
	}

	var customValidationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		tag := err.Tag()
		structField := err.StructField()
		if tag == "required" {
			customValidationErrors = append(customValidationErrors, structField+" is required.")
		} else if tag == "max" {
			customValidationErrors = append(customValidationErrors, structField+" must be less than or equal to 10.")
		}
	}

	if len(customValidationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(customValidationErrors, " ")
		return errors.New(errorMessage)
	}

	return nil
}
