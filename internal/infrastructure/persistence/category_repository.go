package persistence

import (
	"context"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(d *database.Database, dbName string) *CategoryRepository {
	return &CategoryRepository{
		collection: d.Collection(dbName, "categories"),
	}
}

// CategoryRepository implements repository.CategoryRepositoryI interface
var _ repository.CategoryRepositoryI = &CategoryRepository{}

func (r *CategoryRepository) List() ([]entity.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var categories []entity.Category

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Create(category entity.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	category.Id = int(uuid.New().ID())

	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) GetByID(id int) (entity.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var category entity.Category

	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category); err != nil {
		return entity.Category{}, err
	}

	return category, nil
}
