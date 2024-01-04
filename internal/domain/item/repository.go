package item

import (
	"context"
	"fmt"

	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infra/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepositoryInterface interface {
	List(ctx context.Context) ([]entities.Item, error)
	Create(ctx context.Context, item entities.Item) error
	GetById(ctx context.Context, itemID int64) error
	Delete(ctx context.Context, itemID string) error
}

type ItemRepository struct {
	collection *mongo.Collection
}

func NewItemRepository(d *database.Database, dbName string) *ItemRepository {
	return &ItemRepository{
		collection: d.Collection(dbName, "items"),
	}
}

func (r *ItemRepository) List(ctx context.Context) ([]entities.Item, error) {
	var items []entities.Item

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) Create(ctx context.Context, item entities.Item) error {
	item.Id = int64(uuid.New().ID())
	_, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetById(ctx context.Context, itemID int64) (entities.Item, error) {
	var item entities.Item

	err := r.collection.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return entities.Item{}, fmt.Errorf("could not find item with id %d: %w", itemID, err)
	}

	return item, nil
}

func (r *ItemRepository) Delete(ctx context.Context, itemID int64) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": itemID})
	if err != nil {
		return err
	}

	return nil
}
