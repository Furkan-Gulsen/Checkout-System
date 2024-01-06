package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	collection *mongo.Collection
}

func NewItemRepository(d *database.Database, dbName string) *ItemRepository {
	return &ItemRepository{
		collection: d.Collection(dbName, "items"),
	}
}

// ItemRepository implements repository.ItemRepositoryI interface
var _ repository.ItemRepositoryI = &ItemRepository{}

func (r *ItemRepository) ListByCartId(cartId int) ([]*entity.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var items []*entity.Item

	cursor, err := r.collection.Find(ctx, bson.M{"cartId": cartId})
	if err != nil {
		return items, err
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *ItemRepository) Create(item *entity.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	item.Id = utils.GenerateID()

	_, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetById(itemID int) (*entity.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var item entity.Item

	err := r.collection.FindOne(ctx, bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		return &entity.Item{}, fmt.Errorf("could not find item with id %d: %w", itemID, err)
	}

	return &item, nil
}

func (r *ItemRepository) Delete(itemID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": itemID})
	if err != nil {
		return err
	}

	return nil
}
