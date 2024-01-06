package entity

type Cart struct {
	Id                 int     `json:"id" bson:"_id"`
	TotalAmount        float64 `json:"totalAmount" bson:"totalAmount"`
	TotalPrice         float64 `json:"totalPrice" bson:"totalPrice"`
	TotalDiscount      float64 `json:"totalDiscount" bson:"totalDiscount"`
	AppliedPromotionId int     `json:"appliedPromotionId" bson:"appliedPromotionId"`
}
