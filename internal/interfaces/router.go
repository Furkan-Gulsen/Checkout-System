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
	categoryApp := application.NewCategoryApp(repositories.Category)
	promotionApp := application.NewPromotionApp(repositories.Promotion)
	itemApp := application.NewItemApp(repositories.Item)
	vasitemApp := application.NewVasItemApp(repositories.VasItem, itemApp)
	cartApp := application.NewCartApp(repositories.Cart, itemApp, vasitemApp, promotionApp)

	// * Handlers
	itemHandler := api.NewItemHandler(itemApp)
	categoryHandler := api.NewCategoryHandler(categoryApp)
	promotionHandler := api.NewPromotionHandler(promotionApp)
	vasitemHandler := api.NewVasItemHandler(vasitemApp)
	cartHandler := api.NewCartHandler(cartApp)

	// * Category Routes
	categoryRouterGroup := g.Group("/category")
	categoryRouterGroup.GET("/list", categoryHandler.List)
	categoryRouterGroup.POST("/", categoryHandler.Create)
	categoryRouterGroup.GET("/:id", categoryHandler.GetById)

	// * Item Routes
	itemRouterGroup := g.Group("/item")
	itemRouterGroup.GET("/list", itemHandler.ListByCartId)
	itemRouterGroup.GET("/:id", itemHandler.GetById)

	// * Promotion Routes
	promotionRouterGroup := g.Group("/promotion")
	promotionRouterGroup.GET("/list", promotionHandler.List)
	promotionRouterGroup.POST("/", promotionHandler.Create)
	promotionRouterGroup.GET("/:id", promotionHandler.GetById)

	// * VasItem Routes
	vasitemRouterGroup := g.Group("/vasitem")
	vasitemRouterGroup.GET("/list", vasitemHandler.ListByItemId)
	vasitemRouterGroup.GET("/:id", vasitemHandler.GetById)

	// * Cart Routes
	cartRouterGroup := g.Group("/cart")
	cartRouterGroup.POST("/:cartId/promotion/:promotionId", cartHandler.ApplyPromotion)
	cartRouterGroup.GET("/:cartId", cartHandler.Display)
	cartRouterGroup.DELETE("/:cartId", cartHandler.ResetCart)
	cartRouterGroup.POST("/:cartId/item", cartHandler.AddItem)
	cartRouterGroup.POST("/:cartId/item/:itemId/vas-item/:vasItemId", cartHandler.AddVasItem)
	cartRouterGroup.DELETE("/item/:itemId", cartHandler.RemoveItem)
	cartRouterGroup.DELETE("/vas-item/:vasItemId", cartHandler.RemoveVasItem)
}
