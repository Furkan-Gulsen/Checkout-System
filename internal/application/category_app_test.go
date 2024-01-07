package application

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type mockCategoryRepository struct{}

var (
	listCategoryRepo    func() ([]*entity.Category, error)
	createCategoryRepo  func(category *entity.Category) (*entity.Category, error)
	getByIDCategoryRepo func(id int) (*entity.Category, error)
)

func (m *mockCategoryRepository) List() ([]*entity.Category, error) {
	return listCategoryRepo()
}

func (m *mockCategoryRepository) Create(category *entity.Category) (*entity.Category, error) {
	return createCategoryRepo(category)
}

func (m *mockCategoryRepository) GetByID(id int) (*entity.Category, error) {
	return getByIDCategoryRepo(id)
}

var categoryAppMock CategoryAppInterface = &mockCategoryRepository{}

func TestSaveCategory_Success(t *testing.T) {
	createCategoryRepo = func(category *entity.Category) (*entity.Category, error) {
		return &entity.Category{
			Id:   1,
			Name: "Category1",
		}, nil
	}

	category := &entity.Category{
		Id:   1,
		Name: "Category1",
	}

	category, err := categoryAppMock.Create(category)
	assert.Nil(t, err)
	assert.Equal(t, 1, category.Id)
	assert.Equal(t, "Category1", category.Name)

}

func TestGetCategoryByID_Success(t *testing.T) {
	getByIDCategoryRepo = func(id int) (*entity.Category, error) {
		return &entity.Category{
			Id:   1,
			Name: "Category1",
		}, nil
	}

	category, err := categoryAppMock.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, category.Id)
	assert.Equal(t, "Category1", category.Name)
}

func TestGetCategoryByID_Fail(t *testing.T) {
	getByIDCategoryRepo = func(id int) (*entity.Category, error) {
		return nil, nil
	}

	category, err := categoryAppMock.GetByID(300)
	assert.Nil(t, category)
	assert.Nil(t, err)
}

func TestListCategory_Success(t *testing.T) {
	listCategoryRepo = func() ([]*entity.Category, error) {
		return []*entity.Category{
			{
				Id:   1,
				Name: "Category1",
			},
			{
				Id:   2,
				Name: "Category2",
			},
		}, nil
	}

	categories, err := categoryAppMock.List()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(categories))
	assert.Equal(t, 1, categories[0].Id)
	assert.Equal(t, "Category1", categories[0].Name)
	assert.Equal(t, 2, categories[1].Id)
	assert.Equal(t, "Category2", categories[1].Name)
}

func TestListCategory_Fail(t *testing.T) {
	listCategoryRepo = func() ([]*entity.Category, error) {
		return nil, nil
	}

	categories, err := categoryAppMock.List()
	assert.Nil(t, categories)
	assert.Nil(t, err)
}

func TestListCategory_Empty(t *testing.T) {
	listCategoryRepo = func() ([]*entity.Category, error) {
		return []*entity.Category{}, nil
	}

	categories, err := categoryAppMock.List()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(categories))
}
