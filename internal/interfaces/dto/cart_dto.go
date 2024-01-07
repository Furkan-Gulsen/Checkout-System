package dto

type DisplayCartDTO struct {
	ID                 int        `json:"id"`
	TotalPrice         float64    `json:"totalPrice"`
	TotalDiscount      float64    `json:"totalDiscount"`
	TotalAmount        float64    `json:"totalAmount"`
	AppliedPromotionId int        `json:"appliedPromotionId" `
	Items              []*ItemDTO `json:"items"`
}
