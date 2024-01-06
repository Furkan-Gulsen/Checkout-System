package dto

type DisplayCartResponse struct {
	ID            int         `json:"id"`
	TotalPrice    float64     `json:"totalPrice"`
	TotalDiscount float64     `json:"totalDiscount"`
	TotalAmount   float64     `json:"totalAmount"`
	Items         []Item      `json:"items"`
	Promotions    []Promotion `json:"promotions"`
	VasItems      []VasItem   `json:"vasItems"`
}

// Item, Cart içindeki ürünler için DTO.
type Item struct {
	// Item ile ilgili alanlar...
}

// Promotion, Cart içindeki promosyonlar için DTO.
type Promotion struct {
	// Promotion ile ilgili alanlar...
}

// VasItem, Cart içindeki value-added service item'lar için DTO.
type VasItem struct {
	// VasItem ile ilgili alanlar...
}
