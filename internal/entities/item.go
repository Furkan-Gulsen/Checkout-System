package entities

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type Item struct {
	Id         int64     `json:"id" bson:"_id"`
	CategoryID int64     `json:"categoryId" bson:"categoryId" validate:"required"`
	SellerID   int64     `json:"sellerId" bson:"sellerId" validate:"required"`
	Price      float64   `json:"price" bson:"price" validate:"required"`
	Quantity   int64     `json:"quantity" bson:"quantity" validate:"required,max=10"`
	VasItems   []VasItem `json:"vasItems,omitempty" bson:"vasItems,omitempty"`
}

// Validate function validates the Item structure and handles errors.
func (item Item) Validate() error {
	validate := validator.New()
	err := validate.Struct(item)

	// If validation is successful, return nil error
	if err == nil {
		return nil
	}

	// Custom validation error messages
	var customValidationErrors []string
	// Process errors returned by validator
	for _, err := range err.(validator.ValidationErrors) {
		tag := err.Tag()
		structField := err.StructField()
		// Required field check
		if tag == "required" {
			customValidationErrors = append(customValidationErrors, structField+" is required.")
		} else if tag == "max" {
			customValidationErrors = append(customValidationErrors, structField+" must be less than or equal to 10.")
		}
	}

	// Combine custom errors and return a new error
	if len(customValidationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(customValidationErrors, " ")
		return errors.New(errorMessage)
	}

	return nil
}
