package api

import (
	"fmt"
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/gin-gonic/gin"
)

type VasItemHandler struct {
	vasItemApp application.VasItemAppInterface
}

func NewVasItemHandler(vasItemApp application.VasItemAppInterface) *VasItemHandler {
	return &VasItemHandler{
		vasItemApp: vasItemApp,
	}
}

// @Summary List vas items
// @Description Get a list of vas items
// @Tags VasItem
// @Produce json
// @Param item_id query int true "Item ID" Format(int)
// @Success 200 {object} []entity.VasItem
// @Router /api/v1/vasitem/list [get]
func (h *VasItemHandler) ListByItemId(c *gin.Context) {
	query := c.Request.URL.Query()
	queryItemId := query.Get("item_id")

	if queryItemId == "" {
		c.JSON(400, gin.H{"message": "Item ID is required"})
		return
	}

	itemId, err := strconv.Atoi(queryItemId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Item ID format"})
		return
	}

	vasItems, err := h.vasItemApp.ListByItemId(itemId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if len(vasItems) == 0 {
		c.JSON(200, gin.H{"message": "No vas items found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vas items found",
		"data":    vasItems,
	})
}

// @Summary Get vas item
// @Description Get a vas item
// @Tags VasItem
// @Produce json
// @Param id path int true "Vas Item ID" Format(int)
// @Success 200 {object} []entity.VasItem
// @Router /api/v1/vasitem/{id} [get]
func (h *VasItemHandler) GetById(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(400, gin.H{"message": "Vas Item ID is required"})
		return
	}

	fmt.Println("paramID: ", paramID)

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format"})
		return
	}

	vasItem, err := h.vasItemApp.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if vasItem == nil {
		c.JSON(200, gin.H{"message": "No vas item found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vas item found",
		"data":    vasItem,
	})
}
