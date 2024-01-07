package persistence

import (
	"testing"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/constants"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func setUpCategoryRepo(t *testing.T) (*CategoryRepository, func()) {
	db := utils.SetUpMockDatabase(t)
	repo := NewCategoryRepository(db, constants.MOCK_DB_NAME)

	return repo, func() {
		utils.CleanUpMockDatabase(db, "categories")
	}
}

func TestCreateCategory_Success(t *testing.T) {
	repo, tearDown := setUpCategoryRepo(t)
	defer tearDown()

	category := &entity.Category{
		Id:       1,
		Name:     "Category1",
		ItemType: 1,
	}

	category, err := repo.Create(category)
	assert.Nil(t, err)
	assert.NotNil(t, category)
}

func TestCreateCategory_Failure(t *testing.T) {
	category := &entity.Category{
		Id:       1,
		Name:     "",
		ItemType: 4,
	}

	err := category.Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Validation errors: Name is required. ItemType must be 1 or 2.")
}

func TestListCategory(t *testing.T) {
	repo, tearDown := setUpCategoryRepo(t)
	defer tearDown()

	_, _ = repo.Create(&entity.Category{Id: 1, Name: "Category1", ItemType: 1})
	_, _ = repo.Create(&entity.Category{Id: 2, Name: "Category2", ItemType: 2})

	categories, err := repo.List()
	assert.Nil(t, err)
	assert.Len(t, categories, 2)
}

func TestGetCategoryByID_Success(t *testing.T) {
	repo, tearDown := setUpCategoryRepo(t)
	defer tearDown()

	expectedCategory := &entity.Category{Id: 1, Name: "Category1", ItemType: 1}
	category, err := repo.Create(expectedCategory)

	assert.Nil(t, err)
	assert.NotNil(t, category)

	category, err = repo.GetByID(category.Id)
	assert.Nil(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestGetCategoryByID_NotFound(t *testing.T) {
	repo, tearDown := setUpCategoryRepo(t)
	defer tearDown()

	category, err := repo.GetByID(8488484848484)
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestGetCategoryByID_InvalidID(t *testing.T) {
	repo, tearDown := setUpCategoryRepo(t)
	defer tearDown()

	category, err := repo.GetByID(0)
	assert.NotNil(t, err)
	assert.Nil(t, category)
}
