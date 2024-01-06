package api

import (
	"fmt"
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
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
		c.JSON(400, gin.H{"status": false, "message": "Item ID is required"})
		return
	}

	itemId, err := strconv.Atoi(queryItemId)
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": "Invalid Item ID format"})
		return
	}

	vasItems, err := h.vasItemApp.ListByItemId(itemId)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	if len(vasItems) == 0 {
		c.JSON(200, gin.H{"status": false, "message": "No vas items found"})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": vasItems,
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
		c.JSON(400, gin.H{"status": false, "message": "Vas Item ID is required"})
		return
	}

	fmt.Println("paramID: ", paramID)

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format", "status": false})
		return
	}

	vasItem, err := h.vasItemApp.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	if vasItem == nil {
		c.JSON(200, gin.H{"status": false, "message": "No vas item found"})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": vasItem,
	})
}

// @Summary Create a vas item
// @Description Create a new vas item
// @Tags VasItem
// @Accept json
// @Produce json
// @Param vas_item body entity.VasItem true "Vas Item object"
// @Success 200 {string} string "Vas Item created successfully"
// @Router /api/v1/vasitem [post]
func (h *VasItemHandler) Create(c *gin.Context) {
	var vasItem entity.VasItem

	if err := c.ShouldBindJSON(&vasItem); err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	if err := h.vasItemApp.Create(&vasItem); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Vas Item created successfully",
	})
}
