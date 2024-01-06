package entity

// PromotionType enum for different promotion types
type PromotionType int

const (
	SameSellerPromotion PromotionType = iota
	CategoryPromotion
	TotalPricePromotion
)

// DiscountRates struct for TotalPricePromotion discounts
type DiscountRates struct {
	PriceRangeStart float64
	PriceRangeEnd   float64
	DiscountAmount  float64
}

// Promotion struct representing a promotion
type Promotion struct {
	ID                int             `json:"id" bson:"_id"`
	DiscountRate      float64         `json:"discountRate" bson:"discountRate"`
	RelatedCategoryId int             `json:"relatedCategoryId" bson:"relatedCategoryId"`
	MinCartTotal      float64         `json:"minCartTotal" bson:"minCartTotal"`
	DiscountRates     []DiscountRates `json:"discountRates" bson:"discountRates"`
	PromotionType     PromotionType   `json:"promotionType" bson:"promotionType"`
}
