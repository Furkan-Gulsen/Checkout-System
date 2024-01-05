package category

import (
	"context"

	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infra/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryInterface interface {
	List(ctx context.Context) ([]entities.Category, error)
	Create(ctx context.Context, category entities.Category) error
}

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(d *database.Database, dbName string) *CategoryRepository {
	return &CategoryRepository{
		collection: d.Collection(dbName, "categories"),
	}
}

func (r *CategoryRepository) List(ctx context.Context) ([]entities.Category, error) {
	var categories []entities.Category

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category entities.Category) error {
	category.Id = int(uuid.New().ID())

	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}
