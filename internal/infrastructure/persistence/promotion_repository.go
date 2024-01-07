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

type PromotionRepository struct {
	collection *mongo.Collection
}

func NewPromotionRepository(d *database.Database, dbName string) *PromotionRepository {
	return &PromotionRepository{
		collection: d.Collection(dbName, "promotions"),
	}
}

// PromotionRepository implements repository.PromotionRepositoryI interface
var _ repository.PromotionRepositoryI = &PromotionRepository{}

func (r *PromotionRepository) List() ([]*entity.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var promotions []*entity.Promotion

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &promotions); err != nil {
		return nil, err
	}

	return promotions, nil
}

func (r *PromotionRepository) Create(promotion *entity.Promotion) (*entity.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	promotion.Id = utils.GenerateID()

	_, err := r.collection.InsertOne(ctx, promotion)
	if err != nil {
		return nil, fmt.Errorf("error while creating promotion: %v", err)
	}

	return promotion, nil
}

func (r *PromotionRepository) GetById(promotionID int) (*entity.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var promotion entity.Promotion

	err := r.collection.FindOne(ctx, bson.M{"_id": promotionID}).Decode(&promotion)
	if err != nil {
		return nil, err
	}

	return &promotion, nil
}
