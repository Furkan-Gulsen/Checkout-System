package persistence

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/constants"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func setUpPromotionRepo(t *testing.T) (*PromotionRepository, func()) {
	db := utils.SetUpMockDatabase(t)
	repo := NewPromotionRepository(db, constants.MOCK_DB_NAME)

	return repo, func() {
		utils.CleanUpMockDatabase(db, "promotions")
	}
}

func TestCreateSameSellerPromotion_Success(t *testing.T) {
	repo, tearDown := setUpPromotionRepo(t)
	defer tearDown()

	data := &dto.PromotionRequest{
		PromotionType: 1,
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 10.0,
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.Nil(t, validateErr)

	promotion, err := repo.Create(&dataEntity)
	assert.Nil(t, err)
	assert.NotNil(t, promotion)
	assert.Equal(t, promotion.PromotionType, entity.SameSellerPromotion)
	assert.Equal(t, promotion.SameSellerP.DiscountRate, 10.0)
}

func TestCreateCategoryPromotion_Success(t *testing.T) {
	repo, tearDown := setUpPromotionRepo(t)
	defer tearDown()

	data := &dto.PromotionRequest{
		PromotionType: 2,
		CategoryP: &entity.CategoryPromotionDiscount{
			DiscountRate: 10.0,
			CategoryID:   1,
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.Nil(t, validateErr)

	promotion, err := repo.Create(&dataEntity)
	assert.Nil(t, err)
	assert.NotNil(t, promotion)
	assert.Equal(t, promotion.PromotionType, entity.CategoryPromotion)
	assert.Equal(t, promotion.CategoryP.DiscountRate, 10.0)
	assert.Equal(t, promotion.CategoryP.CategoryID, 1)
}

func TestCreateTotalPricePromotion_Success(t *testing.T) {
	repo, tearDown := setUpPromotionRepo(t)
	defer tearDown()

	data := &dto.PromotionRequest{
		PromotionType: 3,
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 500,
				PriceRangeEnd:   1000,
				DiscountAmount:  50,
			},
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.Nil(t, validateErr)

	promotion, err := repo.Create(&dataEntity)
	assert.Nil(t, err)
	assert.NotNil(t, promotion)
	assert.Equal(t, promotion.PromotionType, entity.TotalPricePromotion)
	assert.Equal(t, promotion.TotalPriceP[0].PriceRangeStart, 500.0)
	assert.Equal(t, promotion.TotalPriceP[0].PriceRangeEnd, 1000.0)
	assert.Equal(t, promotion.TotalPriceP[0].DiscountAmount, 50.0)
}

func TestCreateSameSellerPromotionValidation_Failure(t *testing.T) {
	data := &dto.PromotionRequest{
		PromotionType: 1,
		SameSellerP: &entity.SameSellerPromotionDiscount{
			DiscountRate: 101.0,
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.NotNil(t, validateErr)
	assert.EqualError(t, validateErr, "DiscountRate must be less than or equal to 100")
}

func TestCreateCategoryPromotionValidation_Failure(t *testing.T) {
	data := &dto.PromotionRequest{
		PromotionType: 2,
		CategoryP: &entity.CategoryPromotionDiscount{
			DiscountRate: 10.0,
			CategoryID:   0,
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.NotNil(t, validateErr)
	assert.EqualError(t, validateErr, "CategoryID is required")
}

func TestCreateTotalPricePromotionValidation_Failure(t *testing.T) {
	data := &dto.PromotionRequest{
		PromotionType: 3,
		TotalPriceP: []*entity.TotalPricePromotionDiscount{
			{
				PriceRangeStart: 1000,
				PriceRangeEnd:   500,
				DiscountAmount:  0,
			},
		},
	}

	dataEntity := data.ToEntity()
	validateErr := dataEntity.Validate()
	assert.NotNil(t, validateErr)
	assert.EqualError(t, validateErr, "DiscountAmount is required")
}
