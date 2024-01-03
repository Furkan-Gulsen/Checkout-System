package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/product"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// * Setup Config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config: %v", err)
	}

	fmt.Print("Config is loaded: ", cfg)

	// * Setup Database
	ctxDBTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60) // * 1 minute timeout for database connection
	mongodbURI := "mongodb://" + cfg.Mongo.Host + ":" + cfg.Mongo.Port
	db, err := database.Connect(ctxDBTimeout, mongodbURI)
	if err != nil {
		slog.Error("Failed to setup database: %v", err)
	}
	defer db.Disconnect()
	defer cancel()

	// * Setup Server
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// * Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// * Setup Product Router
	product.SetupRouter(router)

	fmt.Print("Server is running on port: " + cfg.Server.Port + "\n")
	router.Run(":" + cfg.Server.Port)
}
