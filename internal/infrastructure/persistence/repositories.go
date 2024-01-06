package persistence

import (
	"context"
	"log/slog"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
)

type Repositories struct {
	Item      repository.ItemRepositoryI
	Category  repository.CategoryRepositoryI
	Promotion repository.PromotionRepositoryI
	db        *database.Database
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
		Item:      NewItemRepository(db, cfg.Database),
		Category:  NewCategoryRepository(db, cfg.Database),
		Promotion: NewPromotionRepository(db, cfg.Database),
		db:        db,
	}, nil
}

func (r *Repositories) Close() {
	r.db.Disconnect()
}
