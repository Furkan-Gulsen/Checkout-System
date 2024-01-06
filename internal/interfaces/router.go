package interfaces

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/infrastructure/persistence"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, repo *persistence.Repositories) {
	// * Category Routes
	categoryRouter := api.NewCategoryRouter(repo.Category)

	categoryRouterGroup := g.Group("/category")
	categoryRouterGroup.GET("/list", categoryRouter.List)
	categoryRouterGroup.POST("/create", categoryRouter.Create)
	categoryRouterGroup.GET("/:id", categoryRouter.GetById)

	// * Item Routes
	itemRouter := api.NewItemHandler(repo.Item)

	itemRouterGroup := g.Group("/item")
	itemRouterGroup.GET("/list", itemRouter.List)
	itemRouterGroup.POST("/create", itemRouter.Create)
	itemRouterGroup.GET("/:id", itemRouter.GetById)
	itemRouterGroup.DELETE("/:id", itemRouter.Delete)

}
