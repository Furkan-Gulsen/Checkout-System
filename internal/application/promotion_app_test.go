package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

type mockPromotionRepository struct{}

var (
	listPromotionRepo    func() ([]*entity.Promotion, error)
	createPromotionRepo  func(promotion *entity.Promotion) (*entity.Promotion, error)
	getByIDPromotionRepo func(promotionID int) (*entity.Promotion, error)
)

func (m *mockPromotionRepository) List() ([]*entity.Promotion, error) {
	return listPromotionRepo()
}

func (m *mockPromotionRepository) Create(promotion *entity.Promotion) (*entity.Promotion, error) {
	return createPromotionRepo(promotion)
}

func (m *mockPromotionRepository) GetById(promotionID int) (*entity.Promotion, error) {
	return getByIDPromotionRepo(promotionID)
}

var PromotionAppMock repository.PromotionRepositoryI = &mockPromotionRepository{}

func TestGetPromotionByID_Success(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1,
			PromotionType: entity.PromotionType(1),
		}, nil
	}

	app := NewPromotionApp(PromotionAppMock)
	promotion, err := app.GetById(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, promotion.Id)
	assert.Equal(t, entity.PromotionType(1), promotion.PromotionType)
}

func TestListPromotion_Success(t *testing.T) {
	listPromotionRepo = func() ([]*entity.Promotion, error) {
		return []*entity.Promotion{
			{
				Id:            1,
				PromotionType: entity.PromotionType(1),
			},
			{
				Id:            2,
				PromotionType: entity.PromotionType(2),
			},
		}, nil
	}

	app := NewPromotionApp(PromotionAppMock)
	promotions, err := app.List()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(promotions))
	assert.Equal(t, 1, promotions[0].Id)
	assert.Equal(t, entity.PromotionType(1), promotions[0].PromotionType)
	assert.Equal(t, 2, promotions[1].Id)
	assert.Equal(t, entity.PromotionType(2), promotions[1].PromotionType)
}

func TestListPromotion_Fail(t *testing.T) {
	listPromotionRepo = func() ([]*entity.Promotion, error) {
		return nil, nil
	}

	app := NewPromotionApp(PromotionAppMock)
	promotions, err := app.List()
	assert.Nil(t, promotions)
	assert.Nil(t, err)
}

func TestSameSellerPromotionDiscount_Validate_Success(t *testing.T) {
	sameSellerPromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(1),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 10,
		},
	}

	err := sameSellerPromotion.Validate()
	assert.Nil(t, err)
}

func TestSameSellerPromotionDiscount_Validate_Fail(t *testing.T) {
	sameSellerPromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(1),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 0,
		},
	}
	err := sameSellerPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountRate is required", err.Error())

	sameSellerPromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(1),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 150,
		},
	}
	err = sameSellerPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountRate must be less than or equal to 100", err.Error())

	sameSellerPromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(1),
		CategoryP: &entity.CategoryPromotionDiscount{
			CategoryID:   1111,
			DiscountRate: 50,
		},
	}
	err = sameSellerPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "sameSellerP is required", err.Error())
}

func TestCategoryPromotionDiscount_Validate_Success(t *testing.T) {
	categoryPromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(2),
		CategoryP: &entity.CategoryPromotionDiscount{
			CategoryID:   1,
			DiscountRate: 10,
		},
	}

	err := categoryPromotion.Validate()
	assert.Nil(t, err)
}

func TestCategoryPromotionDiscount_Validate_Fail(t *testing.T) {
	categoryPromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(2),
		CategoryP: &entity.CategoryPromotionDiscount{
			CategoryID:   0,
			DiscountRate: 10,
		},
	}
	err := categoryPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "CategoryID is required", err.Error())

	categoryPromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(2),
		CategoryP: &entity.CategoryPromotionDiscount{
			CategoryID:   1,
			DiscountRate: 0,
		},
	}
	err = categoryPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountRate is required", err.Error())

	categoryPromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(2),
		CategoryP: &entity.CategoryPromotionDiscount{
			CategoryID:   1,
			DiscountRate: 150,
		},
	}
	err = categoryPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountRate must be less than or equal to 100", err.Error())

	categoryPromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(2),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 50,
		},
	}
	err = categoryPromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "categoryP is required", err.Error())
}

func TestTotalPricePromotionDiscount_Validate_Success(t *testing.T) {
	totalPricePromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 10,
				PriceRangeEnd:   100,
				DiscountAmount:  10,
			},
		},
	}

	err := totalPricePromotion.Validate()
	assert.Nil(t, err)
}

func TestTotalPricePromotionDiscount_Validate_Fail(t *testing.T) {
	totalPricePromotion := &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 0,
				PriceRangeEnd:   100,
				DiscountAmount:  10,
			},
		},
	}
	err := totalPricePromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "PriceRangeStart is required", err.Error())

	totalPricePromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 10,
				PriceRangeEnd:   0,
				DiscountAmount:  10,
			},
		},
	}
	err = totalPricePromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "PriceRangeEnd is required", err.Error())

	totalPricePromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 10,
				PriceRangeEnd:   100,
				DiscountAmount:  0,
			},
		},
	}
	err = totalPricePromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountAmount is required", err.Error())

	totalPricePromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 500,
				PriceRangeEnd:   1000,
				DiscountAmount:  0,
			},
		},
	}
	err = totalPricePromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "DiscountAmount is required", err.Error())

	totalPricePromotion = &entity.Promotion{
		PromotionType: entity.PromotionType(3),
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 50,
		},
	}
	err = totalPricePromotion.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "totalPriceP is required", err.Error())
}
