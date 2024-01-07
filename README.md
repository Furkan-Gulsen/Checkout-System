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
- type: DefaultItem || DigitalItem

<!-- DigitalItem: max quantity 5, categoryId: 7889 -->
<!-- DefaultItem: VasItem price < DefaultItem price -->

<!-- TODO: VasItem direkt olarak item içerisinde de oluşturulabilir -->

### VasItem

- id: string (PK)
- categoryId: string
- sellerId: string
- price: number
- quantity: number (bundan emin değilim)

### Category

- id: string (PK)
- name: string
- itemType: enum (int)

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
- Jenkins

* Prometheus (vakit kalırsa)

## Mimari

```
cmd/app
- main.go

config
- config.yaml
- config.go

internal/domain
* product, promotion folder etc
  - entity.go (model ve value object değerleri)
  - repository.go (infrastructure)
  - service.go (application)

internal/infrastructure
* repository
  - product_repository.go
  - promotion_repository.go

internal/interfaces
* api
  - product_router.go
  - product_handler.go
  - promotion_router.go
  - promotion_handler.go
* middleware
  - error.go
  - success.go

internal/application
- cart_service.go
- promotion_service.go

pkg
- logger

tests
* unit
  * product
    - service_test.go
  * promotion
    - promotion_test.go
* intergration
  * api
    - product_test.go
    - promotion_test.go

```
