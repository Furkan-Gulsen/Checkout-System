package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/middleware"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/logger"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// * Load configuration settings
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration: ", err)
		return
	}

	// * Create a Gin router and set up its middleware
	router := setupRouter(cfg)
	server := setupServer(cfg.Server.Port, router)

	// * Graceful Shutdown
	go utils.Graceful(server, 10*time.Second)

	// * Start the HTTP server
	if err := startServer(server); err != nil {
		slog.Error("Failed to start the server: ", err)
	}
}

func setupRouter(cfg *config.Config) *gin.Engine {
	// * Create a new Gin router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// * Set up Prometheus metrics collection
	promLogger := logger.NewPrometheusLogger()
	router.Use(logger.PrometheusMiddleware(promLogger))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// * Define a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// * Serve Swagger documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// * Create an API router group for version 1
	apiRouter := router.Group("/api/v1")
	interfaces.RegisterRoutes(apiRouter, cfg)

	return router
}

func setupServer(port string, router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return server
}

func startServer(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
