package interfaces

import (
	"log/slog"

	"github.com/Furkan-Gulsen/Checkout-System/config"
	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/persistence"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, cfg *config.Config) {
	repositories, err := persistence.NewRepositories(cfg.Mongo)
	if err != nil {
		slog.Error("Failed to create data repositories: ", err)
	}

	// * Application Layer
	itemApp := application.NewItemApp(repositories.Item, repositories.Category)
	categoryApp := application.NewCategoryApp(repositories.Category)
	promotionApp := application.NewPromotionApp(repositories.Promotion)

	// * Handlers
	itemHandler := api.NewItemHandler(itemApp)
	categoryHandler := api.NewCategoryHandler(categoryApp)
	promotionHandler := api.NewPromotionHandler(promotionApp)

	// * Category Routes
	categoryRouterGroup := g.Group("/category")
	categoryRouterGroup.GET("/list", categoryHandler.List)
	categoryRouterGroup.POST("/create", categoryHandler.Create)
	categoryRouterGroup.GET("/:id", categoryHandler.GetById)

	// * Item Routes
	itemRouterGroup := g.Group("/item")
	itemRouterGroup.GET("/list", itemHandler.List)
	itemRouterGroup.POST("/create", itemHandler.Create)
	itemRouterGroup.GET("/:id", itemHandler.GetById)
	itemRouterGroup.DELETE("/:id", itemHandler.Delete)

	// * Promotion Routes
	promotionRouterGroup := g.Group("/promotion")
	promotionRouterGroup.GET("/list", promotionHandler.List)
	promotionRouterGroup.POST("/create", promotionHandler.Create)
	promotionRouterGroup.GET("/:id", promotionHandler.GetById)
}
