package main

import (
	"fmt"
	"log/slog"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/persistence"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// * Setup Config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config: %v", err)
	}

	fmt.Print("Config is loaded: ", cfg)

	// * Setup Server
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// * Setup Prometheus
	promLogger := logger.NewPrometheusLogger()
	router.Use(logger.PrometheusMiddleware(promLogger))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// * Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// * Swagger Routes for development
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// * Setup Router
	apiRouter := router.Group("/api/v1")
	repositories, err := persistence.NewRepositories(cfg.Mongo)
	if err != nil {
		slog.Error("Failed to create repositories: %v", err)
	}
	defer repositories.Close()

	interfaces.RegisterRoutes(apiRouter, repositories)

	fmt.Print("Server is running on port: " + cfg.Server.Port + "\n")
	router.Run(":" + cfg.Server.Port)
}
