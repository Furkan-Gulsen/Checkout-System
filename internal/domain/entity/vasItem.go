package entity

type VasItem struct {
	ID         int     `json:"id" bson:"_id"`
	CategoryId int     `json:"categoryId" bson:"categoryId"`
	SellerId   int     `json:"sellerId" bson:"sellerId"`
	Price      float64 `json:"price" bson:"price"`
	Quantity   int     `json:"quantity" bson:"quantity"`
}
