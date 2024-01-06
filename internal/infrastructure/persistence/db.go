package persistence

import (
	"context"
	"log/slog"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Item     repository.ItemRepositoryI
	Category repository.CategoryRepositoryI
	db       *mongo.Collection
}

func NewRepositories(cfg config.MongoDBConfig) (*Repositories, error) {
	// * Setup Database
	ctxDBTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60) // * 1 minute timeout for database connection
	mongodbURI := "mongodb://" + cfg.Host + ":" + cfg.Port
	db, err := database.Connect(ctxDBTimeout, mongodbURI)
	if err != nil {
		slog.Error("Failed to setup database: %v", err)
	}
	defer cancel()

	return &Repositories{
		Item:     NewItemRepository(db, cfg.Database),
		Category: NewCategoryRepository(db, cfg.Database),
		db:       db.Collection(cfg.Database, "items"),
	}, nil
}

func (r *Repositories) Close() {
	ctxDBTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if err := r.db.Database().Client().Disconnect(ctxDBTimeout); err != nil {
		slog.Error("Failed to close database connection: %v", err)
	}
}

func (r *Repositories) Ping() error {
	ctxDBTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if err := r.db.Database().Client().Ping(ctxDBTimeout, nil); err != nil {
		return err
	}

	return nil
}
