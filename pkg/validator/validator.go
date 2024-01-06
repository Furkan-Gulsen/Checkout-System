package validator

type Validator interface {
	ValidateItemQuantity(quantity int) error
	ValidateItemPrice(price float64) error
	// ValidateDigitalItem
}

func NewValidator() Validator {
	return &defaultValidator{}
}
