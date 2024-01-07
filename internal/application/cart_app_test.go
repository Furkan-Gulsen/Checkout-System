package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
)

type mockCartRepository struct{}

var (
	createCartRepo  func(cart *entity.Cart) (*entity.Cart, error)
	getByIDCartRepo func(id int) (*entity.Cart, error)
	updateCartRepo  func(cart *entity.Cart) (*entity.Cart, error)
	deleteCartRepo  func(id int) error
)

func (m *mockCartRepository) Create(cart *entity.Cart) (*entity.Cart, error) {
	return createCartRepo(cart)
}

func (m *mockCartRepository) GetByID(id int) (*entity.Cart, error) {
	return getByIDCartRepo(id)
}

func (m *mockCartRepository) Update(cart *entity.Cart) (*entity.Cart, error) {
	return updateCartRepo(cart)
}

func (m *mockCartRepository) Delete(id int) error {
	return deleteCartRepo(id)
}

var CartAppMock repository.CartRepositoryI = &mockCartRepository{}

func TestAllForCart(t *testing.T) {
	t.Run("TestSaveCartToUncreatedCart_Success", func(t *testing.T) {
		saveCartToUncreatedCart_Success(t)
	})
	t.Run("TestSaveCartToCreatedCart_Success", func(t *testing.T) {
		saveCartToCreatedCart_Success(t)
	})
	t.Run("TestAddItemToCart_Fail", func(t *testing.T) {
		addItemToCart_Fail(t)
	})
	t.Run("TestAddingItemsToCartWithPromotion_Success", func(t *testing.T) {
		addingItemsToCartWithPromotion_Success(t)
	})
	t.Run("TestAddingItemsToCartWithPromotion_NotApplied", func(t *testing.T) {
		addingItemsToCartWithPromotion_NotApplied(t)
	})
	t.Run("TestAddingItemsToCartWithCategoryPromotion_Success", func(t *testing.T) {
		addingItemsToCartWithCategoryPromotion_Success(t)
	})
	t.Run("TestAddingItemsToCartWithTotalPricePromotion_Success", func(t *testing.T) {
		addingItemsToCartWithTotalPricePromotion_Success(t)
	})
	t.Run("TestApplySameSellerPromotion", func(t *testing.T) {
		applySameSellerPromotion(t)
	})
	t.Run("TestApplySameCategoryPromotion", func(t *testing.T) {
		applyCategoryPromotion(t)
	})
	t.Run("TestApplyTotalPricePromotion", func(t *testing.T) {
		applyTotalPricePromotion(t)
	})
}

// * Test of Adding Items to an Uncreated Cart
func saveCartToUncreatedCart_Success(t *testing.T) {
	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:            5463,
			TotalAmount:   0,
			TotalPrice:    0,
			TotalDiscount: 0,
		}, nil
	}

	createItemRepo = func(item *entity.Item) (*entity.Item, error) {
		return &entity.Item{
			Id:         1003,
			CategoryID: 1111,
			SellerID:   1111,
			Price:      100,
			Quantity:   5,
			CartID:     5463,
			ItemType:   entity.DigitalItem,
		}, nil
	}

	item := &entity.Item{
		CategoryID: 1003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DigitalItem,
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return []*entity.VasItem{
			{
				Id:         3456,
				CategoryId: 3242,
				ItemId:     2222,
				SellerId:   5003,
				Price:      340,
				Quantity:   1,
			},
		}, nil
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			{
				Id:         1001,
				CategoryID: 3003,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         1002,
				CategoryID: 3003,
				SellerID:   2222,
				Price:      200,
				Quantity:   10,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			}, item,
		}, nil
	}

	updateCartRepo = func(cart *entity.Cart) (*entity.Cart, error) {
		return &entity.Cart{
			Id:            5463, // * not important
			TotalAmount:   0,    // * not important
			TotalPrice:    0,    // * not important
			TotalDiscount: 0,    // * not important
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(1111, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(4020), cart.TotalAmount)
	// assert.Equal(t, float64(4020), cart.TotalPrice)
	// assert.Equal(t, float64(0), cart.TotalDiscount)
	// assert.Equal(t, 0, cart.AppliedPromotionId)

}

// * Test of Adding Items to a Created Cart
func saveCartToCreatedCart_Success(t *testing.T) {
	item := &entity.Item{
		CategoryID: 1003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DigitalItem,
	}

	createCartRepo = func(cart *entity.Cart) (*entity.Cart, error) {
		return nil, nil
	}

	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:                 1004,
			TotalAmount:        1000,
			TotalPrice:         1000,
			TotalDiscount:      0,
			AppliedPromotionId: 0,
		}, nil
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
		}, nil
	}

	listVasItemRepo = func(itemId int) ([]*entity.VasItem, error) {
		return []*entity.VasItem{}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(1111, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(500), cart.TotalAmount)
	// assert.Equal(t, float64(500), cart.TotalPrice)
	// assert.Equal(t, float64(0), cart.TotalDiscount)
	// assert.Equal(t, 0, cart.AppliedPromotionId)
	// fmt.Println("1003. TotalAmount: ", cart.TotalAmount)
}

