package entities

type VasItem struct {
	ID         int64   `json:"id" bson:"_id"`
	CategoryId int64   `json:"categoryId" bson:"categoryId"`
	SellerId   int64   `json:"sellerId" bson:"sellerId"`
	Price      float64 `json:"price" bson:"price"`
	Quantity   int64   `json:"quantity" bson:"quantity"`
}
