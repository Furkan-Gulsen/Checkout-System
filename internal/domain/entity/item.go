package entity

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type ItemType int

const (
	DigitalItem ItemType = iota + 1
	DefaultItem
)

type Item struct {
	Id         int      `json:"id" bson:"_id"`
	CategoryID int      `json:"categoryId" bson:"categoryId" validate:"required"`
	SellerID   int      `json:"sellerId" bson:"sellerId" validate:"required"`
	CartID     int      `json:"cartId" bson:"cartId" validate:"required"`
	Price      float64  `json:"price" bson:"price" validate:"required"`
	Quantity   int      `json:"quantity" bson:"quantity" validate:"required,max=10"`
	ItemType   ItemType `json:"itemType" bson:"itemType" validate:"oneof=1 2"`
}

func (item Item) Validate() error {
	validate := validator.New()
	err := validate.Struct(item)

	if err == nil {
		return nil
	}

	var customValidationErrors []string

	for _, err := range err.(validator.ValidationErrors) {
		if err.Tag() == "required" {
			customValidationErrors = append(customValidationErrors, err.Field()+" is required.")
		} else if err.Tag() == "max" {
			customValidationErrors = append(customValidationErrors, err.Field()+" max value is "+err.Param()+".")
		} else if err.Tag() == "oneof" {
			customValidationErrors = append(customValidationErrors, err.Field()+" must be one of "+err.Param()+".")
		}
	}

	if len(customValidationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(customValidationErrors, " ")
		return errors.New(errorMessage)
	}

	return nil
}