// * Testing of a maximum of 30 products in the cart
func addItemToCart_Fail(t *testing.T) {
	item := &entity.Item{
		Id:         5000,
		CategoryID: 3003,
		SellerID:   1111,
		Price:      100,
		Quantity:   500,
		CartID:     5463,
		ItemType:   entity.DefaultItem,
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
		}, nil
	}

	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:                 1004,
			TotalAmount:        0,
			TotalPrice:         0,
			TotalDiscount:      0,
			AppliedPromotionId: 1234,
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(5463, item)
	// assert.Nil(t, cart)
	// assert.NotNil(t, err)
	// assert.Equal(t, "failed to add item: Validation errors: Quantity max value is 10.", err.Error())
}

// * Test of Adding Items to a Created Cart with SameSellerPromotion [Applied]
func addingItemsToCartWithPromotion_Success(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1234,
			PromotionType: entity.PromotionType(1),
			SameSellerP: &entity.SameSellerPromotionDiscount{
				DiscountRate: 20,
			},
		}, nil
	}

	item := &entity.Item{
		Id:         5001,
		CategoryID: 3003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DefaultItem,
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(5001, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(400), cart.TotalAmount)
	// assert.Equal(t, float64(500), cart.TotalPrice)
	// assert.Equal(t, float64(100), cart.TotalDiscount)
	// assert.Equal(t, 1234, cart.AppliedPromotionId)
}

// * Test of Adding Items to a Created Cart with SameSellerPromotion [Not Applied]
func addingItemsToCartWithPromotion_NotApplied(t *testing.T) {
	item := &entity.Item{
		Id:         5005,
		CategoryID: 3003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DefaultItem,
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
			{
				Id:         5002,
				CategoryID: 3003,
				SellerID:   2222,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(5002, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(1000), cart.TotalAmount)
	// assert.Equal(t, float64(1000), cart.TotalPrice)
	// assert.Equal(t, float64(0), cart.TotalDiscount)
	// assert.Equal(t, 1234, cart.AppliedPromotionId)
}

// * Test of Adding Items to a Created Cart with CategoryPromotion [Applied]
func addingItemsToCartWithCategoryPromotion_Success(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1235,
			PromotionType: entity.PromotionType(2),
			CategoryP: &entity.CategoryPromotionDiscount{
				CategoryID:   3003,
				DiscountRate: 20,
			},
		}, nil
	}

	item := &entity.Item{
		Id:         5003,
		CategoryID: 3003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DefaultItem,
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
			{
				Id:         5004,
				CategoryID: 4003, // * different category
				SellerID:   2222,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5005,
				CategoryID: 3003,
				SellerID:   3333,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5006,
				CategoryID: 3003,
				SellerID:   4444,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DigitalItem, // * not included in promotion
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(5003, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(1800), cart.TotalAmount)
	// assert.Equal(t, float64(2000), cart.TotalPrice)
	// assert.Equal(t, float64(200), cart.TotalDiscount)
}

// * Test of Adding Items to a Created Cart with TotalPricePromotion [Applied]
func addingItemsToCartWithTotalPricePromotion_Success(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1236,
			PromotionType: entity.PromotionType(3),
			TotalPriceP: []*entity.TotalPricePromotionDiscount{
				{
					PriceRangeStart: 1000,
					PriceRangeEnd:   2000,
					DiscountAmount:  100,
				},
			},
		}, nil
	}

	item := &entity.Item{
		Id:         5007,
		CategoryID: 3003,
		SellerID:   1111,
		Price:      100,
		Quantity:   5,
		CartID:     5463,
		ItemType:   entity.DefaultItem,
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			item,
			{
				Id:         5008,
				CategoryID: 4003,
				SellerID:   2222,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5009,
				CategoryID: 3003,
				SellerID:   3333,
				Price:      100,
				Quantity:   5,
				CartID:     5463,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.AddItem(5007, item)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(1400), cart.TotalAmount)
	// assert.Equal(t, float64(1500), cart.TotalPrice)
	// assert.Equal(t, float64(100), cart.TotalDiscount)
}

func applySameSellerPromotion(t *testing.T) {
	const promotionId = 1234
	const cartId = 1501

	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            promotionId,
			PromotionType: entity.PromotionType(1),
			SameSellerP: &entity.SameSellerPromotionDiscount{
				DiscountRate: 50,
			},
		}, nil
	}

	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:                 cartId,
			TotalAmount:        1000,
			TotalPrice:         1000,
			TotalDiscount:      0,
			AppliedPromotionId: 0,
		}, nil
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			{
				Id:         5001,
				CategoryID: 3003,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5002,
				CategoryID: 3003,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.ApplyPromotion(cartId, promotionId)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(500), cart.TotalAmount)
	// assert.Equal(t, float64(1000), cart.TotalPrice)
	// assert.Equal(t, float64(500), cart.TotalDiscount)
	// assert.Equal(t, 1234, cart.AppliedPromotionId)
}

