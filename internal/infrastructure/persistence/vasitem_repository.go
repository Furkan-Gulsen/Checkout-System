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

type VasItemRepository struct {
	collection *mongo.Collection
}

func NewVasItemRepository(d *database.Database, dbName string) *VasItemRepository {
	return &VasItemRepository{
		collection: d.Collection(dbName, "vasitems"),
	}
}

// VasItemRepository implements repository.VasItemRepositoryI interface
var _ repository.VasItemRepositoryI = &VasItemRepository{}

func (r *VasItemRepository) ListByItemId(itemId int) ([]*entity.VasItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var vasItems []*entity.VasItem

	cursor, err := r.collection.Find(ctx, bson.M{"itemId": itemId})
	if err != nil {
		return vasItems, err
	}

	err = cursor.All(ctx, &vasItems)
	if err != nil {
		return vasItems, err
	}

	return vasItems, nil
}

func (r *VasItemRepository) GetById(vasItemId int) (*entity.VasItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var vasItem *entity.VasItem

	err := r.collection.FindOne(ctx, bson.M{"_id": vasItemId}).Decode(&vasItem)
	if err != nil {
		return vasItem, err
	}

	return vasItem, nil
}

func (r *VasItemRepository) Create(vasItem *entity.VasItem) (*entity.VasItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	vasItem.Id = utils.GenerateID()

	_, err := r.collection.InsertOne(ctx, vasItem)
	if err != nil {
		return nil, fmt.Errorf("error while creating vas item: %v", err)
	}

	return vasItem, nil
}

func (r *VasItemRepository) DeleteById(vasItemId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": vasItemId})
	if err != nil {
		return err
	}

	return nil
}
