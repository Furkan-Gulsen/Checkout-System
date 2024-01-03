# Checkout-System

## Entities

### Cart

- id: string (PK)
<!-- TODO: items içerisinde vasItem ne işe yarıyor bakılacak -->
- items: (DigitalItem || DefaultItem)[] (Max 10 unique, total quantity max 30)
- totalAmount: number [totalPrice - totalDiscount] (max 500k) (Emin değilim, kalkabilir)
- totalPrice: number (tüm ürünler + vas items ürünlerin fiyatı)
- totalDiscount: number
- appliedPromotions: Promotion[]

### Item

- id: string (PK)
- categoryId: string (FK)
- sellerId: string (FK)
- price: number
- quantity: number (max 10)
- vasItems: VasItem[]
<!-- - type: VasItem || DefaultItem || DigitalItem -->

### VasItem

- id: string (PK)
- categoryId: string
- sellerId: string
- price: number
- quantity: number (bundan emin değilim)

### Promotion

- id: string (PK)
- discountRate?: number
- relatedCategoryId?: string (CategoryPromotion)
- minCartTotal?: number (totalPricePromotion)
- discountRates?: Dict<number, number> (TotalPricePromotion)
- promotionType: SameSellerPromotion | CategoryPromotion | TotalPricePromotion (enum)

Mikroservisler:

- Checkout (Cart, Promotion)
- Product (Product, Category)

## Kullanılacak Teknolojiler:

- Go
- MongoDB
- Docker
- Swagger
- Notification system (vakit kalırsa)
- GitHub Actions (vakit kalırsa)
- Prometheus (vakit kalırsa)
- Api Gateway (vakit kalırsa)