func applyCategoryPromotion(t *testing.T) {
	const promotionId = 1234
	const cartId = 1501
	const categoryId = 3001

	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            promotionId,
			PromotionType: entity.PromotionType(2),
			CategoryP: &entity.CategoryPromotionDiscount{
				CategoryID:   categoryId,
				DiscountRate: 50,
			},
		}, nil
	}

	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:                 cartId,
			TotalAmount:        1000,
			TotalPrice:         1000,
			TotalDiscount:      0,
			AppliedPromotionId: 0,
		}, nil
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			{
				Id:         5001,
				CategoryID: categoryId,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5002,
				CategoryID: 3003,
				SellerID:   1111,
				Price:      100,
				Quantity:   5,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.ApplyPromotion(cartId, promotionId)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(750), cart.TotalAmount)
	// assert.Equal(t, float64(1000), cart.TotalPrice)
	// assert.Equal(t, float64(250), cart.TotalDiscount)
	// assert.Equal(t, 1234, cart.AppliedPromotionId)
}

func applyTotalPricePromotion(t *testing.T) {
	const promotionId = 1234
	const cartId = 1501
	const categoryId = 3001

	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            promotionId,
			PromotionType: entity.PromotionType(3),
			TotalPriceP: []*entity.TotalPricePromotionDiscount{
				{
					PriceRangeStart: 1000,
					PriceRangeEnd:   2000,
					DiscountAmount:  200,
				},
			},
		}, nil
	}

	getByIDCartRepo = func(id int) (*entity.Cart, error) {
		return &entity.Cart{
			Id:                 cartId,
			TotalAmount:        1500,
			TotalPrice:         1500,
			TotalDiscount:      0,
			AppliedPromotionId: 0,
		}, nil
	}

	listItemRepo = func(id int) ([]*entity.Item, error) {
		return []*entity.Item{
			{
				Id:         5001,
				CategoryID: categoryId,
				SellerID:   1111,
				Price:      100,
				Quantity:   7,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
			{
				Id:         5002,
				CategoryID: 3003,
				SellerID:   1111,
				Price:      100,
				Quantity:   8,
				CartID:     cartId,
				ItemType:   entity.DefaultItem,
			},
		}, nil
	}

	// app := NewCartApp(CartAppMock, ItemAppMock, VasItemAppMock, PromotionAppMock)
	// cart, err := app.ApplyPromotion(cartId, promotionId)
	// assert.Nil(t, err)
	// assert.NotNil(t, cart)
	// assert.Equal(t, float64(1300), cart.TotalAmount)
	// assert.Equal(t, float64(1500), cart.TotalPrice)
	// assert.Equal(t, float64(200), cart.TotalDiscount)
	// assert.Equal(t, 1234, cart.AppliedPromotionId)
}

// func ResetCart(t *testing.T) {
// 	getByIDCartRepo = func(id int) (*entity.Cart, error) {
// 		return &entity.Cart{
// 			Id:                 1501,
// 			TotalAmount:        0,
// 			TotalPrice:         0,
// 			TotalDiscount:      0,
// 			AppliedPromotionId: 0,
// 		}, nil
// 	}

// 	updateCartRepo = func(cart *entity.Cart) (*entity.Cart, error) {
// 		return &entity.Cart{
// 			Id:                 1501,
// 			TotalAmount:        0,
// 			TotalPrice:         0,
// 			TotalDiscount:      0,
// 			AppliedPromotionId: 0,
// 		}, nil
// 	}

// 	deleteItemRepo = func(id int) error {
// 		return nil
// 	}

// 	deleteVasItemRepo = func(id int) error {
// 		return nil
// 	}

// 	deleteCartRepo = func(id int) error {
// 		return nil
// 	}

// }
