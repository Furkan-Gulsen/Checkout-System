package validator

type Validator interface {
	ValidateItemQuantity(quantity int) error
	ValidateItemPrice(price float64) error
	// Diğer doğrulama metodları...
}

// NewValidator, yeni bir Validator nesnesi oluşturur.
func NewValidator() Validator {
	return &defaultValidator{}
}
