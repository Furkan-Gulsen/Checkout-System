package item

import (
	"strconv"

	_ "github.com/Furkan-Gulsen/Checkout-System/docs"
	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine, db *database.Database, dbName string) {

	itemRouterGroup := g.Group("/item")

	itemRepository := NewItemRepository(db, dbName)
	itemService := NewItemService(*itemRepository)
	itemRouter := NewItemRouter(*itemService)

	itemRouterGroup.GET("/list", itemRouter.list)
	itemRouterGroup.POST("/create", itemRouter.create)
	itemRouterGroup.GET("/:id", itemRouter.getById)
	itemRouterGroup.DELETE("/:id", itemRouter.delete)
}

type ItemRouter struct {
	ItemService ItemService
}

func NewItemRouter(itemService ItemService) *ItemRouter {
	return &ItemRouter{
		ItemService: itemService,
	}
}

// @Summary List items
// @Description Get a list of items
// @Produce json
// @Success 200 {object} []entities.Item
// @Router /item/list [get]
func (r *ItemRouter) list(c *gin.Context) {
	items, err := r.ItemService.List()
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": items,
	})
}

// @Summary Create an item
// @Description Create a new item
// @Accept json
// @Produce json
// @Param item body entities.Item true "Item object"
// @Success 200 {string} string "Item created successfully"
// @Router /item/create [post]
func (r *ItemRouter) create(c *gin.Context) {
	var item entities.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	err := item.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if err := r.ItemService.Create(item); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "Item created successfully"})

}

// @Summary Get an item by ID
// @Description Get an item by its ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID" Format(int64)
// @Success 200 {object} string
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /item/{id} [get]
func (h *ItemRouter) getById(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID == "" {
		ctx.JSON(400, gin.H{"message": "ID is required", "status": false})
		return
	}

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid ID format", "status": false})
		return
	}

	item, err := h.ItemService.GetById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error(), "status": false})
		return
	}

	ctx.JSON(200, gin.H{"message": item, "status": true})
}

// @Summary Delete an item by ID
// @Description Delete an item by its ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID" Format(int64)
// @Success 200 {string} string "Item deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /item/{id} [delete]
func (h *ItemRouter) delete(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID == "" {
		ctx.JSON(400, gin.H{"message": "ID is required", "status": false})
		return
	}

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid ID format", "status": false})
		return
	}

	if err := h.ItemService.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"message": err.Error(), "status": false})
		return
	}

	ctx.JSON(200, gin.H{"message": "Item deleted successfully", "status": true})
}
