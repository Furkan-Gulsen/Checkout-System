package validator

import "fmt"

type defaultValidator struct{}

func (v *defaultValidator) ValidateItemQuantity(quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if quantity > 10 {
		return fmt.Errorf("quantity must not exceed 10")
	}
	return nil
}

func (v *defaultValidator) ValidateItemPrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("price must not be negative")
	}
	return nil
}
