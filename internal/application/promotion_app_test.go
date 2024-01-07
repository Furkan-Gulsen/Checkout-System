package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
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

var promotionAppMock PromotionAppInterface = &mockPromotionRepository{}

func TestSavePromotion_Success(t *testing.T) {
	createPromotionRepo = func(promotion *entity.Promotion) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1,
			PromotionType: entity.PromotionType(1),
		}, nil
	}

	promotion := &entity.Promotion{
		Id:            1,
		PromotionType: entity.PromotionType(1),
	}

	promotion, err := promotionAppMock.Create(promotion)
	assert.Nil(t, err)
	assert.Equal(t, 1, promotion.Id)
	assert.Equal(t, entity.PromotionType(1), promotion.PromotionType)

}

func TestGetPromotionByID_Success(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return &entity.Promotion{
			Id:            1,
			PromotionType: entity.PromotionType(1),
		}, nil
	}

	promotion, err := promotionAppMock.GetById(1)
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

	promotions, err := promotionAppMock.List()
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

	promotions, err := promotionAppMock.List()
	assert.Nil(t, promotions)
	assert.Nil(t, err)
}

func TestGetPromotionByID_Fail(t *testing.T) {
	getByIDPromotionRepo = func(id int) (*entity.Promotion, error) {
		return nil, nil
	}

	promotion, err := promotionAppMock.GetById(300)
	assert.Nil(t, promotion)
	assert.Nil(t, err)
}

func TestSavePromotion_Fail(t *testing.T) {
	createPromotionRepo = func(promotion *entity.Promotion) (*entity.Promotion, error) {
		return nil, nil
	}

	promotion := &entity.Promotion{
		Id:            1,
		PromotionType: entity.PromotionType(1),
	}

	promotion, err := promotionAppMock.Create(promotion)
	assert.Nil(t, promotion)
	assert.Nil(t, err)
}
