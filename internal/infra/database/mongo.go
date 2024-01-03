package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client  *mongo.Client
	Context context.Context
	Address string
}

func Connect(ctx context.Context, address string) (*Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Printf("Connected to database: %s\n", address)
	return &Database{
		Client:  client,
		Context: ctx,
		Address: address,
	}, nil
}

func (d *Database) Disconnect() error {
	if err := d.Client.Disconnect(d.Context); err != nil {
		return fmt.Errorf("failed to disconnect from database: %w", err)
	}
	fmt.Printf("Disconnected from database: %s\n", d.Address)
	return nil
}

func (d *Database) Collection(dbName string, collectionName string) *mongo.Collection {
	return d.Client.Database(dbName).Collection(collectionName)
}
