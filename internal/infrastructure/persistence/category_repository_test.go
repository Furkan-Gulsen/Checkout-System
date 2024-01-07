package persistence

import (
	"context"
	"testing"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

var MOCK_DB_URL = "mongodb://localhost:27017"
var MOCK_DB_NAME = "mock_ty_case"

func setUpDatabase(t *testing.T) *database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := database.Connect(ctx, MOCK_DB_URL)
	if err != nil {
		t.Fatalf("Database connection error: %v", err)
	}

	return db
}

func createTestRepo(db *database.Database) *CategoryRepository {
	return NewCategoryRepository(db, MOCK_DB_NAME)
}

func cleanUpDatabase(db *database.Database) {
	db.Client.Database(MOCK_DB_NAME).Collection("categories").Drop(context.Background())
}

func setUp(t *testing.T) (*CategoryRepository, func()) {
	db := setUpDatabase(t)
	repo := createTestRepo(db)

	return repo, func() {
		cleanUpDatabase(db)
	}
}

func TestCreateCategory_Success(t *testing.T) {
	repo, tearDown := setUp(t)
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
	repo, tearDown := setUp(t)
	defer tearDown()

	_, _ = repo.Create(&entity.Category{Id: 1, Name: "Category1", ItemType: 1})
	_, _ = repo.Create(&entity.Category{Id: 2, Name: "Category2", ItemType: 2})

	categories, err := repo.List()
	assert.Nil(t, err)
	assert.Len(t, categories, 2)
}

func TestGetCategoryByID_Success(t *testing.T) {
	repo, tearDown := setUp(t)
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
	repo, tearDown := setUp(t)
	defer tearDown()

	category, err := repo.GetByID(8488484848484)
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestGetCategoryByID_InvalidID(t *testing.T) {
	repo, tearDown := setUp(t)
	defer tearDown()

	category, err := repo.GetByID(0)
	assert.NotNil(t, err)
	assert.Nil(t, category)
}
