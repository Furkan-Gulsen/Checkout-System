package entities

type Cart struct {
	ID                string      `json:"id" bson:"_id"`
	Items             []Item      `json:"items" bson:"items"`
	TotalAmount       float64     `json:"totalAmount" bson:"totalAmount"`
	TotalPrice        float64     `json:"totalPrice" bson:"totalPrice"`
	TotalDiscount     float64     `json:"totalDiscount" bson:"totalDiscount"`
	AppliedPromotions []Promotion `json:"appliedPromotions" bson:"appliedPromotions"`
}
