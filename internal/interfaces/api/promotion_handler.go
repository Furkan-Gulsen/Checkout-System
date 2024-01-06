package api

import (
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type PromotionHandler struct {
	promotionApp application.PromotionAppInterface
}

func NewPromotionHandler(promotionApp application.PromotionAppInterface) *PromotionHandler {
	return &PromotionHandler{
		promotionApp: promotionApp,
	}
}

// @Summary List promotions
// @Description Get a list of promotions
// @Tags Promotion
// @Produce json
// @Success 200 {object} []entity.Promotion
// @Router /api/v1/promotion/list [get]
func (h *PromotionHandler) List(c *gin.Context) {
	promotions, err := h.promotionApp.List()
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": promotions,
	})
}

// @Summary Create a promotion
// @Description Create a new promotion
// @Tags Promotion
// @Accept json
// @Produce json
// @Param promotion body entity.Promotion true "Promotion object"
// @Success 200 {string} string "Promotion created successfully"
// @Router /api/v1/promotion/create [post]
func (h *PromotionHandler) Create(c *gin.Context) {
	var promotion entity.Promotion

	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	err := promotion.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if err := h.promotionApp.Create(&promotion); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Promotion created successfully",
	})
}

// @Summary Get a promotion by ID
// @Description Get a promotion by its ID
// @Tags Promotion
// @Produce json
// @Param id path int true "Promotion ID"
// @Success 200 {object} string
// @Router /api/v1/promotion/{id} [get]
func (h *PromotionHandler) GetById(c *gin.Context) {
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

	promotion, err := h.promotionApp.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": promotion,
	})
}