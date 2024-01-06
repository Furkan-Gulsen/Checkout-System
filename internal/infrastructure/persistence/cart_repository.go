package persistence

import (
	"context"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	collection *mongo.Collection
}

func NewCartRepository(d *database.Database, dbName string) *CartRepository {
	return &CartRepository{
		collection: d.Collection(dbName, "carts"),
	}
}

// CartRepository implements repository.CartRepositoryI interface
var _ repository.CartRepositoryI = &CartRepository{}

func (r *CartRepository) Create(cart entity.Cart) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cart.Id = utils.GenerateID()

	_, err := r.collection.InsertOne(ctx, cart)
	if err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) GetByID(id int) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var cart entity.Cart

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&cart)
	if err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) Update(cart entity.Cart) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": cart.Id}, bson.M{"$set": cart})
	if err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
