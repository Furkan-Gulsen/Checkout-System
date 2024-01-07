package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/stretchr/testify/assert"
)

type mockCartRepository struct{}

var (
	applyPromotionCartRepo         func(cartId int, promotionId int) (*entity.Cart, error)
	displayCartRepo                func(cartId int) (*dto.DisplayCartDTO, error)
	resetCartRepo                  func(cartId int) error
	addItemRepo                    func(cartId int, item *entity.Item) (*entity.Item, error)
	updateCartPriceAndQuantityRepo func(cartId int) (*entity.Cart, error)
)

func (m *mockCartRepository) ApplyPromotion(cartId int, promotionId int) (*entity.Cart, error) {
	return applyPromotionCartRepo(cartId, promotionId)
}

func (m *mockCartRepository) Display(cartId int) (*dto.DisplayCartDTO, error) {
	return displayCartRepo(cartId)
}

func (m *mockCartRepository) ResetCart(cartId int) error {
	return resetCartRepo(cartId)
}

func (m *mockCartRepository) AddItem(cartId int, item *entity.Item) (*entity.Item, error) {
	return addItemRepo(cartId, item)
}

func (m *mockCartRepository) UpdateCartPriceAndQuantity(cartId int) (*entity.Cart, error) {
	return updateCartPriceAndQuantityRepo(cartId)
}

var cartAppMock CartAppInterface = &mockCartRepository{}

var (
	sameSalePromotion   *entity.Promotion
	categoryPromotion   *entity.Promotion
	totalPricePromotion *entity.Promotion
)

func init() {
	// * Same Seller Promotion
	promotion, err := PromotionAppMock.Create(&entity.Promotion{
		PromotionType: entity.PromotionType(1),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 10,
		},
	})
	if err != nil {
		panic(err)
	}

	sameSalePromotion = promotion

	// * Category Promotion
	promotion, err = PromotionAppMock.Create(&entity.Promotion{
		PromotionType: entity.PromotionType(2),
		CategoryP: &entity.CategoryPromotionDiscount{
			DiscountRate: 5,
			CategoryID:   3242,
		},
	})
	if err != nil {
		panic(err)
	}
	categoryPromotion = promotion

	// * Total Price Promotion
	promotion, err = PromotionAppMock.Create(&entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 500,
				PriceRangeEnd:   4999,
				DiscountAmount:  250,
			},
			{
				PriceRangeStart: 5000,
				PriceRangeEnd:   9999,
				DiscountAmount:  500,
			},
			{
				PriceRangeStart: 10000,
				PriceRangeEnd:   49999,
				DiscountAmount:  1000,
			},
			{
				PriceRangeStart: 50000,
				PriceRangeEnd:   500000,
				DiscountAmount:  2000,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	totalPricePromotion = promotion
}

func TestCreateDefaultCart_Success(t *testing.T) {
	cart := entity.Cart{
		Id:                 1,
		AppliedPromotionId: 0,
		TotalPrice:         0,
		TotalAmount:        0,
		TotalDiscount:      0,
	}
	applyPromotionCartRepo = func(cartId int, promotionId int) (*entity.Cart, error) {
		return &cart, nil
	}

	result, err := cartAppMock.ApplyPromotion(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &cart, result)
}
