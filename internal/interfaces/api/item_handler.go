package api

import (
	"strconv"

	_ "github.com/Furkan-Gulsen/Checkout-System/docs"
	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemApp application.ItemAppInterface
}

func NewItemHandler(itemApp application.ItemAppInterface) *ItemHandler {
	return &ItemHandler{
		itemApp: itemApp,
	}
}

// @Summary List items
// @Description Get a list of items
// @Tags Item
// @Produce json
// @Success 200 {object} []entity.Item
// @Router /api/v1/item/list [get]
func (h *ItemHandler) List(c *gin.Context) {
	items, err := h.itemApp.List()
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
// @Tags Item
// @Accept json
// @Produce json
// @Param item body entity.Item true "Item object"
// @Success 200 {string} string "Item created successfully"
// @Router /api/v1/item/create [post]
func (h *ItemHandler) Create(c *gin.Context) {
	var item entity.Item

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

	if err := h.itemApp.Create(&item); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "Item created successfully"})

}

// @Summary Get an item by ID
// @Description Get an item by its ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID" Format(int)
// @Success 200 {object} string
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/item/{id} [get]
func (h *ItemHandler) GetById(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(400, gin.H{"message": "ID is required", "status": false})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format", "status": false})
		return
	}

	item, err := h.itemApp.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(200, gin.H{"message": item, "status": true})
}

// @Summary Delete an item by ID
// @Description Delete an item by its ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID" Format(int)
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/item/{id} [delete]
func (h *ItemHandler) Delete(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(400, gin.H{"message": "ID is required", "status": false})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format", "status": false})
		return
	}

	if err := h.itemApp.Delete(id); err != nil {
		c.JSON(500, gin.H{"message": err.Error(), "status": false})
		return
	}

	c.JSON(200, gin.H{"message": "Item deleted successfully", "status": true})
}
