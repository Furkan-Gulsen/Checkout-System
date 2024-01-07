package utils

import (
	"context"
	"testing"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/database"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/constants"
)

func SetUpMockDatabase(t *testing.T) *database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := database.Connect(ctx, constants.MOCK_DB_URL)
	if err != nil {
		t.Fatalf("Database connection error: %v", err)
	}

	return db
}

func CleanUpMockDatabase(db *database.Database, collName string) {
	db.Client.Database(constants.MOCK_DB_NAME).Collection(collName).Drop(context.Background())
}
