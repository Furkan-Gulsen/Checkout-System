package entity

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type VasItem struct {
	Id         int     `json:"id" bson:"_id"`
	ItemId     int     `json:"itemId" bson:"itemId" validate:"required"`
	CategoryId int     `json:"categoryId" bson:"categoryId" validate:"required"`
	SellerId   int     `json:"sellerId" bson:"sellerId" validate:"required"`
	Price      float64 `json:"price" bson:"price" validate:"required"`
	Quantity   int     `json:"quantity" bson:"quantity" validate:"required,min=1,max=3"`
}

func (v *VasItem) Validate() error {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		var customValidationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				customValidationErrors = append(customValidationErrors, fmt.Sprintf("%s is required.", err.StructField()))
			} else if err.Tag() == "min" {
				customValidationErrors = append(customValidationErrors, fmt.Sprintf("%s must be greater than 0.", err.StructField()))
			} else if err.Tag() == "max" {
				customValidationErrors = append(customValidationErrors, fmt.Sprintf("%s must be less than 4.", err.StructField()))
			}
		}

		if len(customValidationErrors) > 0 {
			return errors.New("Validation errors: " + strings.Join(customValidationErrors, " "))
		}
	}

	return nil
}
