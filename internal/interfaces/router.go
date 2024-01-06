package interfaces

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/persistence"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, repo *persistence.Repositories) {
	// * Handlers for routes
	categoryHandler := api.NewCategoryHandler(repo.Category)
	itemHandler := api.NewItemHandler(repo.Item)

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

}
