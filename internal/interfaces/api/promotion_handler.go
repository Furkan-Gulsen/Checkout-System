package api

import (
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
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
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Promotions listed successfully",
		"data":    promotions,
	})
}

// @Summary Create a promotion
// @Description Create a new promotion
// @Tags Promotion
// @Accept json
// @Produce json
// @Param promotion body dto.PromotionRequest true "Promotion object"
// @Success 200 {string} string "Promotion created successfully"
// @Router /api/v1/promotion [post]
func (h *PromotionHandler) Create(c *gin.Context) {
	var data dto.PromotionRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	entityData := data.ToEntity()
	err := entityData.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, createErr := h.promotionApp.Create(&entityData)
	if createErr != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Promotion created successfully",
		"data":    response,
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
		c.JSON(400, gin.H{"message": "ID is required"})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format"})
		return
	}

	promotion, err := h.promotionApp.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Promotion found",
		"data":    promotion,
	})
}
