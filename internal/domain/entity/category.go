package entity

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type Category struct {
	Id   int    `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name" validate:"required"`
}

func (category Category) Validate() error {
	validate := validator.New()
	err := validate.Struct(category)

	if err == nil {
		return nil
	}

	var customValidationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		tag := err.Tag()
		structField := err.StructField()
		if tag == "required" {
			customValidationErrors = append(customValidationErrors, structField+" is required.")
		}
	}

	if len(customValidationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(customValidationErrors, " ")
		return errors.New(errorMessage)
	}

	return nil
}
