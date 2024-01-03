package entities

type Item struct {
	ID         string    `json:"id" bson:"_id"`
	CategoryId string    `json:"categoryId" bson:"categoryId"`
	SellerId   string    `json:"sellerId" bson:"sellerId"`
	Price      float64   `json:"price" bson:"price"`
	Quantity   int       `json:"quantity" bson:"quantity"`
	VasItems   []VasItem `json:"vasItems" bson:"vasItems"`
}
